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

// MultiArbitrageExecutorMetaData contains all meta data concerning the MultiArbitrageExecutor contract.
var MultiArbitrageExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"pairs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"feeNumerators\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"feeDenominators\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"maxIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"step\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TradeExecuted\",\"inputs\":[{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"profit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080806040523461001657610ad9908161001b8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c634a98f7eb14610025575f80fd5b346104365760a03660031901126104365760043567ffffffffffffffff8111610436576100569036906004016108c1565b60249291923567ffffffffffffffff8111610436576100799036906004016108c1565b60449491943567ffffffffffffffff81116104365761009c9036906004016108c1565b9290946002851061088e575081841480610885575b15610853576100bf8461092c565b906100c98561092c565b851561083f5760049660206001600160a01b036100e58561096e565b16604051998a8092630dfe168160e01b82525afa9788156103d3575f9861081e575b505f9897985b878110610583575061012036868a6109e4565b9761012c3688846109e4565b995f9a5f9a608435805b6064358111156104c6575050505089156104955760208b60646101588861096e565b6040516323b872dd60e01b81523360048201526001600160a01b03918216602482015260448101939093529193849283915f91165af180156103d357610476575b505f928a5b8985106101d7577f589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec708660408d8d82519182526020820152a1005b60049060206001600160a01b036101f76101f2898f8c61095e565b61096e565b166040519384809263d21220a760e01b82525afa9182156103d3575f92610441575b50610254908a8a61022a898c6109c3565b5161024d8a6102458161023d818c6109c3565b51958c61095e565b35948b61095e565b3593610a5d565b6001600160a01b0361026a6101f2888e8b61095e565b168b5f198101116103de575f198c01870361043a5733905b60405180602081011067ffffffffffffffff60208301111761040a57602081016040525f8152813b156104365760405163022c0d9f60e01b81525f60048201819052602482018690526001600160a01b0390941660448201526080606482015281516084820181905290938492905b82811061041e57505091815f80948260a49183838284010152601f801991011681010301925af180156103d3576103f2575b505f198b018610610340575b61033a9150946109a1565b9361019e565b6001860186116103de5760205f9260446103626101f28f60018c01908d61095e565b60405163a9059cbb60e01b81526001600160a01b039182166004820152602481018690529586938492165af19182156103d35761033a926103a4575b5061032f565b6103c59060203d6020116103cc575b6103bd81836108f2565b810190610a32565b505f61039e565b503d6103b3565b6040513d5f823e3d90fd5b634e487b7160e01b5f52601160045260245ffd5b67ffffffffffffffff811161040a576040525f610323565b634e487b7160e01b5f52604160045260245ffd5b602082820181015160a48884010152869450016102f1565b5f80fd5b3090610282565b6102549192506104689060203d60201161046f575b61046081836108f2565b810190610982565b9190610219565b503d610456565b61048e9060203d6020116103cc576103bd81836108f2565b505f610199565b60405162461bcd60e51b81526020600482015260096024820152681b9bc81c1c9bd99a5d60ba1b6044820152606490fd5b9699919a929b939d949587989198965f975b8b5189101561052d576105218f918f928f938f9461051a8e6105138161050c81610505816105279d6109c3565b51966109c3565b51956109c3565b51946109c3565b5193610a5d565b986109a1565b976104d8565b90969f959d939b989197509b9198939b8181115f1461057b5780828103116103de578190038e5b811161056f575b5083610566916109d7565b92909192610136565b909e509c508d8361055b565b505f8e610554565b979897600460606001600160a01b036105a06101f2858d8a61095e565b1660405192838092630240bc6b60e21b82525afa80156103d3575f915f916107c5575b506001600160701b038092166105d984896109c3565b52166105e582856109c3565b52875f198101116103de575f19880181101561071857600460206001600160a01b036106156101f2858d8a61095e565b166040519283809263d21220a760e01b82525afa9081156103d3575f916106f9575b506001820182116103de57600460206001600160a01b0361065f6101f2600187018e8b61095e565b1660405192838092630dfe168160e01b82525afa9081156103d3575f916106da575b506001600160a01b039081169116036106a55761069d906109a1565b98979861010d565b60405162461bcd60e51b815260206004820152600d60248201526c0e0c2d2e440dad2e6dac2e8c6d609b1b6044820152606490fd5b6106f3915060203d60201161046f5761046081836108f2565b5f610681565b610712915060203d60201161046f5761046081836108f2565b5f610637565b600460206001600160a01b036107326101f2858d8a61095e565b166040519283809263d21220a760e01b82525afa9081156103d3575f916107a6575b506001600160a01b038a81169116036107705761069d906109a1565b60405162461bcd60e51b815260206004820152600e60248201526d0c6f2c6d8ca40dad2e6dac2e8c6d60931b6044820152606490fd5b6107bf915060203d60201161046f5761046081836108f2565b5f610754565b9150506060813d606011610816575b816107e1606093836108f2565b81010312610436576107f2816109af565b6040610800602084016109af565b92015163ffffffff81160361043657905f6105c3565b3d91506107d4565b61083891985060203d60201161046f5761046081836108f2565b965f610107565b634e487b7160e01b5f52603260045260245ffd5b60405162461bcd60e51b815260206004820152600a6024820152690cccaca40d8cadccee8d60b31b6044820152606490fd5b508284146100b1565b62461bcd60e51b815260206004820152600e60248201526d6e656564203e3d3220706169727360901b6044820152606490fd5b9181601f840112156104365782359167ffffffffffffffff8311610436576020808501948460051b01011161043657565b90601f8019910116810190811067ffffffffffffffff82111761040a57604052565b67ffffffffffffffff811161040a5760051b60200190565b9061093682610914565b61094360405191826108f2565b8281528092610954601f1991610914565b0190602036910137565b919081101561083f5760051b0190565b356001600160a01b03811681036104365790565b9081602091031261043657516001600160a01b03811681036104365790565b5f1981146103de5760010190565b51906001600160701b038216820361043657565b805182101561083f5760209160051b010190565b919082018092116103de57565b92916109ef82610914565b916109fd60405193846108f2565b829481845260208094019160051b810192831161043657905b828210610a235750505050565b81358152908301908301610a16565b90816020910312610436575180151581036104365790565b818102929181159184041417156103de57565b93610a79610a72610a8495610a7f9497610a4a565b9384610a4a565b94610a4a565b6109d7565b908115610a8f570490565b634e487b7160e01b5f52601260045260245ffdfea2646970667358221220aadd3213fb8af6674caecd41cb296ff55c9640b066d003b19cdb6b6c81cc248d64736f6c63430008140033",
}

// MultiArbitrageExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiArbitrageExecutorMetaData.ABI instead.
var MultiArbitrageExecutorABI = MultiArbitrageExecutorMetaData.ABI

// MultiArbitrageExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiArbitrageExecutorMetaData.Bin instead.
var MultiArbitrageExecutorBin = MultiArbitrageExecutorMetaData.Bin

// DeployMultiArbitrageExecutor deploys a new Ethereum contract, binding an instance of MultiArbitrageExecutor to it.
func DeployMultiArbitrageExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MultiArbitrageExecutor, error) {
	parsed, err := MultiArbitrageExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiArbitrageExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiArbitrageExecutor{MultiArbitrageExecutorCaller: MultiArbitrageExecutorCaller{contract: contract}, MultiArbitrageExecutorTransactor: MultiArbitrageExecutorTransactor{contract: contract}, MultiArbitrageExecutorFilterer: MultiArbitrageExecutorFilterer{contract: contract}}, nil
}

// MultiArbitrageExecutor is an auto generated Go binding around an Ethereum contract.
type MultiArbitrageExecutor struct {
	MultiArbitrageExecutorCaller     // Read-only binding to the contract
	MultiArbitrageExecutorTransactor // Write-only binding to the contract
	MultiArbitrageExecutorFilterer   // Log filterer for contract events
}

// MultiArbitrageExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiArbitrageExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiArbitrageExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiArbitrageExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiArbitrageExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiArbitrageExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiArbitrageExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiArbitrageExecutorSession struct {
	Contract     *MultiArbitrageExecutor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MultiArbitrageExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiArbitrageExecutorCallerSession struct {
	Contract *MultiArbitrageExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MultiArbitrageExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiArbitrageExecutorTransactorSession struct {
	Contract     *MultiArbitrageExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MultiArbitrageExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiArbitrageExecutorRaw struct {
	Contract *MultiArbitrageExecutor // Generic contract binding to access the raw methods on
}

// MultiArbitrageExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiArbitrageExecutorCallerRaw struct {
	Contract *MultiArbitrageExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// MultiArbitrageExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiArbitrageExecutorTransactorRaw struct {
	Contract *MultiArbitrageExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiArbitrageExecutor creates a new instance of MultiArbitrageExecutor, bound to a specific deployed contract.
func NewMultiArbitrageExecutor(address common.Address, backend bind.ContractBackend) (*MultiArbitrageExecutor, error) {
	contract, err := bindMultiArbitrageExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiArbitrageExecutor{MultiArbitrageExecutorCaller: MultiArbitrageExecutorCaller{contract: contract}, MultiArbitrageExecutorTransactor: MultiArbitrageExecutorTransactor{contract: contract}, MultiArbitrageExecutorFilterer: MultiArbitrageExecutorFilterer{contract: contract}}, nil
}

// NewMultiArbitrageExecutorCaller creates a new read-only instance of MultiArbitrageExecutor, bound to a specific deployed contract.
func NewMultiArbitrageExecutorCaller(address common.Address, caller bind.ContractCaller) (*MultiArbitrageExecutorCaller, error) {
	contract, err := bindMultiArbitrageExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiArbitrageExecutorCaller{contract: contract}, nil
}

// NewMultiArbitrageExecutorTransactor creates a new write-only instance of MultiArbitrageExecutor, bound to a specific deployed contract.
func NewMultiArbitrageExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiArbitrageExecutorTransactor, error) {
	contract, err := bindMultiArbitrageExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiArbitrageExecutorTransactor{contract: contract}, nil
}

// NewMultiArbitrageExecutorFilterer creates a new log filterer instance of MultiArbitrageExecutor, bound to a specific deployed contract.
func NewMultiArbitrageExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiArbitrageExecutorFilterer, error) {
	contract, err := bindMultiArbitrageExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiArbitrageExecutorFilterer{contract: contract}, nil
}

// bindMultiArbitrageExecutor binds a generic wrapper to an already deployed contract.
func bindMultiArbitrageExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiArbitrageExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiArbitrageExecutor *MultiArbitrageExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiArbitrageExecutor.Contract.MultiArbitrageExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiArbitrageExecutor *MultiArbitrageExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiArbitrageExecutor.Contract.MultiArbitrageExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiArbitrageExecutor *MultiArbitrageExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiArbitrageExecutor.Contract.MultiArbitrageExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiArbitrageExecutor *MultiArbitrageExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiArbitrageExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiArbitrageExecutor *MultiArbitrageExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiArbitrageExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiArbitrageExecutor *MultiArbitrageExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiArbitrageExecutor.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x4a98f7eb.
//
// Solidity: function execute(address[] pairs, uint256[] feeNumerators, uint256[] feeDenominators, uint256 maxIn, uint256 step) returns()
func (_MultiArbitrageExecutor *MultiArbitrageExecutorTransactor) Execute(opts *bind.TransactOpts, pairs []common.Address, feeNumerators []*big.Int, feeDenominators []*big.Int, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _MultiArbitrageExecutor.contract.Transact(opts, "execute", pairs, feeNumerators, feeDenominators, maxIn, step)
}

// Execute is a paid mutator transaction binding the contract method 0x4a98f7eb.
//
// Solidity: function execute(address[] pairs, uint256[] feeNumerators, uint256[] feeDenominators, uint256 maxIn, uint256 step) returns()
func (_MultiArbitrageExecutor *MultiArbitrageExecutorSession) Execute(pairs []common.Address, feeNumerators []*big.Int, feeDenominators []*big.Int, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _MultiArbitrageExecutor.Contract.Execute(&_MultiArbitrageExecutor.TransactOpts, pairs, feeNumerators, feeDenominators, maxIn, step)
}

// Execute is a paid mutator transaction binding the contract method 0x4a98f7eb.
//
// Solidity: function execute(address[] pairs, uint256[] feeNumerators, uint256[] feeDenominators, uint256 maxIn, uint256 step) returns()
func (_MultiArbitrageExecutor *MultiArbitrageExecutorTransactorSession) Execute(pairs []common.Address, feeNumerators []*big.Int, feeDenominators []*big.Int, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _MultiArbitrageExecutor.Contract.Execute(&_MultiArbitrageExecutor.TransactOpts, pairs, feeNumerators, feeDenominators, maxIn, step)
}

// MultiArbitrageExecutorTradeExecutedIterator is returned from FilterTradeExecuted and is used to iterate over the raw logs and unpacked data for TradeExecuted events raised by the MultiArbitrageExecutor contract.
type MultiArbitrageExecutorTradeExecutedIterator struct {
	Event *MultiArbitrageExecutorTradeExecuted // Event containing the contract specifics and raw log

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
func (it *MultiArbitrageExecutorTradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiArbitrageExecutorTradeExecuted)
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
		it.Event = new(MultiArbitrageExecutorTradeExecuted)
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
func (it *MultiArbitrageExecutorTradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiArbitrageExecutorTradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiArbitrageExecutorTradeExecuted represents a TradeExecuted event raised by the MultiArbitrageExecutor contract.
type MultiArbitrageExecutorTradeExecuted struct {
	AmountIn *big.Int
	Profit   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTradeExecuted is a free log retrieval operation binding the contract event 0x589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec7086.
//
// Solidity: event TradeExecuted(uint256 amountIn, uint256 profit)
func (_MultiArbitrageExecutor *MultiArbitrageExecutorFilterer) FilterTradeExecuted(opts *bind.FilterOpts) (*MultiArbitrageExecutorTradeExecutedIterator, error) {

	logs, sub, err := _MultiArbitrageExecutor.contract.FilterLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return &MultiArbitrageExecutorTradeExecutedIterator{contract: _MultiArbitrageExecutor.contract, event: "TradeExecuted", logs: logs, sub: sub}, nil
}

// WatchTradeExecuted is a free log subscription operation binding the contract event 0x589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec7086.
//
// Solidity: event TradeExecuted(uint256 amountIn, uint256 profit)
func (_MultiArbitrageExecutor *MultiArbitrageExecutorFilterer) WatchTradeExecuted(opts *bind.WatchOpts, sink chan<- *MultiArbitrageExecutorTradeExecuted) (event.Subscription, error) {

	logs, sub, err := _MultiArbitrageExecutor.contract.WatchLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiArbitrageExecutorTradeExecuted)
				if err := _MultiArbitrageExecutor.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
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
func (_MultiArbitrageExecutor *MultiArbitrageExecutorFilterer) ParseTradeExecuted(log types.Log) (*MultiArbitrageExecutorTradeExecuted, error) {
	event := new(MultiArbitrageExecutorTradeExecuted)
	if err := _MultiArbitrageExecutor.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
