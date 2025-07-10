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

// TriangularArbitrageExecutorMetaData contains all meta data concerning the TriangularArbitrageExecutor contract.
var TriangularArbitrageExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"pairAB\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pairBC\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pairCA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"step\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TradeExecuted\",\"inputs\":[{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"profit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080806040523461001657610ad8908161001b8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f803560e01c633343a58514610026575f80fd5b346108945760a036600319011261089457600435916001600160a01b0383168303610890576024356001600160a01b0381169003610890576044356001600160a01b038116900361089057630240bc6b60e21b80825260609081836004816001600160a01b0389165afa918215610885578493859361085f575b5060405182815281816004816024356001600160a01b03165afa9081156108545786908792610833575b5060405193845282846004816044356001600160a01b03165afa93841561062057879388956107fb575b5050604051630dfe168160e01b81526020816004816001600160a01b038d165afa9081156107b25788916107dc575b5060405163d21220a760e01b81526020816004816044356001600160a01b03165afa908115610665579061016d92918a91610774575b506001600160a01b0391821691161461096a565b60405163d21220a760e01b81526020816004816001600160a01b038d165afa9081156107b25788916107bd575b50604051630dfe168160e01b81526020816004816024356001600160a01b03165afa90811561066557906101e392918a9161077457506001600160a01b0391821691161461096a565b60405163d21220a760e01b81526020816004816024356001600160a01b03165afa9081156107b2578891610793575b50604051630dfe168160e01b81526020816004816044356001600160a01b03165afa908115610665579061025b92918a9161077457506001600160a01b0391821691161461096a565b86958795608435805b6064358111156106ee57505086156106bd57604051630dfe168160e01b81526020816004816001600160a01b038f165afa908115610693578a9161069e575b506040516323b872dd60e01b81523360048201526001600160a01b03808d166024830152604482018b90529091602091839160649183918f91165af180156106935761030493926001600160701b03928392610674575b5016911688610a42565b604051610310816108bf565b8881526001600160a01b038a163b1561067057886103488b829360405194858094819363022c0d9f60e01b8352308a600485016109be565b03926001600160a01b03165af180156106655761064a575b5060405163d21220a760e01b8152979896979596899690602090829060049082906001600160a01b03165afa90811561062057879161062b575b5060405163a9059cbb60e01b81526001600160a01b03602480358216600484015282018490529091602091839160449183918c91165af18015610620576103f5946001600160701b0392839261059b575b5016921690610a42565b604051610401816108bf565b8481526024356001600160a01b03163b1561061c5784610437916040518093819263022c0d9f60e01b83523087600485016109be565b0381836024356001600160a01b03165af19081156105ca578591610604575b505060405163d21220a760e01b81526020816004816024356001600160a01b03165afa9081156105ca5785916105d5575b5060405163a9059cbb60e01b81526001600160a01b0360448035821660048401526024830185905291926020928492909183918a91165af180156105ca576104e2946001600160701b0392839261059b575016921690610a42565b604051906104ef826108bf565b8282526044356001600160a01b03163b1561058c5760405163022c0d9f60e01b81529183918391829161052891903390600485016109be565b0381836044356001600160a01b03165af1801561059057610578575b507f589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec70866040848482519182526020820152a180f35b61058190610897565b61058c57825f610544565b8280fd5b6040513d84823e3d90fd5b6105bc9060203d6020116105c3575b6105b481836108db565b8101906109a6565b505f6103eb565b503d6105aa565b6040513d87823e3d90fd5b6105f7915060203d6020116105fd575b6105ef81836108db565b81019061094b565b5f610487565b503d6105e5565b61060d90610897565b61061857835f610456565b8380fd5b8480fd5b6040513d89823e3d90fd5b610644915060203d6020116105fd576105ef81836108db565b5f61039a565b60209996979861065b600492610897565b9897969950610360565b6040513d8b823e3d90fd5b8880fd5b61068c9060203d6020116105c3576105b481836108db565b505f6102fa565b6040513d8c823e3d90fd5b6106b7915060203d6020116105fd576105ef81836108db565b5f6102a3565b60405162461bcd60e51b81526020600482015260096024820152681b9bc81c1c9bd99a5d60ba1b6044820152606490fd5b6107206001600160701b03808a169061071b818b1691808b169061071b818c1691808b16908c1689610a42565b610a42565b8181111561076d578082810311610759578190035b89811161074d575b508161074891610a21565b610264565b9099509750888161073d565b634e487b7160e01b8c52601160045260248cfd5b508a610735565b61078d915060203d6020116105fd576105ef81836108db565b5f610159565b6107ac915060203d6020116105fd576105ef81836108db565b5f610212565b6040513d8a823e3d90fd5b6107d6915060203d6020116105fd576105ef81836108db565b5f61019a565b6107f5915060203d6020116105fd576105ef81836108db565b5f610123565b8091929550610820939450903d1061082c575b61081881836108db565b810190610915565b50919091925f806100f4565b503d61080e565b905061084c9150823d841161082c5761081881836108db565b50905f6100ca565b6040513d88823e3d90fd5b9061087a9294508093503d841161082c5761081881836108db565b50929092915f6100a0565b6040513d86823e3d90fd5b5080fd5b80fd5b67ffffffffffffffff81116108ab57604052565b634e487b7160e01b5f52604160045260245ffd5b6020810190811067ffffffffffffffff8211176108ab57604052565b90601f8019910116810190811067ffffffffffffffff8211176108ab57604052565b51906001600160701b038216820361091157565b5f80fd5b9081606091031261091157610929816108fd565b916040610938602084016108fd565b92015163ffffffff811681036109115790565b9081602091031261091157516001600160a01b03811681036109115790565b1561097157565b60405162461bcd60e51b815260206004820152600d60248201526c0e0c2d2e440dad2e6dac2e8c6d609b1b6044820152606490fd5b90816020910312610911575180151581036109115790565b91939290935f83526020948584015260018060a01b03166040830152608060608301528051908160808401525f5b828110610a0d57505060a09293505f838284010152601f8019910116010190565b81810186015184820160a0015285016109ec565b91908201809211610a2e57565b634e487b7160e01b5f52601160045260245ffd5b6103e580820291808304821481151715610a2e57840202928184041481151715610a2e576103e891828102928184041490151715610a2e57610a8391610a21565b908115610a8e570490565b634e487b7160e01b5f52601260045260245ffdfea2646970667358221220b15fb8223865b3968860d754ddb9def79769a41063024eef07b4d6957b1b08ca64736f6c63430008140033",
}

// TriangularArbitrageExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use TriangularArbitrageExecutorMetaData.ABI instead.
var TriangularArbitrageExecutorABI = TriangularArbitrageExecutorMetaData.ABI

// TriangularArbitrageExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TriangularArbitrageExecutorMetaData.Bin instead.
var TriangularArbitrageExecutorBin = TriangularArbitrageExecutorMetaData.Bin

// DeployTriangularArbitrageExecutor deploys a new Ethereum contract, binding an instance of TriangularArbitrageExecutor to it.
func DeployTriangularArbitrageExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TriangularArbitrageExecutor, error) {
	parsed, err := TriangularArbitrageExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TriangularArbitrageExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TriangularArbitrageExecutor{TriangularArbitrageExecutorCaller: TriangularArbitrageExecutorCaller{contract: contract}, TriangularArbitrageExecutorTransactor: TriangularArbitrageExecutorTransactor{contract: contract}, TriangularArbitrageExecutorFilterer: TriangularArbitrageExecutorFilterer{contract: contract}}, nil
}

// TriangularArbitrageExecutor is an auto generated Go binding around an Ethereum contract.
type TriangularArbitrageExecutor struct {
	TriangularArbitrageExecutorCaller     // Read-only binding to the contract
	TriangularArbitrageExecutorTransactor // Write-only binding to the contract
	TriangularArbitrageExecutorFilterer   // Log filterer for contract events
}

// TriangularArbitrageExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TriangularArbitrageExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriangularArbitrageExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TriangularArbitrageExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriangularArbitrageExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TriangularArbitrageExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriangularArbitrageExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TriangularArbitrageExecutorSession struct {
	Contract     *TriangularArbitrageExecutor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TriangularArbitrageExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TriangularArbitrageExecutorCallerSession struct {
	Contract *TriangularArbitrageExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// TriangularArbitrageExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TriangularArbitrageExecutorTransactorSession struct {
	Contract     *TriangularArbitrageExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// TriangularArbitrageExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TriangularArbitrageExecutorRaw struct {
	Contract *TriangularArbitrageExecutor // Generic contract binding to access the raw methods on
}

// TriangularArbitrageExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TriangularArbitrageExecutorCallerRaw struct {
	Contract *TriangularArbitrageExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// TriangularArbitrageExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TriangularArbitrageExecutorTransactorRaw struct {
	Contract *TriangularArbitrageExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTriangularArbitrageExecutor creates a new instance of TriangularArbitrageExecutor, bound to a specific deployed contract.
func NewTriangularArbitrageExecutor(address common.Address, backend bind.ContractBackend) (*TriangularArbitrageExecutor, error) {
	contract, err := bindTriangularArbitrageExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TriangularArbitrageExecutor{TriangularArbitrageExecutorCaller: TriangularArbitrageExecutorCaller{contract: contract}, TriangularArbitrageExecutorTransactor: TriangularArbitrageExecutorTransactor{contract: contract}, TriangularArbitrageExecutorFilterer: TriangularArbitrageExecutorFilterer{contract: contract}}, nil
}

// NewTriangularArbitrageExecutorCaller creates a new read-only instance of TriangularArbitrageExecutor, bound to a specific deployed contract.
func NewTriangularArbitrageExecutorCaller(address common.Address, caller bind.ContractCaller) (*TriangularArbitrageExecutorCaller, error) {
	contract, err := bindTriangularArbitrageExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TriangularArbitrageExecutorCaller{contract: contract}, nil
}

// NewTriangularArbitrageExecutorTransactor creates a new write-only instance of TriangularArbitrageExecutor, bound to a specific deployed contract.
func NewTriangularArbitrageExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*TriangularArbitrageExecutorTransactor, error) {
	contract, err := bindTriangularArbitrageExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TriangularArbitrageExecutorTransactor{contract: contract}, nil
}

// NewTriangularArbitrageExecutorFilterer creates a new log filterer instance of TriangularArbitrageExecutor, bound to a specific deployed contract.
func NewTriangularArbitrageExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*TriangularArbitrageExecutorFilterer, error) {
	contract, err := bindTriangularArbitrageExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TriangularArbitrageExecutorFilterer{contract: contract}, nil
}

// bindTriangularArbitrageExecutor binds a generic wrapper to an already deployed contract.
func bindTriangularArbitrageExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TriangularArbitrageExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TriangularArbitrageExecutor.Contract.TriangularArbitrageExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriangularArbitrageExecutor.Contract.TriangularArbitrageExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TriangularArbitrageExecutor.Contract.TriangularArbitrageExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TriangularArbitrageExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriangularArbitrageExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TriangularArbitrageExecutor.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x3343a585.
//
// Solidity: function execute(address pairAB, address pairBC, address pairCA, uint256 maxIn, uint256 step) returns()
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorTransactor) Execute(opts *bind.TransactOpts, pairAB common.Address, pairBC common.Address, pairCA common.Address, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _TriangularArbitrageExecutor.contract.Transact(opts, "execute", pairAB, pairBC, pairCA, maxIn, step)
}

// Execute is a paid mutator transaction binding the contract method 0x3343a585.
//
// Solidity: function execute(address pairAB, address pairBC, address pairCA, uint256 maxIn, uint256 step) returns()
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorSession) Execute(pairAB common.Address, pairBC common.Address, pairCA common.Address, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _TriangularArbitrageExecutor.Contract.Execute(&_TriangularArbitrageExecutor.TransactOpts, pairAB, pairBC, pairCA, maxIn, step)
}

// Execute is a paid mutator transaction binding the contract method 0x3343a585.
//
// Solidity: function execute(address pairAB, address pairBC, address pairCA, uint256 maxIn, uint256 step) returns()
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorTransactorSession) Execute(pairAB common.Address, pairBC common.Address, pairCA common.Address, maxIn *big.Int, step *big.Int) (*types.Transaction, error) {
	return _TriangularArbitrageExecutor.Contract.Execute(&_TriangularArbitrageExecutor.TransactOpts, pairAB, pairBC, pairCA, maxIn, step)
}

// TriangularArbitrageExecutorTradeExecutedIterator is returned from FilterTradeExecuted and is used to iterate over the raw logs and unpacked data for TradeExecuted events raised by the TriangularArbitrageExecutor contract.
type TriangularArbitrageExecutorTradeExecutedIterator struct {
	Event *TriangularArbitrageExecutorTradeExecuted // Event containing the contract specifics and raw log

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
func (it *TriangularArbitrageExecutorTradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriangularArbitrageExecutorTradeExecuted)
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
		it.Event = new(TriangularArbitrageExecutorTradeExecuted)
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
func (it *TriangularArbitrageExecutorTradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriangularArbitrageExecutorTradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriangularArbitrageExecutorTradeExecuted represents a TradeExecuted event raised by the TriangularArbitrageExecutor contract.
type TriangularArbitrageExecutorTradeExecuted struct {
	AmountIn *big.Int
	Profit   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTradeExecuted is a free log retrieval operation binding the contract event 0x589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec7086.
//
// Solidity: event TradeExecuted(uint256 amountIn, uint256 profit)
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorFilterer) FilterTradeExecuted(opts *bind.FilterOpts) (*TriangularArbitrageExecutorTradeExecutedIterator, error) {

	logs, sub, err := _TriangularArbitrageExecutor.contract.FilterLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return &TriangularArbitrageExecutorTradeExecutedIterator{contract: _TriangularArbitrageExecutor.contract, event: "TradeExecuted", logs: logs, sub: sub}, nil
}

// WatchTradeExecuted is a free log subscription operation binding the contract event 0x589de19ac049f650e30154ebb6ba9be12c5394027648b0ff506705d197ec7086.
//
// Solidity: event TradeExecuted(uint256 amountIn, uint256 profit)
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorFilterer) WatchTradeExecuted(opts *bind.WatchOpts, sink chan<- *TriangularArbitrageExecutorTradeExecuted) (event.Subscription, error) {

	logs, sub, err := _TriangularArbitrageExecutor.contract.WatchLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriangularArbitrageExecutorTradeExecuted)
				if err := _TriangularArbitrageExecutor.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
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
func (_TriangularArbitrageExecutor *TriangularArbitrageExecutorFilterer) ParseTradeExecuted(log types.Log) (*TriangularArbitrageExecutorTradeExecuted, error) {
	event := new(TriangularArbitrageExecutorTradeExecuted)
	if err := _TriangularArbitrageExecutor.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
