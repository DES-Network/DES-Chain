package private

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"sync"
	// "time"

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
	// client             *ethclient.Client
	ipc = os.Getenv("IPC")
)

// RegulatorClient struct for communicating with the regulator
type RegulatorClient struct {
	contract *des.Des
	client   *ethclient.Client
	events   chan types.Log
	regMap   map[string]bool
	whitelist map[string]bool
}

func getClient() (client *ethclient.Client) {
	once.Do(func() {
		if ipc != "" {
			client, err := ethclient.Dial(ipc)
			if err != nil {
				log.Error("Failed to connect Ethereum client:", "error", err.Error())
			} else {
				log.Trace("Client successfully created: ", "client", client)
			}
		} else {
			log.Error("Failure to get client since IPC url is not set in the environment.")
		}
	})
	return
}

// NewRegulatorClient initializes returns the regulator client
func NewRegulatorClient() (*RegulatorClient, error) {
	//TODO: add check for genesis in association with contract address in init()
	// contract, err := des.NewDes(RegContractAddress, getClient())
	// if err != nil {
	// 	log.Error("Failed to create a Regulator Client", "error", err)
	// 	return nil, err
	// }
	return &RegulatorClient{regMap: make(map[string]bool), whitelist: make(map[string]bool)}, nil
}

// IsRegulatorPresent checks if regulator is one of the privateFor
func (r *RegulatorClient) IsRegulatorPresent(privateFor []string) (bool, error) {
	isPresent := false
	var err error
	// TODO: could make it more efficient
	if r.contract == nil {
		log.Trace("Contract is empty so trying to instantiate contract object")
		client, err := ethclient.Dial(ipc)
		contract, err := des.NewDes(RegContractAddress, client)
		if err != nil || client == nil {
			log.Error("Failed to communicate with regulator contract", "error", err)
			return isPresent, nil
		}
		r.contract = contract
		r.client = client
	}
	for i := range privateFor {
		if isPresent = r.regMap[privateFor[i]]; isPresent {
			break
		}
		isPresent, err = r.contract.Exists(&bind.CallOpts{Context: context.TODO()}, privateFor[i])
		if err != nil {
			log.Error("Couldn't communicate with regulator contract", "error", err)
			break
		} else if isPresent {
			r.regMap[privateFor[i]] = true
			break
		}
	}
	return isPresent, nil
}


// IsRegulatorPresent checks if regulator is one of the privateFor
func (r *RegulatorClient) IsWhitelisted(enode string) (bool, error) {
	isWhitelisted := false
	var err error
	if r.contract == nil {
		log.Trace("Contract is empty so trying to instantiate contract object")
		client, err := ethclient.Dial(ipc)
		contract, err := des.NewDes(RegContractAddress, client)
		if err != nil || client == nil {
			log.Error("Failed to communicate with permissions contract", "error", err)
			return isWhitelisted, nil
		}
		r.contract = contract
		r.client = client
	}

		if isWhitelisted = r.whitelist[enode]; isWhitelisted {
			return isWhitelisted, nil
		}
		isWhitelisted, err = r.contract.NodeExists(&bind.CallOpts{Context: context.TODO()}, enode)
		if err != nil {
			log.Error("Couldn't communicate with regulator contract", "error", err)
		} else if isWhitelisted {
			r.whitelist[enode] = true
		}

	return isWhitelisted, nil
}
