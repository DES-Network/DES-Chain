// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package des

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DesABI is the input ABI used to generate the binding from.
const DesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"nodeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"RegulatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"RegulatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"NodeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

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
