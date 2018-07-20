package private

import (
	"errors"
	"sync"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/contracts/regulators"
)

var (
	RegContractAddress = common.HexToAddress("0x314159265dD8dbb310642f98f50C066173C1259b")
	ErrNoRegulator = errors.New("invalid transaction because no regulator provided")
	once           sync.Once
	client *ethclient.Client
	ipc = os.Getenv("IPC")
)

// RegulatorClient struct for communicating with the regulator
type RegulatorClient struct {
	contract *regulators.Regulators
	client   *ethclient.Client
	events   chan types.Log
	regMap	 map[string]bool
}

func getClient() (*ethclient.Client) {
	once.Do(func() {
		if ipc != "" {
		client, err := ethclient.Dial(ipc)
		if err != nil {
			log.Error("Failed to connect Ethereum client", "error", err.Error())
		} else {
			log.Trace("Client successfully created: ", client)
		}
	} else {
		log.Error("Failure to get client since IPC url is not set in the environment.")
	}
	})
	return client
}

// NewRegulatorClient initializes returns the regulator client
func NewRegulatorClient() (*RegulatorClient, error) {
	//TODO: add check for genesis in association with contract address
	contract, err := regulators.NewRegulators(RegContractAddress, getClient())
	if err != nil {
		log.Error("%v", err)
		return nil, err
	}
	return &RegulatorClient{client: getClient(), regMap: make(map[string]bool), contract: contract}, nil
}

// IsRegulatorPresent checks if regulator is one of the privateFor
func (r *RegulatorClient) IsRegulatorPresent(privateFor []string) (bool, error) {
	isPresent := false
	var err error
	for i := range privateFor {
		if isPresent = r.regMap[privateFor[i]]; isPresent {
			break
		}
		isPresent, err = r.contract.Exists(nil, privateFor[i])
		if err != nil {
			log.Error("%v: %v\n", "Couldn't communicate with regulator contract with the following error", err)
			break
		} else if isPresent {
			r.regMap[privateFor[i]] = true
			break
		}
	}
	return isPresent, nil
}
