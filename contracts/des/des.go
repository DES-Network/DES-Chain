// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package des

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DesABI is the input ABI used to generate the binding from.
const DesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"nodeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"enodes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"reg\",\"type\":\"string\"},{\"name\":\"enode\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"RegulatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"RegulatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"NodeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DESByteCode is the compiled bytecode
const DESByteCode = "0x6080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631b128f61146100a9578063261a323e1461012a578063715018a6146101ab57806380599e4b146101c25780638994dd8e1461022b5780638da5cb5b146102945780639890b329146102eb578063a0c15b7714610391578063f2c298be146103fa578063f2fde38b14610463575b600080fd5b3480156100b557600080fd5b50610110600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506104a6565b604051808215151515815260200191505060405180910390f35b34801561013657600080fd5b50610191600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610528565b604051808215151515815260200191505060405180910390f35b3480156101b757600080fd5b506101c06105aa565b005b3480156101ce57600080fd5b50610229600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506106ac565b005b34801561023757600080fd5b50610292600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610831565b005b3480156102a057600080fd5b506102a96109c8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102f757600080fd5b50610316600480360381019080803590602001909291905050506109ed565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561035657808201518184015260208101905061033b565b50505050905090810190601f1680156103835780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561039d57600080fd5b506103f8600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610aa8565b005b34801561040657600080fd5b50610461600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610c2d565b005b34801561046f57600080fd5b506104a4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610dc5565b005b60006001826040518082805190602001908083835b6020831015156104e057805182526020820191506020810190506020830392506104bb565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff169050919050565b60006004826040518082805190602001908083835b602083101515610562578051825260208201915060208101905060208303925061053d565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561060557600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561070757600080fd5b61071081610528565b151561071b57600080fd5b60006004826040518082805190602001908083835b6020831015156107555780518252602082019150602081019050602083039250610730565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff021916908315150217905550806040518082805190602001908083835b6020831015156107d557805182526020820191506020810190506020830392506107b0565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f0f3c2abe8eaf0ca41a7df24a358452cd475b189fe258bb725c444ac1a52184e460405160405180910390a250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561088c57600080fd5b610895816104a6565b1515156108a157600080fd5b600180826040518082805190602001908083835b6020831015156108da57805182526020820191506020810190506020830392506108b5565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff021916908315150217905550600360008154809291906001019190505550806040518082805190602001908083835b60208310151561096c5780518252602082019150602081019050602083039250610947565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f0a7183d5fa1d84a7c428d9c0a71fb9b95337bbffe2a054b99eda22f4178278bf60405160405180910390a250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6002818154811015156109fc57fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610aa05780601f10610a7557610100808354040283529160200191610aa0565b820191906000526020600020905b815481529060010190602001808311610a8357829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b0357600080fd5b610b0c816104a6565b1515610b1757600080fd5b60006001826040518082805190602001908083835b602083101515610b515780518252602082019150602081019050602083039250610b2c565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff021916908315150217905550806040518082805190602001908083835b602083101515610bd15780518252602082019150602081019050602083039250610bac565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f40a06e475ce2f72df9022ace74c6a720d51ee1539203c4e682ad2c919775f1c360405160405180910390a250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c8857600080fd5b610c9181610528565b151515610c9d57600080fd5b60016004826040518082805190602001908083835b602083101515610cd75780518252602082019150602081019050602083039250610cb2565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff021916908315150217905550600560008154809291906001019190505550806040518082805190602001908083835b602083101515610d695780518252602082019150602081019050602083039250610d44565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f22f1f4ff1f21662ed311579542b2058bb9dae330fc8a5b57aa05d1a1f56e078060405160405180910390a250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e2057600080fd5b610e2981610e2c565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610e6857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058209c0c88c31a23eb709f6aa353a5073fd15c2ce352b2a60ff4562d92f5ae78e0990029"

// Des is an auto generated Go binding around an Ethereum contract.
type Des struct {
	DesCaller     // Read-only binding to the contract
	DesTransactor // Write-only binding to the contract
}

// DesCaller is an auto generated read-only Go binding around an Ethereum contract.
type DesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DesSession struct {
	Contract     *Des              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DesCallerSession struct {
	Contract *DesCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DesTransactorSession struct {
	Contract     *DesTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DesRaw is an auto generated low-level Go binding around an Ethereum contract.
type DesRaw struct {
	Contract *Des // Generic contract binding to access the raw methods on
}

// DesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DesCallerRaw struct {
	Contract *DesCaller // Generic read-only contract binding to access the raw methods on
}

// DesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DesTransactorRaw struct {
	Contract *DesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDes creates a new instance of Des, bound to a specific deployed contract.
func NewDes(address common.Address, backend bind.ContractBackend) (*Des, error) {
	contract, err := bindDes(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Des{DesCaller: DesCaller{contract: contract}, DesTransactor: DesTransactor{contract: contract}}, nil
}

// NewDesCaller creates a new read-only instance of Des, bound to a specific deployed contract.
func NewDesCaller(address common.Address, caller bind.ContractCaller) (*DesCaller, error) {
	contract, err := bindDes(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DesCaller{contract: contract}, nil
}

// NewDesTransactor creates a new write-only instance of Des, bound to a specific deployed contract.
func NewDesTransactor(address common.Address, transactor bind.ContractTransactor) (*DesTransactor, error) {
	contract, err := bindDes(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DesTransactor{contract: contract}, nil
}

// bindDes binds a generic wrapper to an already deployed contract.
func bindDes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Des *DesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Des.Contract.DesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Des *DesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Des.Contract.DesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Des *DesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Des.Contract.DesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Des *DesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Des.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Des *DesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Des.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Des *DesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Des.Contract.contract.Transact(opts, method, params...)
}

// Enodes is a free data retrieval call binding the contract method 0x9890b329.
//
// Solidity: function enodes( uint256) constant returns(string)
func (_Des *DesCaller) Enodes(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Des.contract.Call(opts, out, "enodes", arg0)
	return *ret0, err
}

// Enodes is a free data retrieval call binding the contract method 0x9890b329.
//
// Solidity: function enodes( uint256) constant returns(string)
func (_Des *DesSession) Enodes(arg0 *big.Int) (string, error) {
	return _Des.Contract.Enodes(&_Des.CallOpts, arg0)
}

// Enodes is a free data retrieval call binding the contract method 0x9890b329.
//
// Solidity: function enodes( uint256) constant returns(string)
func (_Des *DesCallerSession) Enodes(arg0 *big.Int) (string, error) {
	return _Des.Contract.Enodes(&_Des.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(_address string) constant returns(bool)
func (_Des *DesCaller) Exists(opts *bind.CallOpts, _address string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Des.contract.Call(opts, out, "exists", _address)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(_address string) constant returns(bool)
func (_Des *DesSession) Exists(_address string) (bool, error) {
	return _Des.Contract.Exists(&_Des.CallOpts, _address)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(_address string) constant returns(bool)
func (_Des *DesCallerSession) Exists(_address string) (bool, error) {
	return _Des.Contract.Exists(&_Des.CallOpts, _address)
}

// NodeExists is a free data retrieval call binding the contract method 0x1b128f61.
//
// Solidity: function nodeExists(enode string) constant returns(bool)
func (_Des *DesCaller) NodeExists(opts *bind.CallOpts, enode string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Des.contract.Call(opts, out, "nodeExists", enode)
	return *ret0, err
}

// NodeExists is a free data retrieval call binding the contract method 0x1b128f61.
//
// Solidity: function nodeExists(enode string) constant returns(bool)
func (_Des *DesSession) NodeExists(enode string) (bool, error) {
	return _Des.Contract.NodeExists(&_Des.CallOpts, enode)
}

// NodeExists is a free data retrieval call binding the contract method 0x1b128f61.
//
// Solidity: function nodeExists(enode string) constant returns(bool)
func (_Des *DesCallerSession) NodeExists(enode string) (bool, error) {
	return _Des.Contract.NodeExists(&_Des.CallOpts, enode)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Des *DesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Des.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Des *DesSession) Owner() (common.Address, error) {
	return _Des.Contract.Owner(&_Des.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Des *DesCallerSession) Owner() (common.Address, error) {
	return _Des.Contract.Owner(&_Des.CallOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0x8994dd8e.
//
// Solidity: function addNode(enode string) returns()
func (_Des *DesTransactor) AddNode(opts *bind.TransactOpts, enode string) (*types.Transaction, error) {
	return _Des.contract.Transact(opts, "addNode", enode)
}

// AddNode is a paid mutator transaction binding the contract method 0x8994dd8e.
//
// Solidity: function addNode(enode string) returns()
func (_Des *DesSession) AddNode(enode string) (*types.Transaction, error) {
	return _Des.Contract.AddNode(&_Des.TransactOpts, enode)
}

// AddNode is a paid mutator transaction binding the contract method 0x8994dd8e.
//
// Solidity: function addNode(enode string) returns()
func (_Des *DesTransactorSession) AddNode(enode string) (*types.Transaction, error) {
	return _Des.Contract.AddNode(&_Des.TransactOpts, enode)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(enode string) returns()
func (_Des *DesTransactor) DeleteNode(opts *bind.TransactOpts, enode string) (*types.Transaction, error) {
	return _Des.contract.Transact(opts, "deleteNode", enode)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(enode string) returns()
func (_Des *DesSession) DeleteNode(enode string) (*types.Transaction, error) {
	return _Des.Contract.DeleteNode(&_Des.TransactOpts, enode)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(enode string) returns()
func (_Des *DesTransactorSession) DeleteNode(enode string) (*types.Transaction, error) {
	return _Des.Contract.DeleteNode(&_Des.TransactOpts, enode)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_address string) returns()
func (_Des *DesTransactor) Register(opts *bind.TransactOpts, _address string) (*types.Transaction, error) {
	return _Des.contract.Transact(opts, "register", _address)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_address string) returns()
func (_Des *DesSession) Register(_address string) (*types.Transaction, error) {
	return _Des.Contract.Register(&_Des.TransactOpts, _address)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_address string) returns()
func (_Des *DesTransactorSession) Register(_address string) (*types.Transaction, error) {
	return _Des.Contract.Register(&_Des.TransactOpts, _address)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(_address string) returns()
func (_Des *DesTransactor) Remove(opts *bind.TransactOpts, _address string) (*types.Transaction, error) {
	return _Des.contract.Transact(opts, "remove", _address)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(_address string) returns()
func (_Des *DesSession) Remove(_address string) (*types.Transaction, error) {
	return _Des.Contract.Remove(&_Des.TransactOpts, _address)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(_address string) returns()
func (_Des *DesTransactorSession) Remove(_address string) (*types.Transaction, error) {
	return _Des.Contract.Remove(&_Des.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Des *DesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Des.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Des *DesSession) RenounceOwnership() (*types.Transaction, error) {
	return _Des.Contract.RenounceOwnership(&_Des.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Des *DesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Des.Contract.RenounceOwnership(&_Des.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Des *DesTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Des.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Des *DesSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Des.Contract.TransferOwnership(&_Des.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Des *DesTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Des.Contract.TransferOwnership(&_Des.TransactOpts, _newOwner)
}
