// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissioning

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PermissioningABI is the input ABI used to generate the binding from.
const PermissioningABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeName\",\"type\":\"bytes\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeName\",\"type\":\"bytes\"}],\"name\":\"NodeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Permissioning is an auto generated Go binding around an Ethereum contract.
type Permissioning struct {
	PermissioningCaller     // Read-only binding to the contract
	PermissioningTransactor // Write-only binding to the contract
}

// PermissioningCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissioningCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissioningTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissioningTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissioningSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissioningSession struct {
	Contract     *Permissioning    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissioningCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissioningCallerSession struct {
	Contract *PermissioningCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PermissioningTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissioningTransactorSession struct {
	Contract     *PermissioningTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PermissioningRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissioningRaw struct {
	Contract *Permissioning // Generic contract binding to access the raw methods on
}

// PermissioningCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissioningCallerRaw struct {
	Contract *PermissioningCaller // Generic read-only contract binding to access the raw methods on
}

// PermissioningTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissioningTransactorRaw struct {
	Contract *PermissioningTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissioning creates a new instance of Permissioning, bound to a specific deployed contract.
func NewPermissioning(address common.Address, backend bind.ContractBackend) (*Permissioning, error) {
	contract, err := bindPermissioning(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Permissioning{PermissioningCaller: PermissioningCaller{contract: contract}, PermissioningTransactor: PermissioningTransactor{contract: contract}}, nil
}

// NewPermissioningCaller creates a new read-only instance of Permissioning, bound to a specific deployed contract.
func NewPermissioningCaller(address common.Address, caller bind.ContractCaller) (*PermissioningCaller, error) {
	contract, err := bindPermissioning(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PermissioningCaller{contract: contract}, nil
}

// NewPermissioningTransactor creates a new write-only instance of Permissioning, bound to a specific deployed contract.
func NewPermissioningTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissioningTransactor, error) {
	contract, err := bindPermissioning(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PermissioningTransactor{contract: contract}, nil
}

// bindPermissioning binds a generic wrapper to an already deployed contract.
func bindPermissioning(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissioningABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissioning *PermissioningRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissioning.Contract.PermissioningCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissioning *PermissioningRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissioning.Contract.PermissioningTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissioning *PermissioningRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissioning.Contract.PermissioningTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissioning *PermissioningCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissioning.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissioning *PermissioningTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissioning.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissioning *PermissioningTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissioning.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(name bytes) constant returns(bool)
func (_Permissioning *PermissioningCaller) Exists(opts *bind.CallOpts, name []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "exists", name)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(name bytes) constant returns(bool)
func (_Permissioning *PermissioningSession) Exists(name []byte) (bool, error) {
	return _Permissioning.Contract.Exists(&_Permissioning.CallOpts, name)
}

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(name bytes) constant returns(bool)
func (_Permissioning *PermissioningCallerSession) Exists(name []byte) (bool, error) {
	return _Permissioning.Contract.Exists(&_Permissioning.CallOpts, name)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Permissioning *PermissioningCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Permissioning.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Permissioning *PermissioningSession) Owner() (common.Address, error) {
	return _Permissioning.Contract.Owner(&_Permissioning.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Permissioning *PermissioningCallerSession) Owner() (common.Address, error) {
	return _Permissioning.Contract.Owner(&_Permissioning.CallOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0x7e10c84f.
//
// Solidity: function addNode(name bytes) returns()
func (_Permissioning *PermissioningTransactor) AddNode(opts *bind.TransactOpts, name []byte) (*types.Transaction, error) {
	return _Permissioning.contract.Transact(opts, "addNode", name)
}

// AddNode is a paid mutator transaction binding the contract method 0x7e10c84f.
//
// Solidity: function addNode(name bytes) returns()
func (_Permissioning *PermissioningSession) AddNode(name []byte) (*types.Transaction, error) {
	return _Permissioning.Contract.AddNode(&_Permissioning.TransactOpts, name)
}

// AddNode is a paid mutator transaction binding the contract method 0x7e10c84f.
//
// Solidity: function addNode(name bytes) returns()
func (_Permissioning *PermissioningTransactorSession) AddNode(name []byte) (*types.Transaction, error) {
	return _Permissioning.Contract.AddNode(&_Permissioning.TransactOpts, name)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xc8a2639c.
//
// Solidity: function deleteNode(name bytes) returns()
func (_Permissioning *PermissioningTransactor) DeleteNode(opts *bind.TransactOpts, name []byte) (*types.Transaction, error) {
	return _Permissioning.contract.Transact(opts, "deleteNode", name)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xc8a2639c.
//
// Solidity: function deleteNode(name bytes) returns()
func (_Permissioning *PermissioningSession) DeleteNode(name []byte) (*types.Transaction, error) {
	return _Permissioning.Contract.DeleteNode(&_Permissioning.TransactOpts, name)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xc8a2639c.
//
// Solidity: function deleteNode(name bytes) returns()
func (_Permissioning *PermissioningTransactorSession) DeleteNode(name []byte) (*types.Transaction, error) {
	return _Permissioning.Contract.DeleteNode(&_Permissioning.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Permissioning *PermissioningTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissioning.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Permissioning *PermissioningSession) RenounceOwnership() (*types.Transaction, error) {
	return _Permissioning.Contract.RenounceOwnership(&_Permissioning.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Permissioning *PermissioningTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Permissioning.Contract.RenounceOwnership(&_Permissioning.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Permissioning *PermissioningTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Permissioning.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Permissioning *PermissioningSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Permissioning.Contract.TransferOwnership(&_Permissioning.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Permissioning *PermissioningTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Permissioning.Contract.TransferOwnership(&_Permissioning.TransactOpts, _newOwner)
}
