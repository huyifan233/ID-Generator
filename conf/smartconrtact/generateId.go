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
const GenerateIdABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"intId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"step\",\"type\":\"uint256\"}],\"name\":\"getId\",\"outputs\":[{\"name\":\"returnedId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"currentId\",\"type\":\"uint256\"}],\"name\":\"LogId\",\"type\":\"event\"}]"

// GenerateIdBin is the compiled bytecode used for deploying new contracts.
const GenerateIdBin = `608060405234801561001057600080fd5b506040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506000908051906020019061005c929190610069565b506001808190555061010e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100aa57805160ff19168380011785556100d8565b828001600101855582156100d8579182015b828111156100d75782518255916020019190600101906100bc565b5b5090506100e591906100e9565b5090565b61010b91905b808211156101075760008160009055506001016100ef565b5090565b90565b61025d8061011d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634f4550a114610046578063545a153a14610064578063af640d0f146100a6575b600080fd5b61004e610129565b6040518082815260200191505060405180910390f35b6100906004803603602081101561007a57600080fd5b810190808035906020019092919050505061012f565b6040518082815260200191505060405180910390f35b6100ae61018a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100ee5780820151818401526020810190506100d3565b50505050905090810190601f16801561011b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60015481565b60008060009050826001600082825401925050819055507f20b868ed97d841eef3a78720764709b7c99a97815fc550ab1fbbbafe64afe6f36001546040518082815260200191505060405180910390a1600154915050919050565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102205780601f106101f557610100808354040283529160200191610220565b820191906000526020600020905b81548152906001019060200180831161020357829003601f168201915b50505050508156fea265627a7a72305820ff6252d302da9a3db383355f3f81e5aff8087f8c6569bcf8eb1e298dbc493a7a64736f6c63430005090032`

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
// Solidity: function id() constant returns(bytes)
func (_GenerateId *GenerateIdCaller) Id(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _GenerateId.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes)
func (_GenerateId *GenerateIdSession) Id() ([]byte, error) {
	return _GenerateId.Contract.Id(&_GenerateId.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes)
func (_GenerateId *GenerateIdCallerSession) Id() ([]byte, error) {
	return _GenerateId.Contract.Id(&_GenerateId.CallOpts)
}

// IntId is a free data retrieval call binding the contract method 0x4f4550a1.
//
// Solidity: function intId() constant returns(uint256)
func (_GenerateId *GenerateIdCaller) IntId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GenerateId.contract.Call(opts, out, "intId")
	return *ret0, err
}

// IntId is a free data retrieval call binding the contract method 0x4f4550a1.
//
// Solidity: function intId() constant returns(uint256)
func (_GenerateId *GenerateIdSession) IntId() (*big.Int, error) {
	return _GenerateId.Contract.IntId(&_GenerateId.CallOpts)
}

// IntId is a free data retrieval call binding the contract method 0x4f4550a1.
//
// Solidity: function intId() constant returns(uint256)
func (_GenerateId *GenerateIdCallerSession) IntId() (*big.Int, error) {
	return _GenerateId.Contract.IntId(&_GenerateId.CallOpts)
}

// GetId is a paid mutator transaction binding the contract method 0x545a153a.
//
// Solidity: function getId(uint256 step) returns(uint256 returnedId)
func (_GenerateId *GenerateIdTransactor) GetId(opts *bind.TransactOpts, step *big.Int) (*types.Transaction, error) {
	return _GenerateId.contract.Transact(opts, "getId", step)
}

// GetId is a paid mutator transaction binding the contract method 0x545a153a.
//
// Solidity: function getId(uint256 step) returns(uint256 returnedId)
func (_GenerateId *GenerateIdSession) GetId(step *big.Int) (*types.Transaction, error) {
	return _GenerateId.Contract.GetId(&_GenerateId.TransactOpts, step)
}

// GetId is a paid mutator transaction binding the contract method 0x545a153a.
//
// Solidity: function getId(uint256 step) returns(uint256 returnedId)
func (_GenerateId *GenerateIdTransactorSession) GetId(step *big.Int) (*types.Transaction, error) {
	return _GenerateId.Contract.GetId(&_GenerateId.TransactOpts, step)
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
	CurrentId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogId is a free log retrieval operation binding the contract event 0x20b868ed97d841eef3a78720764709b7c99a97815fc550ab1fbbbafe64afe6f3.
//
// Solidity: event LogId(uint256 currentId)
func (_GenerateId *GenerateIdFilterer) FilterLogId(opts *bind.FilterOpts) (*GenerateIdLogIdIterator, error) {

	logs, sub, err := _GenerateId.contract.FilterLogs(opts, "LogId")
	if err != nil {
		return nil, err
	}
	return &GenerateIdLogIdIterator{contract: _GenerateId.contract, event: "LogId", logs: logs, sub: sub}, nil
}

// WatchLogId is a free log subscription operation binding the contract event 0x20b868ed97d841eef3a78720764709b7c99a97815fc550ab1fbbbafe64afe6f3.
//
// Solidity: event LogId(uint256 currentId)
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
