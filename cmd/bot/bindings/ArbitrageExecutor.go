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

// ArbitrageExecutorMetaData contains all meta data concerning the ArbitrageExecutor contract.
var ArbitrageExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"pairA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pairB\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"step\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TradeExecuted\",\"inputs\":[{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"profit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080806040523461001657610893908161001b8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f90813560e01c636142251c14610027575f80fd5b3461064e57608036600319011261064e576004356001600160a01b03811681036103de57602435916001600160a01b03831683036103ed57630240bc6b60e21b8082526060939084836004816001600160a01b0388165afa9485156104205786938796610628575b5060405191825280826004816001600160a01b0387165afa91821561061d57879188936105ea575b5050604051630dfe168160e01b81526020816004816001600160a01b038a165afa9081156104415788916105cb575b5060405163d21220a760e01b8152906020826004816001600160a01b0389165afa80156105a15761012c928a91610582575b506001600160a01b03918216911614610725565b60405163d21220a760e01b81526020816004816001600160a01b038a165afa9081156104415788916105ac575b50604051630dfe168160e01b8152906020826004816001600160a01b0389165afa80156105a15761019e928a9161058257506001600160a01b03918216911614610725565b86948594606435604435815b818111156105095750505085156104d857604051630dfe168160e01b8152899890916020836004816001600160a01b0388165afa9283156104cd578a936104ac575b5060405163d21220a760e01b81526020816004816001600160a01b0389165afa93841561046f578b91829561047a575b506040516323b872dd60e01b81523360048201526001600160a01b038781166024830152604482018d90529092602092849260649284929091165af1801561046f5761027c936001600160701b03928392610450575b50169116896107fd565b916040516102898161067a565b8981526001600160a01b0382163b1561044c5789916102c0916040519c8d8094819363022c0d9f60e01b8352308a60048501610779565b03926001600160a01b03165af180156104415761042b575b60405163a9059cbb60e01b81526001600160a01b038681166004830152602482018490529899509697959689969091602091839160449183918b91165af180156104205761033a946001600160701b039283926103f1575b50169216906107fd565b6040516103468161067a565b8381526001600160a01b0383163b156103ed5761037e9284928360405180968195829463022c0d9f60e01b8452339060048501610779565b03926001600160a01b03165af180156103e2576103ca575b507f589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec70866040848482519182526020820152a180f35b6103d390610652565b6103de57825f610396565b8280fd5b6040513d84823e3d90fd5b8380fd5b6104129060203d602011610419575b61040a8183610696565b810190610761565b505f610330565b503d610400565b6040513d88823e3d90fd5b949596610439602099610652565b9695946102d8565b6040513d8a823e3d90fd5b8980fd5b6104689060203d6020116104195761040a8183610696565b505f610272565b6040513d8d823e3d90fd5b602091955061049e90823d84116104a5575b6104968183610696565b810190610706565b949061021c565b503d61048c565b6104c691935060203d6020116104a5576104968183610696565b915f6101ec565b6040513d8c823e3d90fd5b60405162461bcd60e51b81526020600482015260096024820152681b9bc81c1c9bd99a5d60ba1b6044820152606490fd5b826105316001600160701b038d61052c828c169280808d169316908a16876107fd565b6107fd565b828111610548575b50610543916107dc565b6101aa565b8281039150811161056e57898111610562575b8390610539565b9099509750888261055b565b634e487b7160e01b8d52601160045260248dfd5b61059b915060203d6020116104a5576104968183610696565b5f610118565b6040513d8b823e3d90fd5b6105c5915060203d6020116104a5576104968183610696565b5f610159565b6105e4915060203d6020116104a5576104968183610696565b5f6100e6565b61060d935080919250903d10610616575b6106058183610696565b8101906106d0565b50905f806100b7565b503d6105fb565b6040513d89823e3d90fd5b908096506106439294503d8711610616576106058183610696565b50929092945f61008f565b5080fd5b67ffffffffffffffff811161066657604052565b634e487b7160e01b5f52604160045260245ffd5b6020810190811067ffffffffffffffff82111761066657604052565b90601f8019910116810190811067ffffffffffffffff82111761066657604052565b51906001600160701b03821682036106cc57565b5f80fd5b908160609103126106cc576106e4816106b8565b9160406106f3602084016106b8565b92015163ffffffff811681036106cc5790565b908160209103126106cc57516001600160a01b03811681036106cc5790565b1561072c57565b60405162461bcd60e51b815260206004820152600d60248201526c0e0c2d2e440dad2e6dac2e8c6d609b1b6044820152606490fd5b908160209103126106cc575180151581036106cc5790565b91939290935f83526020948584015260018060a01b03166040830152608060608301528051908160808401525f5b8281106107c857505060a09293505f838284010152601f8019910116010190565b81810186015184820160a0015285016107a7565b919082018092116107e957565b634e487b7160e01b5f52601160045260245ffd5b6103e5808202918083048214811517156107e9578402029281840414811517156107e9576103e8918281029281840414901517156107e95761083e916107dc565b908115610849570490565b634e487b7160e01b5f52601260045260245ffdfea26469706673582212209ba68aab20ca39b5fa97fc52981fda706724bdf8375d5f27b486e615ce52913364736f6c63430008140033",
}

// ArbitrageExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbitrageExecutorMetaData.ABI instead.
var ArbitrageExecutorABI = ArbitrageExecutorMetaData.ABI

// ArbitrageExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArbitrageExecutorMetaData.Bin instead.
var ArbitrageExecutorBin = ArbitrageExecutorMetaData.Bin

// DeployArbitrageExecutor deploys a new Ethereum contract, binding an instance of ArbitrageExecutor to it.
func DeployArbitrageExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbitrageExecutor, error) {
	parsed, err := ArbitrageExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrageExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrageExecutor{ArbitrageExecutorCaller: ArbitrageExecutorCaller{contract: contract}, ArbitrageExecutorTransactor: ArbitrageExecutorTransactor{contract: contract}, ArbitrageExecutorFilterer: ArbitrageExecutorFilterer{contract: contract}}, nil
}

// ArbitrageExecutor is an auto generated Go binding around an Ethereum contract.
type ArbitrageExecutor struct {
	ArbitrageExecutorCaller     // Read-only binding to the contract
	ArbitrageExecutorTransactor // Write-only binding to the contract
	ArbitrageExecutorFilterer   // Log filterer for contract events
}

// ArbitrageExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbitrageExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbitrageExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbitrageExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitrageExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbitrageExecutorSession struct {
	Contract     *ArbitrageExecutor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ArbitrageExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbitrageExecutorCallerSession struct {
	Contract *ArbitrageExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ArbitrageExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbitrageExecutorTransactorSession struct {
	Contract     *ArbitrageExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ArbitrageExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbitrageExecutorRaw struct {
	Contract *ArbitrageExecutor // Generic contract binding to access the raw methods on
}

// ArbitrageExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbitrageExecutorCallerRaw struct {
	Contract *ArbitrageExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// ArbitrageExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbitrageExecutorTransactorRaw struct {
	Contract *ArbitrageExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbitrageExecutor creates a new instance of ArbitrageExecutor, bound to a specific deployed contract.
func NewArbitrageExecutor(address common.Address, backend bind.ContractBackend) (*ArbitrageExecutor, error) {
	contract, err := bindArbitrageExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrageExecutor{ArbitrageExecutorCaller: ArbitrageExecutorCaller{contract: contract}, ArbitrageExecutorTransactor: ArbitrageExecutorTransactor{contract: contract}, ArbitrageExecutorFilterer: ArbitrageExecutorFilterer{contract: contract}}, nil
}

// NewArbitrageExecutorCaller creates a new read-only instance of ArbitrageExecutor, bound to a specific deployed contract.
func NewArbitrageExecutorCaller(address common.Address, caller bind.ContractCaller) (*ArbitrageExecutorCaller, error) {
	contract, err := bindArbitrageExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrageExecutorCaller{contract: contract}, nil
}

// NewArbitrageExecutorTransactor creates a new write-only instance of ArbitrageExecutor, bound to a specific deployed contract.
func NewArbitrageExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrageExecutorTransactor, error) {
	contract, err := bindArbitrageExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrageExecutorTransactor{contract: contract}, nil
}

// NewArbitrageExecutorFilterer creates a new log filterer instance of ArbitrageExecutor, bound to a specific deployed contract.
func NewArbitrageExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrageExecutorFilterer, error) {
	contract, err := bindArbitrageExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrageExecutorFilterer{contract: contract}, nil
}

// bindArbitrageExecutor binds a generic wrapper to an already deployed contract.
func bindArbitrageExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrageExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbitrageExecutor *ArbitrageExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrageExecutor.Contract.ArbitrageExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbitrageExecutor *ArbitrageExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrageExecutor.Contract.ArbitrageExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbitrageExecutor *ArbitrageExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrageExecutor.Contract.ArbitrageExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbitrageExecutor *ArbitrageExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrageExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbitrageExecutor *ArbitrageExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrageExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbitrageExecutor *ArbitrageExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrageExecutor.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x6142251c.
//
// Solidity: function execute(address pairA, address pairB, uint256 maxIn, uint256 step) returns()
func (_ArbitrageExecutor *ArbitrageExecutorTransactor) Execute(opts *bind.TransactOpts, pairA common.Address, pairB common.Address, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _ArbitrageExecutor.contract.Transact(opts, "execute", pairA, pairB, maxIn, step)
}

// Execute is a paid mutator transaction binding the contract method 0x6142251c.
//
// Solidity: function execute(address pairA, address pairB, uint256 maxIn, uint256 step) returns()
func (_ArbitrageExecutor *ArbitrageExecutorSession) Execute(pairA common.Address, pairB common.Address, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _ArbitrageExecutor.Contract.Execute(&_ArbitrageExecutor.TransactOpts, pairA, pairB, maxIn, step)
}

// Execute is a paid mutator transaction binding the contract method 0x6142251c.
//
// Solidity: function execute(address pairA, address pairB, uint256 maxIn, uint256 step) returns()
func (_ArbitrageExecutor *ArbitrageExecutorTransactorSession) Execute(pairA common.Address, pairB common.Address, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _ArbitrageExecutor.Contract.Execute(&_ArbitrageExecutor.TransactOpts, pairA, pairB, maxIn, step)
}

// ArbitrageExecutorTradeExecutedIterator is returned from FilterTradeExecuted and is used to iterate over the raw logs and unpacked data for TradeExecuted events raised by the ArbitrageExecutor contract.
type ArbitrageExecutorTradeExecutedIterator struct {
	Event *ArbitrageExecutorTradeExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbitrageExecutorTradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrageExecutorTradeExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbitrageExecutorTradeExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbitrageExecutorTradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbitrageExecutorTradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbitrageExecutorTradeExecuted represents a TradeExecuted event raised by the ArbitrageExecutor contract.
type ArbitrageExecutorTradeExecuted struct {
	AmountIn *big.Int
	Profit   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTradeExecuted is a free log retrieval operation binding the contract event 0x589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec7086.
//
// Solidity: event TradeExecuted(uint256 amountIn, uint256 profit)
func (_ArbitrageExecutor *ArbitrageExecutorFilterer) FilterTradeExecuted(opts *bind.FilterOpts) (*ArbitrageExecutorTradeExecutedIterator, error) {

	logs, sub, err := _ArbitrageExecutor.contract.FilterLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return &ArbitrageExecutorTradeExecutedIterator{contract: _ArbitrageExecutor.contract, event: "TradeExecuted", logs: logs, sub: sub}, nil
}

// WatchTradeExecuted is a free log subscription operation binding the contract event 0x589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec7086.
//
// Solidity: event TradeExecuted(uint256 amountIn, uint256 profit)
func (_ArbitrageExecutor *ArbitrageExecutorFilterer) WatchTradeExecuted(opts *bind.WatchOpts, sink chan<- *ArbitrageExecutorTradeExecuted) (event.Subscription, error) {

	logs, sub, err := _ArbitrageExecutor.contract.WatchLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbitrageExecutorTradeExecuted)
				if err := _ArbitrageExecutor.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradeExecuted is a log parse operation binding the contract event 0x589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec7086.
//
// Solidity: event TradeExecuted(uint256 amountIn, uint256 profit)
func (_ArbitrageExecutor *ArbitrageExecutorFilterer) ParseTradeExecuted(log types.Log) (*ArbitrageExecutorTradeExecuted, error) {
	event := new(ArbitrageExecutorTradeExecuted)
	if err := _ArbitrageExecutor.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
