package private

import (
	"bytes"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/des"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var (
	DESContractAddress = common.HexToAddress("0x1932c48b2bF8102Ba33B4A6B545C32236e342f34")
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
	deployed bool
	lastUpdated time.Time

}

// NewRegulatorClient initializes returns the regulator client
func NewRegulatorClient() *RegulatorClient {
	return &RegulatorClient{regMap: make(map[string]bool), whitelist: make(map[string]bool), deployed: false}
}

func (r *RegulatorClient) getClient() (b bool) {
	b = false
	var err error
	if r.client != nil {
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

// IsRegulatorPresent checks if regulator is one of the privateFor
func (r *RegulatorClient) IsRegulatorPresent(from string, privateFor []string) (isPresent bool, err error) {
	isPresent = false
	err = nil
	// TODO: could make it more efficient
	if r.contract == nil {
		log.Trace("Contract is empty so trying to instantiate contract object")
		if r.getClient() {
			r.contract, err = des.NewDes(DESContractAddress, r.client)
			if err != nil || r.client == nil {
				log.Error("Failed to communicate with regulator contract", "error", err)
				return isPresent, nil
			}
		} else {
			log.Error("Can't seem to connect to an IPC client")
		}
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
			r.lastUpdated = time.Now()
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
	go r.update()
	return
}

// IsRegulatorPresent checks if regulator is one of the privateFor
func (r *RegulatorClient) IsWhitelisted(enode string) (bool, error) {
	isWhitelisted := false
	var err error
	if r.getClient() {
		log.Trace("Contract is empty so trying to instantiate contract object")
		r.contract, err = des.NewDes(DESContractAddress, r.client)
		if err != nil || r.client == nil {
			log.Error("Failed to communicate with permissions contract", "error", err)
			return isWhitelisted, nil
		}
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
		r.lastUpdated = time.Now()
	}
	go r.update()
	return isWhitelisted, nil
}

// IsDeployed returns the state of the contract, whether it's been deployed yet.
func (r *RegulatorClient) IsDeployed() bool {
	if !r.deployed {
	if r.getClient() {
		code, err := r.client.CodeAt(context.Background(), DESContractAddress, nil)
		if err == nil {
			r.deployed = bytes.Equal(code, common.FromHex(des.DESByteCode))
			log.Debug("Checking whether DES contract has been deployed", "deployed", r.deployed, "codeLength", len(code))
		} else {
			log.Error("Can't find contract code", "error", err)
		}
	} else {
		log.Trace("Can't seem to connect to an IPC client. Maybe it's too early?")
	}
	}
	log.Trace("IsDeployed", "deployed", r.deployed)
	return r.deployed
}

// make sure the maps don't get stagnant, so renew them after ten minutes or so
func (r *RegulatorClient) update() {
	defer func() {
        // recover from panic if one occured.
		if err := recover(); err != nil { //catch
            log.Error("Exception", "error", err)
        }
    }()
	d, err := time.ParseDuration("10m")
	if err == nil {
	// TODO: time should be configurable
	if time.Now().Sub(r.lastUpdated) >  d {
		r.whitelist = make(map[string]bool)
		r.regMap = make(map[string]bool)
	}
    }
}