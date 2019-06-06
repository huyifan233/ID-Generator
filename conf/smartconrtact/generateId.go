// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"intId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"step\",\"type\":\"uint256\"}],\"name\":\"getId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"currentId\",\"type\":\"uint256\"}],\"name\":\"LogId\",\"type\":\"event\"}]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
const UtilsBin = `608060405234801561001057600080fd5b506040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506000908051906020019061005c929190610069565b506001808190555061010e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100aa57805160ff19168380011785556100d8565b828001600101855582156100d8579182015b828111156100d75782518255916020019190600101906100bc565b5b5090506100e591906100e9565b5090565b61010b91905b808211156101075760008160009055506001016100ef565b5090565b90565b6102408061011d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634f4550a114610046578063545a153a14610064578063af640d0f14610092575b600080fd5b61004e610115565b6040518082815260200191505060405180910390f35b6100906004803603602081101561007a57600080fd5b810190808035906020019092919050505061011b565b005b61009a61016d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100da5780820151818401526020810190506100bf565b50505050905090810190601f1680156101075780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60015481565b6000809050816001600082825401925050819055507f20b868ed97d841eef3a78720764709b7c99a97815fc550ab1fbbbafe64afe6f36001546040518082815260200191505060405180910390a15050565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102035780601f106101d857610100808354040283529160200191610203565b820191906000526020600020905b8154815290600101906020018083116101e657829003601f168201915b50505050508156fea265627a7a72305820ef1c6bc21ea273b54390ad84341d1a37ed487a8abde46a6fc3f8a48cdefb7a2364736f6c63430005090032`

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes)
func (_Utils *UtilsCaller) Id(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Utils.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes)
func (_Utils *UtilsSession) Id() ([]byte, error) {
	return _Utils.Contract.Id(&_Utils.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes)
func (_Utils *UtilsCallerSession) Id() ([]byte, error) {
	return _Utils.Contract.Id(&_Utils.CallOpts)
}

// IntId is a free data retrieval call binding the contract method 0x4f4550a1.
//
// Solidity: function intId() constant returns(uint256)
func (_Utils *UtilsCaller) IntId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Utils.contract.Call(opts, out, "intId")
	return *ret0, err
}

// IntId is a free data retrieval call binding the contract method 0x4f4550a1.
//
// Solidity: function intId() constant returns(uint256)
func (_Utils *UtilsSession) IntId() (*big.Int, error) {
	return _Utils.Contract.IntId(&_Utils.CallOpts)
}

// IntId is a free data retrieval call binding the contract method 0x4f4550a1.
//
// Solidity: function intId() constant returns(uint256)
func (_Utils *UtilsCallerSession) IntId() (*big.Int, error) {
	return _Utils.Contract.IntId(&_Utils.CallOpts)
}

// GetId is a paid mutator transaction binding the contract method 0x545a153a.
//
// Solidity: function getId(uint256 step) returns()
func (_Utils *UtilsTransactor) GetId(opts *bind.TransactOpts, step *big.Int) (*types.Transaction, error) {
	return _Utils.contract.Transact(opts, "getId", step)
}

// GetId is a paid mutator transaction binding the contract method 0x545a153a.
//
// Solidity: function getId(uint256 step) returns()
func (_Utils *UtilsSession) GetId(step *big.Int) (*types.Transaction, error) {
	return _Utils.Contract.GetId(&_Utils.TransactOpts, step)
}

// GetId is a paid mutator transaction binding the contract method 0x545a153a.
//
// Solidity: function getId(uint256 step) returns()
func (_Utils *UtilsTransactorSession) GetId(step *big.Int) (*types.Transaction, error) {
	return _Utils.Contract.GetId(&_Utils.TransactOpts, step)
}

// UtilsLogIdIterator is returned from FilterLogId and is used to iterate over the raw logs and unpacked data for LogId events raised by the Utils contract.
type UtilsLogIdIterator struct {
	Event *UtilsLogId // Event containing the contract specifics and raw log

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
func (it *UtilsLogIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UtilsLogId)
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
		it.Event = new(UtilsLogId)
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
func (it *UtilsLogIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UtilsLogIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UtilsLogId represents a LogId event raised by the Utils contract.
type UtilsLogId struct {
	CurrentId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogId is a free log retrieval operation binding the contract event 0x20b868ed97d841eef3a78720764709b7c99a97815fc550ab1fbbbafe64afe6f3.
//
// Solidity: event LogId(uint256 currentId)
func (_Utils *UtilsFilterer) FilterLogId(opts *bind.FilterOpts) (*UtilsLogIdIterator, error) {

	logs, sub, err := _Utils.contract.FilterLogs(opts, "LogId")
	if err != nil {
		return nil, err
	}
	return &UtilsLogIdIterator{contract: _Utils.contract, event: "LogId", logs: logs, sub: sub}, nil
}

// WatchLogId is a free log subscription operation binding the contract event 0x20b868ed97d841eef3a78720764709b7c99a97815fc550ab1fbbbafe64afe6f3.
//
// Solidity: event LogId(uint256 currentId)
func (_Utils *UtilsFilterer) WatchLogId(opts *bind.WatchOpts, sink chan<- *UtilsLogId) (event.Subscription, error) {

	logs, sub, err := _Utils.contract.WatchLogs(opts, "LogId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UtilsLogId)
				if err := _Utils.contract.UnpackLog(event, "LogId", log); err != nil {
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
