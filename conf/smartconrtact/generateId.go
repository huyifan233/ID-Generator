// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generateId

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

// GenerateIdABI is the input ABI used to generate the binding from.
const GenerateIdABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"setId\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"LogId\",\"type\":\"event\"}]"

// GenerateIdBin is the compiled bytecode used for deploying new contracts.
const GenerateIdBin = `608060405234801561001057600080fd5b507f300000000000000000000000000000000000000000000000000000000000000060008190555060fc806100466000396000f3fe60806040526004361060265760003560e01c806349bfcf0214602b578063af640d0f146056575b600080fd5b605460048036036020811015603f57600080fd5b8101908080359060200190929190505050607e565b005b348015606157600080fd5b50606860c1565b6040518082815260200191505060405180910390f35b806000819055507f2ec078554f30cba9aad72c89ddc37203c0e04df1058fa6c1959f5be52ca282d46000546040518082815260200191505060405180910390a150565b6000548156fea265627a7a72305820829398546dbcf67bba9333d3db024a17d501d25356086c9d9996672df12012af64736f6c63430005090032`

// DeployGenerateId deploys a new Ethereum contract, binding an instance of GenerateId to it.
func DeployGenerateId(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GenerateId, error) {
	parsed, err := abi.JSON(strings.NewReader(GenerateIdABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenerateIdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenerateId{GenerateIdCaller: GenerateIdCaller{contract: contract}, GenerateIdTransactor: GenerateIdTransactor{contract: contract}, GenerateIdFilterer: GenerateIdFilterer{contract: contract}}, nil
}

// GenerateId is an auto generated Go binding around an Ethereum contract.
type GenerateId struct {
	GenerateIdCaller     // Read-only binding to the contract
	GenerateIdTransactor // Write-only binding to the contract
	GenerateIdFilterer   // Log filterer for contract events
}

// GenerateIdCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenerateIdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenerateIdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenerateIdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenerateIdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenerateIdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenerateIdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenerateIdSession struct {
	Contract     *GenerateId       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenerateIdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenerateIdCallerSession struct {
	Contract *GenerateIdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GenerateIdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenerateIdTransactorSession struct {
	Contract     *GenerateIdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GenerateIdRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenerateIdRaw struct {
	Contract *GenerateId // Generic contract binding to access the raw methods on
}

// GenerateIdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenerateIdCallerRaw struct {
	Contract *GenerateIdCaller // Generic read-only contract binding to access the raw methods on
}

// GenerateIdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenerateIdTransactorRaw struct {
	Contract *GenerateIdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenerateId creates a new instance of GenerateId, bound to a specific deployed contract.
func NewGenerateId(address common.Address, backend bind.ContractBackend) (*GenerateId, error) {
	contract, err := bindGenerateId(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenerateId{GenerateIdCaller: GenerateIdCaller{contract: contract}, GenerateIdTransactor: GenerateIdTransactor{contract: contract}, GenerateIdFilterer: GenerateIdFilterer{contract: contract}}, nil
}

// NewGenerateIdCaller creates a new read-only instance of GenerateId, bound to a specific deployed contract.
func NewGenerateIdCaller(address common.Address, caller bind.ContractCaller) (*GenerateIdCaller, error) {
	contract, err := bindGenerateId(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenerateIdCaller{contract: contract}, nil
}

// NewGenerateIdTransactor creates a new write-only instance of GenerateId, bound to a specific deployed contract.
func NewGenerateIdTransactor(address common.Address, transactor bind.ContractTransactor) (*GenerateIdTransactor, error) {
	contract, err := bindGenerateId(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenerateIdTransactor{contract: contract}, nil
}

// NewGenerateIdFilterer creates a new log filterer instance of GenerateId, bound to a specific deployed contract.
func NewGenerateIdFilterer(address common.Address, filterer bind.ContractFilterer) (*GenerateIdFilterer, error) {
	contract, err := bindGenerateId(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenerateIdFilterer{contract: contract}, nil
}

// bindGenerateId binds a generic wrapper to an already deployed contract.
func bindGenerateId(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenerateIdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenerateId *GenerateIdRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenerateId.Contract.GenerateIdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenerateId *GenerateIdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenerateId.Contract.GenerateIdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenerateId *GenerateIdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenerateId.Contract.GenerateIdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenerateId *GenerateIdCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GenerateId.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenerateId *GenerateIdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenerateId.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenerateId *GenerateIdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenerateId.Contract.contract.Transact(opts, method, params...)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_GenerateId *GenerateIdCaller) Id(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _GenerateId.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_GenerateId *GenerateIdSession) Id() ([32]byte, error) {
	return _GenerateId.Contract.Id(&_GenerateId.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_GenerateId *GenerateIdCallerSession) Id() ([32]byte, error) {
	return _GenerateId.Contract.Id(&_GenerateId.CallOpts)
}

// SetId is a paid mutator transaction binding the contract method 0x49bfcf02.
//
// Solidity: function setId(bytes32 _id) returns()
func (_GenerateId *GenerateIdTransactor) SetId(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _GenerateId.contract.Transact(opts, "setId", _id)
}

// SetId is a paid mutator transaction binding the contract method 0x49bfcf02.
//
// Solidity: function setId(bytes32 _id) returns()
func (_GenerateId *GenerateIdSession) SetId(_id [32]byte) (*types.Transaction, error) {
	return _GenerateId.Contract.SetId(&_GenerateId.TransactOpts, _id)
}

// SetId is a paid mutator transaction binding the contract method 0x49bfcf02.
//
// Solidity: function setId(bytes32 _id) returns()
func (_GenerateId *GenerateIdTransactorSession) SetId(_id [32]byte) (*types.Transaction, error) {
	return _GenerateId.Contract.SetId(&_GenerateId.TransactOpts, _id)
}

// GenerateIdLogIdIterator is returned from FilterLogId and is used to iterate over the raw logs and unpacked data for LogId events raised by the GenerateId contract.
type GenerateIdLogIdIterator struct {
	Event *GenerateIdLogId // Event containing the contract specifics and raw log

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
func (it *GenerateIdLogIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenerateIdLogId)
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
		it.Event = new(GenerateIdLogId)
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
func (it *GenerateIdLogIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenerateIdLogIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenerateIdLogId represents a LogId event raised by the GenerateId contract.
type GenerateIdLogId struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogId is a free log retrieval operation binding the contract event 0x2ec078554f30cba9aad72c89ddc37203c0e04df1058fa6c1959f5be52ca282d4.
//
// Solidity: event LogId(bytes32 _id)
func (_GenerateId *GenerateIdFilterer) FilterLogId(opts *bind.FilterOpts) (*GenerateIdLogIdIterator, error) {

	logs, sub, err := _GenerateId.contract.FilterLogs(opts, "LogId")
	if err != nil {
		return nil, err
	}
	return &GenerateIdLogIdIterator{contract: _GenerateId.contract, event: "LogId", logs: logs, sub: sub}, nil
}

// WatchLogId is a free log subscription operation binding the contract event 0x2ec078554f30cba9aad72c89ddc37203c0e04df1058fa6c1959f5be52ca282d4.
//
// Solidity: event LogId(bytes32 _id)
func (_GenerateId *GenerateIdFilterer) WatchLogId(opts *bind.WatchOpts, sink chan<- *GenerateIdLogId) (event.Subscription, error) {

	logs, sub, err := _GenerateId.contract.WatchLogs(opts, "LogId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenerateIdLogId)
				if err := _GenerateId.contract.UnpackLog(event, "LogId", log); err != nil {
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
