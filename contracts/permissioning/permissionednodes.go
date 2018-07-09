// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/block360/quorum/accounts/abi"
	"github.com/block360/quorum/accounts/abi/bind"
	"github.com/block360/quorum/common"
	"github.com/block360/quorum/core/types"
)

// PermissionedNodesABI is the input ABI used to generate the binding from.
const PermissionedNodesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeName\",\"type\":\"bytes\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeName\",\"type\":\"bytes\"}],\"name\":\"NodeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PermissionedNodes is an auto generated Go binding around an Ethereum contract.
type PermissionedNodes struct {
	PermissionedNodesCaller     // Read-only binding to the contract
	PermissionedNodesTransactor // Write-only binding to the contract
}

// PermissionedNodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionedNodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedNodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionedNodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedNodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionedNodesSession struct {
	Contract     *PermissionedNodes // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PermissionedNodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionedNodesCallerSession struct {
	Contract *PermissionedNodesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PermissionedNodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionedNodesTransactorSession struct {
	Contract     *PermissionedNodesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PermissionedNodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionedNodesRaw struct {
	Contract *PermissionedNodes // Generic contract binding to access the raw methods on
}

// PermissionedNodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionedNodesCallerRaw struct {
	Contract *PermissionedNodesCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionedNodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionedNodesTransactorRaw struct {
	Contract *PermissionedNodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionedNodes creates a new instance of PermissionedNodes, bound to a specific deployed contract.
func NewPermissionedNodes(address common.Address, backend bind.ContractBackend) (*PermissionedNodes, error) {
	contract, err := bindPermissionedNodes(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodes{PermissionedNodesCaller: PermissionedNodesCaller{contract: contract}, PermissionedNodesTransactor: PermissionedNodesTransactor{contract: contract}}, nil
}

// NewPermissionedNodesCaller creates a new read-only instance of PermissionedNodes, bound to a specific deployed contract.
func NewPermissionedNodesCaller(address common.Address, caller bind.ContractCaller) (*PermissionedNodesCaller, error) {
	contract, err := bindPermissionedNodes(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodesCaller{contract: contract}, nil
}

// NewPermissionedNodesTransactor creates a new write-only instance of PermissionedNodes, bound to a specific deployed contract.
func NewPermissionedNodesTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionedNodesTransactor, error) {
	contract, err := bindPermissionedNodes(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodesTransactor{contract: contract}, nil
}

// bindPermissionedNodes binds a generic wrapper to an already deployed contract.
func bindPermissionedNodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionedNodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedNodes *PermissionedNodesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermissionedNodes.Contract.PermissionedNodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedNodes *PermissionedNodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.PermissionedNodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedNodes *PermissionedNodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.PermissionedNodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedNodes *PermissionedNodesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermissionedNodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedNodes *PermissionedNodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedNodes *PermissionedNodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(name bytes) constant returns(bool)
func (_PermissionedNodes *PermissionedNodesCaller) Exists(opts *bind.CallOpts, name []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermissionedNodes.contract.Call(opts, out, "exists", name)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(name bytes) constant returns(bool)
func (_PermissionedNodes *PermissionedNodesSession) Exists(name []byte) (bool, error) {
	return _PermissionedNodes.Contract.Exists(&_PermissionedNodes.CallOpts, name)
}

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(name bytes) constant returns(bool)
func (_PermissionedNodes *PermissionedNodesCallerSession) Exists(name []byte) (bool, error) {
	return _PermissionedNodes.Contract.Exists(&_PermissionedNodes.CallOpts, name)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PermissionedNodes *PermissionedNodesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermissionedNodes.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PermissionedNodes *PermissionedNodesSession) Owner() (common.Address, error) {
	return _PermissionedNodes.Contract.Owner(&_PermissionedNodes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PermissionedNodes *PermissionedNodesCallerSession) Owner() (common.Address, error) {
	return _PermissionedNodes.Contract.Owner(&_PermissionedNodes.CallOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0x7e10c84f.
//
// Solidity: function addNode(name bytes) returns()
func (_PermissionedNodes *PermissionedNodesTransactor) AddNode(opts *bind.TransactOpts, name []byte) (*types.Transaction, error) {
	return _PermissionedNodes.contract.Transact(opts, "addNode", name)
}

// AddNode is a paid mutator transaction binding the contract method 0x7e10c84f.
//
// Solidity: function addNode(name bytes) returns()
func (_PermissionedNodes *PermissionedNodesSession) AddNode(name []byte) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.AddNode(&_PermissionedNodes.TransactOpts, name)
}

// AddNode is a paid mutator transaction binding the contract method 0x7e10c84f.
//
// Solidity: function addNode(name bytes) returns()
func (_PermissionedNodes *PermissionedNodesTransactorSession) AddNode(name []byte) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.AddNode(&_PermissionedNodes.TransactOpts, name)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xc8a2639c.
//
// Solidity: function deleteNode(name bytes) returns()
func (_PermissionedNodes *PermissionedNodesTransactor) DeleteNode(opts *bind.TransactOpts, name []byte) (*types.Transaction, error) {
	return _PermissionedNodes.contract.Transact(opts, "deleteNode", name)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xc8a2639c.
//
// Solidity: function deleteNode(name bytes) returns()
func (_PermissionedNodes *PermissionedNodesSession) DeleteNode(name []byte) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.DeleteNode(&_PermissionedNodes.TransactOpts, name)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xc8a2639c.
//
// Solidity: function deleteNode(name bytes) returns()
func (_PermissionedNodes *PermissionedNodesTransactorSession) DeleteNode(name []byte) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.DeleteNode(&_PermissionedNodes.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PermissionedNodes *PermissionedNodesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedNodes.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PermissionedNodes *PermissionedNodesSession) RenounceOwnership() (*types.Transaction, error) {
	return _PermissionedNodes.Contract.RenounceOwnership(&_PermissionedNodes.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PermissionedNodes *PermissionedNodesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PermissionedNodes.Contract.RenounceOwnership(&_PermissionedNodes.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PermissionedNodes *PermissionedNodesTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PermissionedNodes.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PermissionedNodes *PermissionedNodesSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.TransferOwnership(&_PermissionedNodes.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PermissionedNodes *PermissionedNodesTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PermissionedNodes.Contract.TransferOwnership(&_PermissionedNodes.TransactOpts, _newOwner)
}
