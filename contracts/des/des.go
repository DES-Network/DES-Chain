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

//go:generate abigen --sol contract/des.sol --pkg des --out contract/des.go

// DESByteCode is the contract's compiled byte code for verification.
const DESByteCode = "0x6080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631b128f61811461009d578063261a323e1461010a578063715018a61461016357806380599e4b1461017a5780638994dd8e146101d35780638da5cb5b1461022c578063a0c15b771461025d578063f2c298be146102b6578063f2fde38b1461030f575b600080fd5b3480156100a957600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100f69436949293602493928401919081908401838280828437509497506103309650505050505050565b604080519115158252519081900360200190f35b34801561011657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100f694369492936024939284019190819084018382808284375094975061039b9650505050505050565b34801561016f57600080fd5b506101786103ce565b005b34801561018657600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261017894369492936024939284019190819084018382808284375094975061043a9650505050505050565b3480156101df57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101789436949293602493928401919081908401838280828437509497506105599650505050505050565b34801561023857600080fd5b50610241610676565b60408051600160a060020a039092168252519081900360200190f35b34801561026957600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101789436949293602493928401919081908401838280828437509497506106859650505050505050565b3480156102c257600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101789436949293602493928401919081908401838280828437509497506107a49650505050505050565b34801561031b57600080fd5b50610178600160a060020a03600435166108c2565b60006001826040518082805190602001908083835b602083106103645780518252601f199092019160209182019101610345565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff16949350505050565b6000600282604051808280519060200190808383602083106103645780518252601f199092019160209182019101610345565b600054600160a060020a031633146103e557600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461045157600080fd5b61045a8161039b565b151561046557600080fd5b60006002826040518082805190602001908083835b602083106104995780518252601f19909201916020918201910161047a565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420805460ff1916951515959095179094555050825183928291908401908083835b602083106105035780518252601f1990920191602091820191016104e4565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507f0f3c2abe8eaf0ca41a7df24a358452cd475b189fe258bb725c444ac1a52184e492506000919050a250565b600054600160a060020a0316331461057057600080fd5b61057981610330565b1561058357600080fd5b600180826040518082805190602001908083835b602083106105b65780518252601f199092019160209182019101610597565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420805460ff1916951515959095179094555050825183928291908401908083835b602083106106205780518252601f199092019160209182019101610601565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507f0a7183d5fa1d84a7c428d9c0a71fb9b95337bbffe2a054b99eda22f4178278bf92506000919050a250565b600054600160a060020a031681565b600054600160a060020a0316331461069c57600080fd5b6106a581610330565b15156106b057600080fd5b60006001826040518082805190602001908083835b602083106106e45780518252601f1990920191602091820191016106c5565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420805460ff1916951515959095179094555050825183928291908401908083835b6020831061074e5780518252601f19909201916020918201910161072f565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507f40a06e475ce2f72df9022ace74c6a720d51ee1539203c4e682ad2c919775f1c392506000919050a250565b600054600160a060020a031633146107bb57600080fd5b6107c48161039b565b156107ce57600080fd5b60016002826040518082805190602001908083835b602083106108025780518252601f1990920191602091820191016107e3565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420805460ff1916951515959095179094555050825183928291908401908083835b6020831061086c5780518252601f19909201916020918201910161084d565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507f22f1f4ff1f21662ed311579542b2058bb9dae330fc8a5b57aa05d1a1f56e078092506000919050a250565b600054600160a060020a031633146108d957600080fd5b6108e2816108e5565b50565b600160a060020a03811615156108fa57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058208f22ed71c74bbc73824c9cca855f20a1741218ed6de13fa8d769c982f0c355380029"


// DesABI is the input ABI used to generate the binding from.
const DesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"nodeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"reg\",\"type\":\"string\"},{\"name\":\"enode\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"RegulatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"RegulatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"enode\",\"type\":\"string\"}],\"name\":\"NodeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

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
