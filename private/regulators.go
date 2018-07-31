package private

import (
	"bytes"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/des"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

// TODO: refactor
var (
	RegContractAddress = common.HexToAddress("0x1932c48b2bF8102Ba33B4A6B545C32236e342f34")
	ErrNoRegulator     = errors.New("invalid transaction because no regulator provided")
	once               sync.Once
	ipc                = os.Getenv("IPC")
)

// RegulatorClient struct for communicating with the regulator
type RegulatorClient struct {
	contract  *des.Des
	client    *ethclient.Client
	events    chan types.Log
	regMap    map[string]bool
	whitelist map[string]bool
}

func (r *RegulatorClient) getClient() (b bool) {
	b = false
	var err error
	if r.client != nil {
		log.Debug("IPC Client exists")
		b = true
		return
	}
	// create once
	if _, err = os.Stat(ipc); err == nil {
		if ipc != "" {
			r.client, err = ethclient.Dial(ipc)
			if err != nil {
				log.Error("Failed to connect Ethereum client:", "error", err.Error())
			} else {
				log.Debug("Client successfully created: ", "client", r.client)
				b = true
			}
		} else {
			log.Error("Failure to get client since IPC url is not set in the environment.")
		}
	} else {
		log.Debug("IPC file not available yet")
	}
	return
}

// NewRegulatorClient initializes returns the regulator client
func NewRegulatorClient() (*RegulatorClient, error) {
	return &RegulatorClient{regMap: make(map[string]bool), whitelist: make(map[string]bool)}, nil
}

// IsRegulatorPresent checks if regulator is one of the privateFor
func (r *RegulatorClient) IsRegulatorPresent(from string, privateFor []string) (isPresent bool, err error) {
	isPresent = false
	err = nil
	// TODO: could make it more efficient
	if r.contract == nil {
		log.Trace("Contract is empty so trying to instantiate contract object")
		if r.getClient() {
			r.contract, err = des.NewDes(RegContractAddress, r.client)
			if err != nil || r.client == nil {
				log.Error("Failed to communicate with regulator contract", "error", err)
				return isPresent, nil
			}
		} else {
			log.Error("Can't seem to connect to an IPC client")
		}
		// r.contract = contract
		// r.client = client
	}

	// first check if sender is a regulator
	if from != "" {
		if isPresent = r.regMap[from]; isPresent {
			return
		}
		isPresent, err = r.contract.Exists(&bind.CallOpts{Context: context.Background()}, from)
		if err != nil {
			log.Error("Couldn't communicate with regulator contract", "error", err)
		} else if isPresent {
			r.regMap[from] = true
			return
		}
	}

	// check to see if any of the recipients are regulators
	for i := range privateFor {
		if isPresent = r.regMap[privateFor[i]]; isPresent {
			break
		}
		isPresent, err = r.contract.Exists(&bind.CallOpts{Context: context.Background()}, privateFor[i])
		if err != nil {
			log.Error("Couldn't communicate with regulator contract", "error", err)
			break
		} else if isPresent {
			r.regMap[privateFor[i]] = true
			break
		}
	}
	return
}

// IsRegulatorPresent checks if regulator is one of the privateFor
func (r *RegulatorClient) IsWhitelisted(enode string) (bool, error) {
	isWhitelisted := false
	var err error
	if r.getClient() {
		log.Trace("Contract is empty so trying to instantiate contract object")
		// if r.client == nil {
		// 	r.client, err = ethclient.Dial(ipc)
		// }
		r.contract, err = des.NewDes(RegContractAddress, r.client)
		if err != nil || r.client == nil {
			log.Error("Failed to communicate with permissions contract", "error", err)
			return isWhitelisted, nil
		}
		// r.contract = contract
	} else {
		log.Error("Can't seem to connect to an IPC client")
	}

	if isWhitelisted = r.whitelist[enode]; isWhitelisted {
		return isWhitelisted, nil
	}
	isWhitelisted, err = r.contract.NodeExists(&bind.CallOpts{Context: context.Background()}, enode)
	if err != nil {
		log.Error("Couldn't communicate with regulator contract", "error", err)
	} else if isWhitelisted {
		r.whitelist[enode] = true
	}

	return isWhitelisted, nil
}

// IsDeployed returns the state of the contract, whether it's been deployed yet.
func (r *RegulatorClient) IsDeployed() (deployed bool) {
	deployed = false
	// if r.client == nil {
	// 	r.client, err = ethclient.Dial(ipc)
	// 	defer r.client.Close()
	// 	if err != nil || r.client == nil {
	// 		log.Error("Error creating client", "err", err)
	// 		return
	// 	}
	// }
	if r.getClient() {
		code, err := r.client.CodeAt(context.Background(), RegContractAddress, nil)
		if err == nil {
			deployed = bytes.Equal(code, common.FromHex(des.DESByteCode))
			log.Debug("Checking whether contract has been deployed", "deployed", deployed, "codeLength", len(code))
			return
		} else {
			log.Error("Can't find contract code", "error", err)
		}
	} else {
		log.Error("Can't seem to connect to an IPC client")
	}
	log.Info("IsDeployed", "deployed", deployed)
	return
}
