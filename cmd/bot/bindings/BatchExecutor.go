// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BatchExecutorMetaData contains all meta data concerning the BatchExecutor contract.
var BatchExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080806040523461001657610277908161001b8239f35b5f80fdfe608060409080825260049182361015610016575f80fd5b5f92833560e01c63c8d18a451461002b575f80fd5b346101b057816003193601126101b05767ffffffffffffffff81358181116102085761005a903690840161020c565b939091602495863582811161020457610076903690870161020c565b9490918588036101d257508897985b878110610090578880f35b600581901b82810135906001600160a01b03821682036101ce57878310156101bc5799809b9a850135601e19863603018112156101b8578501918235908782116101b45760208094019180360383136101b057838093828c519384928337810182815203925af13d156101ab573d86811161019957875190601f19603f81601f840116011682018281108982111761018757895281528c833d92013e5b1561015857505f19811461014657600101989798610085565b634e487b7160e01b8a5260118752888afd5b87600b8b60649389519362461bcd60e51b85528401528201526a18d85b1b0819985a5b195960aa1b6044820152fd5b634e487b7160e01b8f5260418c528d8ffd5b634e487b7160e01b8d5260418a528b8dfd5b61012d565b8380fd5b8280fd5b5080fd5b634e487b7160e01b8b52603289528b8bfd5b8a80fd5b62461bcd60e51b8152602087820152600f898201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606490fd5b8880fd5b8580fd5b9181601f8401121561023d5782359167ffffffffffffffff831161023d576020808501948460051b01011161023d57565b5f80fdfea2646970667358221220ad0f3bfd42de1340971f1715344856e8af4a00dc2f0893372981cef161d1712864736f6c63430008140033",
}

// BatchExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchExecutorMetaData.ABI instead.
var BatchExecutorABI = BatchExecutorMetaData.ABI

// BatchExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchExecutorMetaData.Bin instead.
var BatchExecutorBin = BatchExecutorMetaData.Bin

// DeployBatchExecutor deploys a new Ethereum contract, binding an instance of BatchExecutor to it.
func DeployBatchExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchExecutor, error) {
	parsed, err := BatchExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchExecutor{BatchExecutorCaller: BatchExecutorCaller{contract: contract}, BatchExecutorTransactor: BatchExecutorTransactor{contract: contract}, BatchExecutorFilterer: BatchExecutorFilterer{contract: contract}}, nil
}

// BatchExecutor is an auto generated Go binding around an Ethereum contract.
type BatchExecutor struct {
	BatchExecutorCaller     // Read-only binding to the contract
	BatchExecutorTransactor // Write-only binding to the contract
	BatchExecutorFilterer   // Log filterer for contract events
}

// BatchExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchExecutorSession struct {
	Contract     *BatchExecutor    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchExecutorCallerSession struct {
	Contract *BatchExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BatchExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchExecutorTransactorSession struct {
	Contract     *BatchExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BatchExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchExecutorRaw struct {
	Contract *BatchExecutor // Generic contract binding to access the raw methods on
}

// BatchExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchExecutorCallerRaw struct {
	Contract *BatchExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// BatchExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchExecutorTransactorRaw struct {
	Contract *BatchExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchExecutor creates a new instance of BatchExecutor, bound to a specific deployed contract.
func NewBatchExecutor(address common.Address, backend bind.ContractBackend) (*BatchExecutor, error) {
	contract, err := bindBatchExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchExecutor{BatchExecutorCaller: BatchExecutorCaller{contract: contract}, BatchExecutorTransactor: BatchExecutorTransactor{contract: contract}, BatchExecutorFilterer: BatchExecutorFilterer{contract: contract}}, nil
}

// NewBatchExecutorCaller creates a new read-only instance of BatchExecutor, bound to a specific deployed contract.
func NewBatchExecutorCaller(address common.Address, caller bind.ContractCaller) (*BatchExecutorCaller, error) {
	contract, err := bindBatchExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchExecutorCaller{contract: contract}, nil
}

// NewBatchExecutorTransactor creates a new write-only instance of BatchExecutor, bound to a specific deployed contract.
func NewBatchExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchExecutorTransactor, error) {
	contract, err := bindBatchExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchExecutorTransactor{contract: contract}, nil
}

// NewBatchExecutorFilterer creates a new log filterer instance of BatchExecutor, bound to a specific deployed contract.
func NewBatchExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchExecutorFilterer, error) {
	contract, err := bindBatchExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchExecutorFilterer{contract: contract}, nil
}

// bindBatchExecutor binds a generic wrapper to an already deployed contract.
func bindBatchExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchExecutor *BatchExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchExecutor.Contract.BatchExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchExecutor *BatchExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchExecutor.Contract.BatchExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchExecutor *BatchExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchExecutor.Contract.BatchExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchExecutor *BatchExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchExecutor *BatchExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchExecutor *BatchExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchExecutor.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0xc8d18a45.
//
// Solidity: function execute(address[] targets, bytes[] data) returns()
func (_BatchExecutor *BatchExecutorTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, data [][]byte) (*types.Transaction, error) {
	return _BatchExecutor.contract.Transact(opts, "execute", targets, data)
}

// Execute is a paid mutator transaction binding the contract method 0xc8d18a45.
//
// Solidity: function execute(address[] targets, bytes[] data) returns()
func (_BatchExecutor *BatchExecutorSession) Execute(targets []common.Address, data [][]byte) (*types.Transaction, error) {
	return _BatchExecutor.Contract.Execute(&_BatchExecutor.TransactOpts, targets, data)
}

// Execute is a paid mutator transaction binding the contract method 0xc8d18a45.
//
// Solidity: function execute(address[] targets, bytes[] data) returns()
func (_BatchExecutor *BatchExecutorTransactorSession) Execute(targets []common.Address, data [][]byte) (*types.Transaction, error) {
	return _BatchExecutor.Contract.Execute(&_BatchExecutor.TransactOpts, targets, data)
}
