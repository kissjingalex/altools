// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// TradeTypesTradeOrder is an auto generated low-level Go binding around an user-defined struct.
type TradeTypesTradeOrder struct {
	Signer common.Address
	ItemId [32]byte
	Price  *big.Int
	V      uint8
	R      [32]byte
	S      [32]byte
}

// DefoSignMetaData contains all meta data concerning the DefoSign contract.
var DefoSignMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"itemId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structTradeTypes.TradeOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"Recipient\",\"type\":\"address\"}],\"name\":\"executeOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"itemId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structTradeTypes.TradeOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"itemId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structTradeTypes.TradeOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderTypedDataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"itemId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structTradeTypes.TradeOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"verifyOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DefoSignABI is the input ABI used to generate the binding from.
// Deprecated: Use DefoSignMetaData.ABI instead.
var DefoSignABI = DefoSignMetaData.ABI

// DefoSign is an auto generated Go binding around an Ethereum contract.
type DefoSign struct {
	DefoSignCaller     // Read-only binding to the contract
	DefoSignTransactor // Write-only binding to the contract
	DefoSignFilterer   // Log filterer for contract events
}

// DefoSignCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefoSignCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoSignTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefoSignTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoSignFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefoSignFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoSignSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefoSignSession struct {
	Contract     *DefoSign         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefoSignCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefoSignCallerSession struct {
	Contract *DefoSignCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DefoSignTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefoSignTransactorSession struct {
	Contract     *DefoSignTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DefoSignRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefoSignRaw struct {
	Contract *DefoSign // Generic contract binding to access the raw methods on
}

// DefoSignCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefoSignCallerRaw struct {
	Contract *DefoSignCaller // Generic read-only contract binding to access the raw methods on
}

// DefoSignTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefoSignTransactorRaw struct {
	Contract *DefoSignTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefoSign creates a new instance of DefoSign, bound to a specific deployed contract.
func NewDefoSign(address common.Address, backend bind.ContractBackend) (*DefoSign, error) {
	contract, err := bindDefoSign(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefoSign{DefoSignCaller: DefoSignCaller{contract: contract}, DefoSignTransactor: DefoSignTransactor{contract: contract}, DefoSignFilterer: DefoSignFilterer{contract: contract}}, nil
}

// NewDefoSignCaller creates a new read-only instance of DefoSign, bound to a specific deployed contract.
func NewDefoSignCaller(address common.Address, caller bind.ContractCaller) (*DefoSignCaller, error) {
	contract, err := bindDefoSign(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefoSignCaller{contract: contract}, nil
}

// NewDefoSignTransactor creates a new write-only instance of DefoSign, bound to a specific deployed contract.
func NewDefoSignTransactor(address common.Address, transactor bind.ContractTransactor) (*DefoSignTransactor, error) {
	contract, err := bindDefoSign(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefoSignTransactor{contract: contract}, nil
}

// NewDefoSignFilterer creates a new log filterer instance of DefoSign, bound to a specific deployed contract.
func NewDefoSignFilterer(address common.Address, filterer bind.ContractFilterer) (*DefoSignFilterer, error) {
	contract, err := bindDefoSign(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefoSignFilterer{contract: contract}, nil
}

// bindDefoSign binds a generic wrapper to an already deployed contract.
func bindDefoSign(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefoSignMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefoSign *DefoSignRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefoSign.Contract.DefoSignCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefoSign *DefoSignRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoSign.Contract.DefoSignTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefoSign *DefoSignRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefoSign.Contract.DefoSignTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefoSign *DefoSignCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefoSign.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefoSign *DefoSignTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoSign.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefoSign *DefoSignTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefoSign.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DefoSign *DefoSignCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _DefoSign.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DefoSign *DefoSignSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DefoSign.Contract.Eip712Domain(&_DefoSign.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DefoSign *DefoSignCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DefoSign.Contract.Eip712Domain(&_DefoSign.CallOpts)
}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoSign *DefoSignCaller) GetDomainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DefoSign.contract.Call(opts, &out, "getDomainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoSign *DefoSignSession) GetDomainName() (string, error) {
	return _DefoSign.Contract.GetDomainName(&_DefoSign.CallOpts)
}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoSign *DefoSignCallerSession) GetDomainName() (string, error) {
	return _DefoSign.Contract.GetDomainName(&_DefoSign.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoSign *DefoSignCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DefoSign.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoSign *DefoSignSession) GetDomainSeparator() ([32]byte, error) {
	return _DefoSign.Contract.GetDomainSeparator(&_DefoSign.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoSign *DefoSignCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _DefoSign.Contract.GetDomainSeparator(&_DefoSign.CallOpts)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xf7f2a1ff.
//
// Solidity: function getOrderHash((address,bytes32,uint256,uint8,bytes32,bytes32) order) pure returns(bytes32)
func (_DefoSign *DefoSignCaller) GetOrderHash(opts *bind.CallOpts, order TradeTypesTradeOrder) ([32]byte, error) {
	var out []interface{}
	err := _DefoSign.contract.Call(opts, &out, "getOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0xf7f2a1ff.
//
// Solidity: function getOrderHash((address,bytes32,uint256,uint8,bytes32,bytes32) order) pure returns(bytes32)
func (_DefoSign *DefoSignSession) GetOrderHash(order TradeTypesTradeOrder) ([32]byte, error) {
	return _DefoSign.Contract.GetOrderHash(&_DefoSign.CallOpts, order)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xf7f2a1ff.
//
// Solidity: function getOrderHash((address,bytes32,uint256,uint8,bytes32,bytes32) order) pure returns(bytes32)
func (_DefoSign *DefoSignCallerSession) GetOrderHash(order TradeTypesTradeOrder) ([32]byte, error) {
	return _DefoSign.Contract.GetOrderHash(&_DefoSign.CallOpts, order)
}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x9b78493c.
//
// Solidity: function getOrderTypedDataHash((address,bytes32,uint256,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoSign *DefoSignCaller) GetOrderTypedDataHash(opts *bind.CallOpts, order TradeTypesTradeOrder) ([32]byte, error) {
	var out []interface{}
	err := _DefoSign.contract.Call(opts, &out, "getOrderTypedDataHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x9b78493c.
//
// Solidity: function getOrderTypedDataHash((address,bytes32,uint256,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoSign *DefoSignSession) GetOrderTypedDataHash(order TradeTypesTradeOrder) ([32]byte, error) {
	return _DefoSign.Contract.GetOrderTypedDataHash(&_DefoSign.CallOpts, order)
}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x9b78493c.
//
// Solidity: function getOrderTypedDataHash((address,bytes32,uint256,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoSign *DefoSignCallerSession) GetOrderTypedDataHash(order TradeTypesTradeOrder) ([32]byte, error) {
	return _DefoSign.Contract.GetOrderTypedDataHash(&_DefoSign.CallOpts, order)
}

// VerifyOrder is a free data retrieval call binding the contract method 0xdb9cfda4.
//
// Solidity: function verifyOrder((address,bytes32,uint256,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoSign *DefoSignCaller) VerifyOrder(opts *bind.CallOpts, order TradeTypesTradeOrder) ([32]byte, error) {
	var out []interface{}
	err := _DefoSign.contract.Call(opts, &out, "verifyOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyOrder is a free data retrieval call binding the contract method 0xdb9cfda4.
//
// Solidity: function verifyOrder((address,bytes32,uint256,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoSign *DefoSignSession) VerifyOrder(order TradeTypesTradeOrder) ([32]byte, error) {
	return _DefoSign.Contract.VerifyOrder(&_DefoSign.CallOpts, order)
}

// VerifyOrder is a free data retrieval call binding the contract method 0xdb9cfda4.
//
// Solidity: function verifyOrder((address,bytes32,uint256,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoSign *DefoSignCallerSession) VerifyOrder(order TradeTypesTradeOrder) ([32]byte, error) {
	return _DefoSign.Contract.VerifyOrder(&_DefoSign.CallOpts, order)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x49000237.
//
// Solidity: function executeOrder((address,bytes32,uint256,uint8,bytes32,bytes32) order, address Recipient) payable returns()
func (_DefoSign *DefoSignTransactor) ExecuteOrder(opts *bind.TransactOpts, order TradeTypesTradeOrder, recipient common.Address) (*types.Transaction, error) {
	return _DefoSign.contract.Transact(opts, "executeOrder", order, recipient)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x49000237.
//
// Solidity: function executeOrder((address,bytes32,uint256,uint8,bytes32,bytes32) order, address Recipient) payable returns()
func (_DefoSign *DefoSignSession) ExecuteOrder(order TradeTypesTradeOrder, recipient common.Address) (*types.Transaction, error) {
	return _DefoSign.Contract.ExecuteOrder(&_DefoSign.TransactOpts, order, recipient)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x49000237.
//
// Solidity: function executeOrder((address,bytes32,uint256,uint8,bytes32,bytes32) order, address Recipient) payable returns()
func (_DefoSign *DefoSignTransactorSession) ExecuteOrder(order TradeTypesTradeOrder, recipient common.Address) (*types.Transaction, error) {
	return _DefoSign.Contract.ExecuteOrder(&_DefoSign.TransactOpts, order, recipient)
}

// DefoSignEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the DefoSign contract.
type DefoSignEIP712DomainChangedIterator struct {
	Event *DefoSignEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *DefoSignEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoSignEIP712DomainChanged)
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
		it.Event = new(DefoSignEIP712DomainChanged)
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
func (it *DefoSignEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoSignEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoSignEIP712DomainChanged represents a EIP712DomainChanged event raised by the DefoSign contract.
type DefoSignEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DefoSign *DefoSignFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*DefoSignEIP712DomainChangedIterator, error) {

	logs, sub, err := _DefoSign.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &DefoSignEIP712DomainChangedIterator{contract: _DefoSign.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DefoSign *DefoSignFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *DefoSignEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _DefoSign.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoSignEIP712DomainChanged)
				if err := _DefoSign.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DefoSign *DefoSignFilterer) ParseEIP712DomainChanged(log types.Log) (*DefoSignEIP712DomainChanged, error) {
	event := new(DefoSignEIP712DomainChanged)
	if err := _DefoSign.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
