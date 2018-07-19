// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package regulators

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RegulatorsABI is the input ABI used to generate the binding from.
const RegulatorsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"RegulatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"RegulatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Regulators is an auto generated Go binding around an Ethereum contract.
type Regulators struct {
	RegulatorsCaller     // Read-only binding to the contract
	RegulatorsTransactor // Write-only binding to the contract
}

// RegulatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegulatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegulatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegulatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegulatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegulatorsSession struct {
	Contract     *Regulators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegulatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegulatorsCallerSession struct {
	Contract *RegulatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RegulatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegulatorsTransactorSession struct {
	Contract     *RegulatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RegulatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegulatorsRaw struct {
	Contract *Regulators // Generic contract binding to access the raw methods on
}

// RegulatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegulatorsCallerRaw struct {
	Contract *RegulatorsCaller // Generic read-only contract binding to access the raw methods on
}

// RegulatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegulatorsTransactorRaw struct {
	Contract *RegulatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegulators creates a new instance of Regulators, bound to a specific deployed contract.
func NewRegulators(address common.Address, backend bind.ContractBackend) (*Regulators, error) {
	contract, err := bindRegulators(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Regulators{RegulatorsCaller: RegulatorsCaller{contract: contract}, RegulatorsTransactor: RegulatorsTransactor{contract: contract}}, nil
}

// NewRegulatorsCaller creates a new read-only instance of Regulators, bound to a specific deployed contract.
func NewRegulatorsCaller(address common.Address, caller bind.ContractCaller) (*RegulatorsCaller, error) {
	contract, err := bindRegulators(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RegulatorsCaller{contract: contract}, nil
}

// NewRegulatorsTransactor creates a new write-only instance of Regulators, bound to a specific deployed contract.
func NewRegulatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*RegulatorsTransactor, error) {
	contract, err := bindRegulators(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RegulatorsTransactor{contract: contract}, nil
}

// bindRegulators binds a generic wrapper to an already deployed contract.
func bindRegulators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegulatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Regulators *RegulatorsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Regulators.Contract.RegulatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Regulators *RegulatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Regulators.Contract.RegulatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Regulators *RegulatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Regulators.Contract.RegulatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Regulators *RegulatorsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Regulators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Regulators *RegulatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Regulators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Regulators *RegulatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Regulators.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(_address string) constant returns(bool)
func (_Regulators *RegulatorsCaller) Exists(opts *bind.CallOpts, _address string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Regulators.contract.Call(opts, out, "exists", _address)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(_address string) constant returns(bool)
func (_Regulators *RegulatorsSession) Exists(_address string) (bool, error) {
	return _Regulators.Contract.Exists(&_Regulators.CallOpts, _address)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(_address string) constant returns(bool)
func (_Regulators *RegulatorsCallerSession) Exists(_address string) (bool, error) {
	return _Regulators.Contract.Exists(&_Regulators.CallOpts, _address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Regulators *RegulatorsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Regulators.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Regulators *RegulatorsSession) Owner() (common.Address, error) {
	return _Regulators.Contract.Owner(&_Regulators.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Regulators *RegulatorsCallerSession) Owner() (common.Address, error) {
	return _Regulators.Contract.Owner(&_Regulators.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_address string) returns()
func (_Regulators *RegulatorsTransactor) Register(opts *bind.TransactOpts, _address string) (*types.Transaction, error) {
	return _Regulators.contract.Transact(opts, "register", _address)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_address string) returns()
func (_Regulators *RegulatorsSession) Register(_address string) (*types.Transaction, error) {
	return _Regulators.Contract.Register(&_Regulators.TransactOpts, _address)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(_address string) returns()
func (_Regulators *RegulatorsTransactorSession) Register(_address string) (*types.Transaction, error) {
	return _Regulators.Contract.Register(&_Regulators.TransactOpts, _address)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(_address string) returns()
func (_Regulators *RegulatorsTransactor) Remove(opts *bind.TransactOpts, _address string) (*types.Transaction, error) {
	return _Regulators.contract.Transact(opts, "remove", _address)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(_address string) returns()
func (_Regulators *RegulatorsSession) Remove(_address string) (*types.Transaction, error) {
	return _Regulators.Contract.Remove(&_Regulators.TransactOpts, _address)
}

// Remove is a paid mutator transaction binding the contract method 0x80599e4b.
//
// Solidity: function remove(_address string) returns()
func (_Regulators *RegulatorsTransactorSession) Remove(_address string) (*types.Transaction, error) {
	return _Regulators.Contract.Remove(&_Regulators.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Regulators *RegulatorsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Regulators.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Regulators *RegulatorsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Regulators.Contract.RenounceOwnership(&_Regulators.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Regulators *RegulatorsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Regulators.Contract.RenounceOwnership(&_Regulators.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Regulators *RegulatorsTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Regulators.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Regulators *RegulatorsSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Regulators.Contract.TransferOwnership(&_Regulators.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Regulators *RegulatorsTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Regulators.Contract.TransferOwnership(&_Regulators.TransactOpts, _newOwner)
}
