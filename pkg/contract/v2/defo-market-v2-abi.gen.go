// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v2

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

// OrderTypesV2EthscriptionOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderTypesV2EthscriptionOrder struct {
	Seller                common.Address
	Buyer                 common.Address
	EthscriptionId        [32]byte
	Quantity              *big.Int
	Currency              common.Address
	Price                 *big.Int
	Nonce                 *big.Int
	StartTime             uint64
	EndTime               uint64
	ProtocolFeeDiscounted uint16
	V                     uint8
	R                     [32]byte
	S                     [32]byte
}

// DefoMarketV2MetaData contains all meta data concerning the DefoMarketV2 contract.
var DefoMarketV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrencyInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyOrderCancelList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthscriptionInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoncesInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNonceTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedSignatureInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"CancelAllOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"CancelMultipleOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"EthscriptionDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"EthscriptionOrderExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"EthscriptionWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"creatorFeeBps\",\"type\":\"uint96\"}],\"name\":\"NewCreatorFeeBps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currencyManager\",\"type\":\"address\"}],\"name\":\"NewCurrencyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"protocolFeeBps\",\"type\":\"uint96\"}],\"name\":\"NewProtocolFeeBps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"NewProtocolFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trustedVerifier\",\"type\":\"address\"}],\"name\":\"NewTrustedVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ethscriptions_protocol_TransferEthscriptionForPreviousOwner\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"cancelAllOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"cancelMultipleMakerOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"protocolFeeDiscounted\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypesV2.EthscriptionOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"executeEthscriptionOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"protocolFeeDiscounted\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypesV2.EthscriptionOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderTypedDataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMinOrderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expiration\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"trustedSign\",\"type\":\"bytes\"}],\"name\":\"withdrawEthscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DefoMarketV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use DefoMarketV2MetaData.ABI instead.
var DefoMarketV2ABI = DefoMarketV2MetaData.ABI

// DefoMarketV2 is an auto generated Go binding around an Ethereum contract.
type DefoMarketV2 struct {
	DefoMarketV2Caller     // Read-only binding to the contract
	DefoMarketV2Transactor // Write-only binding to the contract
	DefoMarketV2Filterer   // Log filterer for contract events
}

// DefoMarketV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type DefoMarketV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoMarketV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DefoMarketV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoMarketV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefoMarketV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoMarketV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefoMarketV2Session struct {
	Contract     *DefoMarketV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefoMarketV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefoMarketV2CallerSession struct {
	Contract *DefoMarketV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DefoMarketV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefoMarketV2TransactorSession struct {
	Contract     *DefoMarketV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DefoMarketV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type DefoMarketV2Raw struct {
	Contract *DefoMarketV2 // Generic contract binding to access the raw methods on
}

// DefoMarketV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefoMarketV2CallerRaw struct {
	Contract *DefoMarketV2Caller // Generic read-only contract binding to access the raw methods on
}

// DefoMarketV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefoMarketV2TransactorRaw struct {
	Contract *DefoMarketV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDefoMarketV2 creates a new instance of DefoMarketV2, bound to a specific deployed contract.
func NewDefoMarketV2(address common.Address, backend bind.ContractBackend) (*DefoMarketV2, error) {
	contract, err := bindDefoMarketV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2{DefoMarketV2Caller: DefoMarketV2Caller{contract: contract}, DefoMarketV2Transactor: DefoMarketV2Transactor{contract: contract}, DefoMarketV2Filterer: DefoMarketV2Filterer{contract: contract}}, nil
}

// NewDefoMarketV2Caller creates a new read-only instance of DefoMarketV2, bound to a specific deployed contract.
func NewDefoMarketV2Caller(address common.Address, caller bind.ContractCaller) (*DefoMarketV2Caller, error) {
	contract, err := bindDefoMarketV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2Caller{contract: contract}, nil
}

// NewDefoMarketV2Transactor creates a new write-only instance of DefoMarketV2, bound to a specific deployed contract.
func NewDefoMarketV2Transactor(address common.Address, transactor bind.ContractTransactor) (*DefoMarketV2Transactor, error) {
	contract, err := bindDefoMarketV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2Transactor{contract: contract}, nil
}

// NewDefoMarketV2Filterer creates a new log filterer instance of DefoMarketV2, bound to a specific deployed contract.
func NewDefoMarketV2Filterer(address common.Address, filterer bind.ContractFilterer) (*DefoMarketV2Filterer, error) {
	contract, err := bindDefoMarketV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2Filterer{contract: contract}, nil
}

// bindDefoMarketV2 binds a generic wrapper to an already deployed contract.
func bindDefoMarketV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefoMarketV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefoMarketV2 *DefoMarketV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefoMarketV2.Contract.DefoMarketV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefoMarketV2 *DefoMarketV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.DefoMarketV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefoMarketV2 *DefoMarketV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.DefoMarketV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefoMarketV2 *DefoMarketV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefoMarketV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefoMarketV2 *DefoMarketV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefoMarketV2 *DefoMarketV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DefoMarketV2 *DefoMarketV2Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _DefoMarketV2.contract.Call(opts, &out, "eip712Domain")

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
func (_DefoMarketV2 *DefoMarketV2Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DefoMarketV2.Contract.Eip712Domain(&_DefoMarketV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DefoMarketV2 *DefoMarketV2CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DefoMarketV2.Contract.Eip712Domain(&_DefoMarketV2.CallOpts)
}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoMarketV2 *DefoMarketV2Caller) GetDomainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DefoMarketV2.contract.Call(opts, &out, "getDomainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoMarketV2 *DefoMarketV2Session) GetDomainName() (string, error) {
	return _DefoMarketV2.Contract.GetDomainName(&_DefoMarketV2.CallOpts)
}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoMarketV2 *DefoMarketV2CallerSession) GetDomainName() (string, error) {
	return _DefoMarketV2.Contract.GetDomainName(&_DefoMarketV2.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoMarketV2 *DefoMarketV2Caller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DefoMarketV2.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoMarketV2 *DefoMarketV2Session) GetDomainSeparator() ([32]byte, error) {
	return _DefoMarketV2.Contract.GetDomainSeparator(&_DefoMarketV2.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoMarketV2 *DefoMarketV2CallerSession) GetDomainSeparator() ([32]byte, error) {
	return _DefoMarketV2.Contract.GetDomainSeparator(&_DefoMarketV2.CallOpts)
}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x1b5a9122.
//
// Solidity: function getOrderTypedDataHash((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoMarketV2 *DefoMarketV2Caller) GetOrderTypedDataHash(opts *bind.CallOpts, order OrderTypesV2EthscriptionOrder) ([32]byte, error) {
	var out []interface{}
	err := _DefoMarketV2.contract.Call(opts, &out, "getOrderTypedDataHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x1b5a9122.
//
// Solidity: function getOrderTypedDataHash((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoMarketV2 *DefoMarketV2Session) GetOrderTypedDataHash(order OrderTypesV2EthscriptionOrder) ([32]byte, error) {
	return _DefoMarketV2.Contract.GetOrderTypedDataHash(&_DefoMarketV2.CallOpts, order)
}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x1b5a9122.
//
// Solidity: function getOrderTypedDataHash((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoMarketV2 *DefoMarketV2CallerSession) GetOrderTypedDataHash(order OrderTypesV2EthscriptionOrder) ([32]byte, error) {
	return _DefoMarketV2.Contract.GetOrderTypedDataHash(&_DefoMarketV2.CallOpts, order)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_DefoMarketV2 *DefoMarketV2Caller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefoMarketV2.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_DefoMarketV2 *DefoMarketV2Session) GetVerifier() (common.Address, error) {
	return _DefoMarketV2.Contract.GetVerifier(&_DefoMarketV2.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_DefoMarketV2 *DefoMarketV2CallerSession) GetVerifier() (common.Address, error) {
	return _DefoMarketV2.Contract.GetVerifier(&_DefoMarketV2.CallOpts)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_DefoMarketV2 *DefoMarketV2Caller) UserMinOrderNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DefoMarketV2.contract.Call(opts, &out, "userMinOrderNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_DefoMarketV2 *DefoMarketV2Session) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _DefoMarketV2.Contract.UserMinOrderNonce(&_DefoMarketV2.CallOpts, arg0)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_DefoMarketV2 *DefoMarketV2CallerSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _DefoMarketV2.Contract.UserMinOrderNonce(&_DefoMarketV2.CallOpts, arg0)
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0x4cee19d2.
//
// Solidity: function cancelAllOrders(address buyer) returns()
func (_DefoMarketV2 *DefoMarketV2Transactor) CancelAllOrders(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _DefoMarketV2.contract.Transact(opts, "cancelAllOrders", buyer)
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0x4cee19d2.
//
// Solidity: function cancelAllOrders(address buyer) returns()
func (_DefoMarketV2 *DefoMarketV2Session) CancelAllOrders(buyer common.Address) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.CancelAllOrders(&_DefoMarketV2.TransactOpts, buyer)
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0x4cee19d2.
//
// Solidity: function cancelAllOrders(address buyer) returns()
func (_DefoMarketV2 *DefoMarketV2TransactorSession) CancelAllOrders(buyer common.Address) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.CancelAllOrders(&_DefoMarketV2.TransactOpts, buyer)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0xa1eb4742.
//
// Solidity: function cancelMultipleMakerOrders(address buyer, uint256[] orderNonces) returns()
func (_DefoMarketV2 *DefoMarketV2Transactor) CancelMultipleMakerOrders(opts *bind.TransactOpts, buyer common.Address, orderNonces []*big.Int) (*types.Transaction, error) {
	return _DefoMarketV2.contract.Transact(opts, "cancelMultipleMakerOrders", buyer, orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0xa1eb4742.
//
// Solidity: function cancelMultipleMakerOrders(address buyer, uint256[] orderNonces) returns()
func (_DefoMarketV2 *DefoMarketV2Session) CancelMultipleMakerOrders(buyer common.Address, orderNonces []*big.Int) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.CancelMultipleMakerOrders(&_DefoMarketV2.TransactOpts, buyer, orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0xa1eb4742.
//
// Solidity: function cancelMultipleMakerOrders(address buyer, uint256[] orderNonces) returns()
func (_DefoMarketV2 *DefoMarketV2TransactorSession) CancelMultipleMakerOrders(buyer common.Address, orderNonces []*big.Int) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.CancelMultipleMakerOrders(&_DefoMarketV2.TransactOpts, buyer, orderNonces)
}

// ExecuteEthscriptionOrder is a paid mutator transaction binding the contract method 0x3b34a2a4.
//
// Solidity: function executeEthscriptionOrder((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint8,bytes32,bytes32) order) payable returns()
func (_DefoMarketV2 *DefoMarketV2Transactor) ExecuteEthscriptionOrder(opts *bind.TransactOpts, order OrderTypesV2EthscriptionOrder) (*types.Transaction, error) {
	return _DefoMarketV2.contract.Transact(opts, "executeEthscriptionOrder", order)
}

// ExecuteEthscriptionOrder is a paid mutator transaction binding the contract method 0x3b34a2a4.
//
// Solidity: function executeEthscriptionOrder((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint8,bytes32,bytes32) order) payable returns()
func (_DefoMarketV2 *DefoMarketV2Session) ExecuteEthscriptionOrder(order OrderTypesV2EthscriptionOrder) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.ExecuteEthscriptionOrder(&_DefoMarketV2.TransactOpts, order)
}

// ExecuteEthscriptionOrder is a paid mutator transaction binding the contract method 0x3b34a2a4.
//
// Solidity: function executeEthscriptionOrder((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint8,bytes32,bytes32) order) payable returns()
func (_DefoMarketV2 *DefoMarketV2TransactorSession) ExecuteEthscriptionOrder(order OrderTypesV2EthscriptionOrder) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.ExecuteEthscriptionOrder(&_DefoMarketV2.TransactOpts, order)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_DefoMarketV2 *DefoMarketV2Transactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefoMarketV2.contract.Transact(opts, "withdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_DefoMarketV2 *DefoMarketV2Session) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.WithdrawETH(&_DefoMarketV2.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_DefoMarketV2 *DefoMarketV2TransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.WithdrawETH(&_DefoMarketV2.TransactOpts, to, amount)
}

// WithdrawEthscription is a paid mutator transaction binding the contract method 0x3910d620.
//
// Solidity: function withdrawEthscription(bytes32 ethscriptionId, address recipient, uint64 expiration, bytes trustedSign) returns()
func (_DefoMarketV2 *DefoMarketV2Transactor) WithdrawEthscription(opts *bind.TransactOpts, ethscriptionId [32]byte, recipient common.Address, expiration uint64, trustedSign []byte) (*types.Transaction, error) {
	return _DefoMarketV2.contract.Transact(opts, "withdrawEthscription", ethscriptionId, recipient, expiration, trustedSign)
}

// WithdrawEthscription is a paid mutator transaction binding the contract method 0x3910d620.
//
// Solidity: function withdrawEthscription(bytes32 ethscriptionId, address recipient, uint64 expiration, bytes trustedSign) returns()
func (_DefoMarketV2 *DefoMarketV2Session) WithdrawEthscription(ethscriptionId [32]byte, recipient common.Address, expiration uint64, trustedSign []byte) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.WithdrawEthscription(&_DefoMarketV2.TransactOpts, ethscriptionId, recipient, expiration, trustedSign)
}

// WithdrawEthscription is a paid mutator transaction binding the contract method 0x3910d620.
//
// Solidity: function withdrawEthscription(bytes32 ethscriptionId, address recipient, uint64 expiration, bytes trustedSign) returns()
func (_DefoMarketV2 *DefoMarketV2TransactorSession) WithdrawEthscription(ethscriptionId [32]byte, recipient common.Address, expiration uint64, trustedSign []byte) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.WithdrawEthscription(&_DefoMarketV2.TransactOpts, ethscriptionId, recipient, expiration, trustedSign)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DefoMarketV2 *DefoMarketV2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DefoMarketV2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DefoMarketV2 *DefoMarketV2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.Fallback(&_DefoMarketV2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DefoMarketV2 *DefoMarketV2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DefoMarketV2.Contract.Fallback(&_DefoMarketV2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefoMarketV2 *DefoMarketV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoMarketV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefoMarketV2 *DefoMarketV2Session) Receive() (*types.Transaction, error) {
	return _DefoMarketV2.Contract.Receive(&_DefoMarketV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefoMarketV2 *DefoMarketV2TransactorSession) Receive() (*types.Transaction, error) {
	return _DefoMarketV2.Contract.Receive(&_DefoMarketV2.TransactOpts)
}

// DefoMarketV2CancelAllOrdersIterator is returned from FilterCancelAllOrders and is used to iterate over the raw logs and unpacked data for CancelAllOrders events raised by the DefoMarketV2 contract.
type DefoMarketV2CancelAllOrdersIterator struct {
	Event *DefoMarketV2CancelAllOrders // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2CancelAllOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2CancelAllOrders)
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
		it.Event = new(DefoMarketV2CancelAllOrders)
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
func (it *DefoMarketV2CancelAllOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2CancelAllOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2CancelAllOrders represents a CancelAllOrders event raised by the DefoMarketV2 contract.
type DefoMarketV2CancelAllOrders struct {
	User        common.Address
	NewMinNonce *big.Int
	Timestamp   uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelAllOrders is a free log retrieval operation binding the contract event 0x692159ac25967247c250ff3ad11f2c390430d4ac38944f63da07860e76feec8b.
//
// Solidity: event CancelAllOrders(address user, uint256 newMinNonce, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterCancelAllOrders(opts *bind.FilterOpts) (*DefoMarketV2CancelAllOrdersIterator, error) {

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "CancelAllOrders")
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2CancelAllOrdersIterator{contract: _DefoMarketV2.contract, event: "CancelAllOrders", logs: logs, sub: sub}, nil
}

// WatchCancelAllOrders is a free log subscription operation binding the contract event 0x692159ac25967247c250ff3ad11f2c390430d4ac38944f63da07860e76feec8b.
//
// Solidity: event CancelAllOrders(address user, uint256 newMinNonce, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchCancelAllOrders(opts *bind.WatchOpts, sink chan<- *DefoMarketV2CancelAllOrders) (event.Subscription, error) {

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "CancelAllOrders")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2CancelAllOrders)
				if err := _DefoMarketV2.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
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

// ParseCancelAllOrders is a log parse operation binding the contract event 0x692159ac25967247c250ff3ad11f2c390430d4ac38944f63da07860e76feec8b.
//
// Solidity: event CancelAllOrders(address user, uint256 newMinNonce, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseCancelAllOrders(log types.Log) (*DefoMarketV2CancelAllOrders, error) {
	event := new(DefoMarketV2CancelAllOrders)
	if err := _DefoMarketV2.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2CancelMultipleOrdersIterator is returned from FilterCancelMultipleOrders and is used to iterate over the raw logs and unpacked data for CancelMultipleOrders events raised by the DefoMarketV2 contract.
type DefoMarketV2CancelMultipleOrdersIterator struct {
	Event *DefoMarketV2CancelMultipleOrders // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2CancelMultipleOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2CancelMultipleOrders)
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
		it.Event = new(DefoMarketV2CancelMultipleOrders)
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
func (it *DefoMarketV2CancelMultipleOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2CancelMultipleOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2CancelMultipleOrders represents a CancelMultipleOrders event raised by the DefoMarketV2 contract.
type DefoMarketV2CancelMultipleOrders struct {
	User        common.Address
	OrderNonces []*big.Int
	Timestamp   uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelMultipleOrders is a free log retrieval operation binding the contract event 0x447a449ff88d1c270a413a1edc839b011295e3f954a7fb37f4bd15d61e2b842b.
//
// Solidity: event CancelMultipleOrders(address user, uint256[] orderNonces, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterCancelMultipleOrders(opts *bind.FilterOpts) (*DefoMarketV2CancelMultipleOrdersIterator, error) {

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "CancelMultipleOrders")
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2CancelMultipleOrdersIterator{contract: _DefoMarketV2.contract, event: "CancelMultipleOrders", logs: logs, sub: sub}, nil
}

// WatchCancelMultipleOrders is a free log subscription operation binding the contract event 0x447a449ff88d1c270a413a1edc839b011295e3f954a7fb37f4bd15d61e2b842b.
//
// Solidity: event CancelMultipleOrders(address user, uint256[] orderNonces, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchCancelMultipleOrders(opts *bind.WatchOpts, sink chan<- *DefoMarketV2CancelMultipleOrders) (event.Subscription, error) {

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "CancelMultipleOrders")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2CancelMultipleOrders)
				if err := _DefoMarketV2.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
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

// ParseCancelMultipleOrders is a log parse operation binding the contract event 0x447a449ff88d1c270a413a1edc839b011295e3f954a7fb37f4bd15d61e2b842b.
//
// Solidity: event CancelMultipleOrders(address user, uint256[] orderNonces, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseCancelMultipleOrders(log types.Log) (*DefoMarketV2CancelMultipleOrders, error) {
	event := new(DefoMarketV2CancelMultipleOrders)
	if err := _DefoMarketV2.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the DefoMarketV2 contract.
type DefoMarketV2EIP712DomainChangedIterator struct {
	Event *DefoMarketV2EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2EIP712DomainChanged)
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
		it.Event = new(DefoMarketV2EIP712DomainChanged)
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
func (it *DefoMarketV2EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2EIP712DomainChanged represents a EIP712DomainChanged event raised by the DefoMarketV2 contract.
type DefoMarketV2EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*DefoMarketV2EIP712DomainChangedIterator, error) {

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2EIP712DomainChangedIterator{contract: _DefoMarketV2.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *DefoMarketV2EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2EIP712DomainChanged)
				if err := _DefoMarketV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseEIP712DomainChanged(log types.Log) (*DefoMarketV2EIP712DomainChanged, error) {
	event := new(DefoMarketV2EIP712DomainChanged)
	if err := _DefoMarketV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2EthscriptionDepositedIterator is returned from FilterEthscriptionDeposited and is used to iterate over the raw logs and unpacked data for EthscriptionDeposited events raised by the DefoMarketV2 contract.
type DefoMarketV2EthscriptionDepositedIterator struct {
	Event *DefoMarketV2EthscriptionDeposited // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2EthscriptionDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2EthscriptionDeposited)
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
		it.Event = new(DefoMarketV2EthscriptionDeposited)
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
func (it *DefoMarketV2EthscriptionDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2EthscriptionDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2EthscriptionDeposited represents a EthscriptionDeposited event raised by the DefoMarketV2 contract.
type DefoMarketV2EthscriptionDeposited struct {
	Owner          common.Address
	EthscriptionId [32]byte
	Timestamp      uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEthscriptionDeposited is a free log retrieval operation binding the contract event 0xb7a0cccb198549b78d6c922c373fd26013f05b1f206cf19fcc3e91b52bb96815.
//
// Solidity: event EthscriptionDeposited(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterEthscriptionDeposited(opts *bind.FilterOpts, owner []common.Address, ethscriptionId [][32]byte) (*DefoMarketV2EthscriptionDepositedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var ethscriptionIdRule []interface{}
	for _, ethscriptionIdItem := range ethscriptionId {
		ethscriptionIdRule = append(ethscriptionIdRule, ethscriptionIdItem)
	}

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "EthscriptionDeposited", ownerRule, ethscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2EthscriptionDepositedIterator{contract: _DefoMarketV2.contract, event: "EthscriptionDeposited", logs: logs, sub: sub}, nil
}

// WatchEthscriptionDeposited is a free log subscription operation binding the contract event 0xb7a0cccb198549b78d6c922c373fd26013f05b1f206cf19fcc3e91b52bb96815.
//
// Solidity: event EthscriptionDeposited(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchEthscriptionDeposited(opts *bind.WatchOpts, sink chan<- *DefoMarketV2EthscriptionDeposited, owner []common.Address, ethscriptionId [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var ethscriptionIdRule []interface{}
	for _, ethscriptionIdItem := range ethscriptionId {
		ethscriptionIdRule = append(ethscriptionIdRule, ethscriptionIdItem)
	}

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "EthscriptionDeposited", ownerRule, ethscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2EthscriptionDeposited)
				if err := _DefoMarketV2.contract.UnpackLog(event, "EthscriptionDeposited", log); err != nil {
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

// ParseEthscriptionDeposited is a log parse operation binding the contract event 0xb7a0cccb198549b78d6c922c373fd26013f05b1f206cf19fcc3e91b52bb96815.
//
// Solidity: event EthscriptionDeposited(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseEthscriptionDeposited(log types.Log) (*DefoMarketV2EthscriptionDeposited, error) {
	event := new(DefoMarketV2EthscriptionDeposited)
	if err := _DefoMarketV2.contract.UnpackLog(event, "EthscriptionDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2EthscriptionOrderExecutedIterator is returned from FilterEthscriptionOrderExecuted and is used to iterate over the raw logs and unpacked data for EthscriptionOrderExecuted events raised by the DefoMarketV2 contract.
type DefoMarketV2EthscriptionOrderExecutedIterator struct {
	Event *DefoMarketV2EthscriptionOrderExecuted // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2EthscriptionOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2EthscriptionOrderExecuted)
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
		it.Event = new(DefoMarketV2EthscriptionOrderExecuted)
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
func (it *DefoMarketV2EthscriptionOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2EthscriptionOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2EthscriptionOrderExecuted represents a EthscriptionOrderExecuted event raised by the DefoMarketV2 contract.
type DefoMarketV2EthscriptionOrderExecuted struct {
	OrderHash      [32]byte
	OrderNonce     *big.Int
	EthscriptionId [32]byte
	Quantity       *big.Int
	Seller         common.Address
	Buyer          common.Address
	Currency       common.Address
	Price          *big.Int
	EndTime        uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEthscriptionOrderExecuted is a free log retrieval operation binding the contract event 0x93a6900c7e12c8592eb245abc171ff4709f08fcba2e290729cd16cfe71380260.
//
// Solidity: event EthscriptionOrderExecuted(bytes32 indexed orderHash, uint256 orderNonce, bytes32 ethscriptionId, uint256 quantity, address seller, address buyer, address currency, uint256 price, uint64 endTime)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterEthscriptionOrderExecuted(opts *bind.FilterOpts, orderHash [][32]byte) (*DefoMarketV2EthscriptionOrderExecutedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "EthscriptionOrderExecuted", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2EthscriptionOrderExecutedIterator{contract: _DefoMarketV2.contract, event: "EthscriptionOrderExecuted", logs: logs, sub: sub}, nil
}

// WatchEthscriptionOrderExecuted is a free log subscription operation binding the contract event 0x93a6900c7e12c8592eb245abc171ff4709f08fcba2e290729cd16cfe71380260.
//
// Solidity: event EthscriptionOrderExecuted(bytes32 indexed orderHash, uint256 orderNonce, bytes32 ethscriptionId, uint256 quantity, address seller, address buyer, address currency, uint256 price, uint64 endTime)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchEthscriptionOrderExecuted(opts *bind.WatchOpts, sink chan<- *DefoMarketV2EthscriptionOrderExecuted, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "EthscriptionOrderExecuted", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2EthscriptionOrderExecuted)
				if err := _DefoMarketV2.contract.UnpackLog(event, "EthscriptionOrderExecuted", log); err != nil {
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

// ParseEthscriptionOrderExecuted is a log parse operation binding the contract event 0x93a6900c7e12c8592eb245abc171ff4709f08fcba2e290729cd16cfe71380260.
//
// Solidity: event EthscriptionOrderExecuted(bytes32 indexed orderHash, uint256 orderNonce, bytes32 ethscriptionId, uint256 quantity, address seller, address buyer, address currency, uint256 price, uint64 endTime)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseEthscriptionOrderExecuted(log types.Log) (*DefoMarketV2EthscriptionOrderExecuted, error) {
	event := new(DefoMarketV2EthscriptionOrderExecuted)
	if err := _DefoMarketV2.contract.UnpackLog(event, "EthscriptionOrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2EthscriptionWithdrawnIterator is returned from FilterEthscriptionWithdrawn and is used to iterate over the raw logs and unpacked data for EthscriptionWithdrawn events raised by the DefoMarketV2 contract.
type DefoMarketV2EthscriptionWithdrawnIterator struct {
	Event *DefoMarketV2EthscriptionWithdrawn // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2EthscriptionWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2EthscriptionWithdrawn)
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
		it.Event = new(DefoMarketV2EthscriptionWithdrawn)
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
func (it *DefoMarketV2EthscriptionWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2EthscriptionWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2EthscriptionWithdrawn represents a EthscriptionWithdrawn event raised by the DefoMarketV2 contract.
type DefoMarketV2EthscriptionWithdrawn struct {
	Owner          common.Address
	EthscriptionId [32]byte
	Timestamp      uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEthscriptionWithdrawn is a free log retrieval operation binding the contract event 0x2de4767ccba4af9cda81fafe1bed205b39a99c1b5a9a38bfdd2d73f9c7d3a1b9.
//
// Solidity: event EthscriptionWithdrawn(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterEthscriptionWithdrawn(opts *bind.FilterOpts, owner []common.Address, ethscriptionId [][32]byte) (*DefoMarketV2EthscriptionWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var ethscriptionIdRule []interface{}
	for _, ethscriptionIdItem := range ethscriptionId {
		ethscriptionIdRule = append(ethscriptionIdRule, ethscriptionIdItem)
	}

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "EthscriptionWithdrawn", ownerRule, ethscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2EthscriptionWithdrawnIterator{contract: _DefoMarketV2.contract, event: "EthscriptionWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthscriptionWithdrawn is a free log subscription operation binding the contract event 0x2de4767ccba4af9cda81fafe1bed205b39a99c1b5a9a38bfdd2d73f9c7d3a1b9.
//
// Solidity: event EthscriptionWithdrawn(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchEthscriptionWithdrawn(opts *bind.WatchOpts, sink chan<- *DefoMarketV2EthscriptionWithdrawn, owner []common.Address, ethscriptionId [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var ethscriptionIdRule []interface{}
	for _, ethscriptionIdItem := range ethscriptionId {
		ethscriptionIdRule = append(ethscriptionIdRule, ethscriptionIdItem)
	}

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "EthscriptionWithdrawn", ownerRule, ethscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2EthscriptionWithdrawn)
				if err := _DefoMarketV2.contract.UnpackLog(event, "EthscriptionWithdrawn", log); err != nil {
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

// ParseEthscriptionWithdrawn is a log parse operation binding the contract event 0x2de4767ccba4af9cda81fafe1bed205b39a99c1b5a9a38bfdd2d73f9c7d3a1b9.
//
// Solidity: event EthscriptionWithdrawn(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseEthscriptionWithdrawn(log types.Log) (*DefoMarketV2EthscriptionWithdrawn, error) {
	event := new(DefoMarketV2EthscriptionWithdrawn)
	if err := _DefoMarketV2.contract.UnpackLog(event, "EthscriptionWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2NewCreatorFeeBpsIterator is returned from FilterNewCreatorFeeBps and is used to iterate over the raw logs and unpacked data for NewCreatorFeeBps events raised by the DefoMarketV2 contract.
type DefoMarketV2NewCreatorFeeBpsIterator struct {
	Event *DefoMarketV2NewCreatorFeeBps // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2NewCreatorFeeBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2NewCreatorFeeBps)
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
		it.Event = new(DefoMarketV2NewCreatorFeeBps)
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
func (it *DefoMarketV2NewCreatorFeeBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2NewCreatorFeeBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2NewCreatorFeeBps represents a NewCreatorFeeBps event raised by the DefoMarketV2 contract.
type DefoMarketV2NewCreatorFeeBps struct {
	CreatorFeeBps *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewCreatorFeeBps is a free log retrieval operation binding the contract event 0xfe3cf77440bf81f0127725a2a60bdb8a2d3244d16fefcd74e890e42233be3a3a.
//
// Solidity: event NewCreatorFeeBps(uint96 creatorFeeBps)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterNewCreatorFeeBps(opts *bind.FilterOpts) (*DefoMarketV2NewCreatorFeeBpsIterator, error) {

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "NewCreatorFeeBps")
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2NewCreatorFeeBpsIterator{contract: _DefoMarketV2.contract, event: "NewCreatorFeeBps", logs: logs, sub: sub}, nil
}

// WatchNewCreatorFeeBps is a free log subscription operation binding the contract event 0xfe3cf77440bf81f0127725a2a60bdb8a2d3244d16fefcd74e890e42233be3a3a.
//
// Solidity: event NewCreatorFeeBps(uint96 creatorFeeBps)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchNewCreatorFeeBps(opts *bind.WatchOpts, sink chan<- *DefoMarketV2NewCreatorFeeBps) (event.Subscription, error) {

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "NewCreatorFeeBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2NewCreatorFeeBps)
				if err := _DefoMarketV2.contract.UnpackLog(event, "NewCreatorFeeBps", log); err != nil {
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

// ParseNewCreatorFeeBps is a log parse operation binding the contract event 0xfe3cf77440bf81f0127725a2a60bdb8a2d3244d16fefcd74e890e42233be3a3a.
//
// Solidity: event NewCreatorFeeBps(uint96 creatorFeeBps)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseNewCreatorFeeBps(log types.Log) (*DefoMarketV2NewCreatorFeeBps, error) {
	event := new(DefoMarketV2NewCreatorFeeBps)
	if err := _DefoMarketV2.contract.UnpackLog(event, "NewCreatorFeeBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2NewCurrencyManagerIterator is returned from FilterNewCurrencyManager and is used to iterate over the raw logs and unpacked data for NewCurrencyManager events raised by the DefoMarketV2 contract.
type DefoMarketV2NewCurrencyManagerIterator struct {
	Event *DefoMarketV2NewCurrencyManager // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2NewCurrencyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2NewCurrencyManager)
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
		it.Event = new(DefoMarketV2NewCurrencyManager)
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
func (it *DefoMarketV2NewCurrencyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2NewCurrencyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2NewCurrencyManager represents a NewCurrencyManager event raised by the DefoMarketV2 contract.
type DefoMarketV2NewCurrencyManager struct {
	CurrencyManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCurrencyManager is a free log retrieval operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterNewCurrencyManager(opts *bind.FilterOpts, currencyManager []common.Address) (*DefoMarketV2NewCurrencyManagerIterator, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2NewCurrencyManagerIterator{contract: _DefoMarketV2.contract, event: "NewCurrencyManager", logs: logs, sub: sub}, nil
}

// WatchNewCurrencyManager is a free log subscription operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchNewCurrencyManager(opts *bind.WatchOpts, sink chan<- *DefoMarketV2NewCurrencyManager, currencyManager []common.Address) (event.Subscription, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2NewCurrencyManager)
				if err := _DefoMarketV2.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
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

// ParseNewCurrencyManager is a log parse operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseNewCurrencyManager(log types.Log) (*DefoMarketV2NewCurrencyManager, error) {
	event := new(DefoMarketV2NewCurrencyManager)
	if err := _DefoMarketV2.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2NewProtocolFeeBpsIterator is returned from FilterNewProtocolFeeBps and is used to iterate over the raw logs and unpacked data for NewProtocolFeeBps events raised by the DefoMarketV2 contract.
type DefoMarketV2NewProtocolFeeBpsIterator struct {
	Event *DefoMarketV2NewProtocolFeeBps // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2NewProtocolFeeBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2NewProtocolFeeBps)
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
		it.Event = new(DefoMarketV2NewProtocolFeeBps)
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
func (it *DefoMarketV2NewProtocolFeeBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2NewProtocolFeeBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2NewProtocolFeeBps represents a NewProtocolFeeBps event raised by the DefoMarketV2 contract.
type DefoMarketV2NewProtocolFeeBps struct {
	ProtocolFeeBps *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeBps is a free log retrieval operation binding the contract event 0xdc3c6f144c8a3540eb29e62ebf4b80d38985663b5c6e2255152a55a0dbb2ee44.
//
// Solidity: event NewProtocolFeeBps(uint96 protocolFeeBps)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterNewProtocolFeeBps(opts *bind.FilterOpts) (*DefoMarketV2NewProtocolFeeBpsIterator, error) {

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "NewProtocolFeeBps")
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2NewProtocolFeeBpsIterator{contract: _DefoMarketV2.contract, event: "NewProtocolFeeBps", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeBps is a free log subscription operation binding the contract event 0xdc3c6f144c8a3540eb29e62ebf4b80d38985663b5c6e2255152a55a0dbb2ee44.
//
// Solidity: event NewProtocolFeeBps(uint96 protocolFeeBps)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchNewProtocolFeeBps(opts *bind.WatchOpts, sink chan<- *DefoMarketV2NewProtocolFeeBps) (event.Subscription, error) {

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "NewProtocolFeeBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2NewProtocolFeeBps)
				if err := _DefoMarketV2.contract.UnpackLog(event, "NewProtocolFeeBps", log); err != nil {
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

// ParseNewProtocolFeeBps is a log parse operation binding the contract event 0xdc3c6f144c8a3540eb29e62ebf4b80d38985663b5c6e2255152a55a0dbb2ee44.
//
// Solidity: event NewProtocolFeeBps(uint96 protocolFeeBps)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseNewProtocolFeeBps(log types.Log) (*DefoMarketV2NewProtocolFeeBps, error) {
	event := new(DefoMarketV2NewProtocolFeeBps)
	if err := _DefoMarketV2.contract.UnpackLog(event, "NewProtocolFeeBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2NewProtocolFeeRecipientIterator is returned from FilterNewProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for NewProtocolFeeRecipient events raised by the DefoMarketV2 contract.
type DefoMarketV2NewProtocolFeeRecipientIterator struct {
	Event *DefoMarketV2NewProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2NewProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2NewProtocolFeeRecipient)
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
		it.Event = new(DefoMarketV2NewProtocolFeeRecipient)
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
func (it *DefoMarketV2NewProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2NewProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2NewProtocolFeeRecipient represents a NewProtocolFeeRecipient event raised by the DefoMarketV2 contract.
type DefoMarketV2NewProtocolFeeRecipient struct {
	ProtocolFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address protocolFeeRecipient)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterNewProtocolFeeRecipient(opts *bind.FilterOpts) (*DefoMarketV2NewProtocolFeeRecipientIterator, error) {

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "NewProtocolFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2NewProtocolFeeRecipientIterator{contract: _DefoMarketV2.contract, event: "NewProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeRecipient is a free log subscription operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address protocolFeeRecipient)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchNewProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *DefoMarketV2NewProtocolFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "NewProtocolFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2NewProtocolFeeRecipient)
				if err := _DefoMarketV2.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
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

// ParseNewProtocolFeeRecipient is a log parse operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address protocolFeeRecipient)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseNewProtocolFeeRecipient(log types.Log) (*DefoMarketV2NewProtocolFeeRecipient, error) {
	event := new(DefoMarketV2NewProtocolFeeRecipient)
	if err := _DefoMarketV2.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2NewTrustedVerifierIterator is returned from FilterNewTrustedVerifier and is used to iterate over the raw logs and unpacked data for NewTrustedVerifier events raised by the DefoMarketV2 contract.
type DefoMarketV2NewTrustedVerifierIterator struct {
	Event *DefoMarketV2NewTrustedVerifier // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2NewTrustedVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2NewTrustedVerifier)
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
		it.Event = new(DefoMarketV2NewTrustedVerifier)
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
func (it *DefoMarketV2NewTrustedVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2NewTrustedVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2NewTrustedVerifier represents a NewTrustedVerifier event raised by the DefoMarketV2 contract.
type DefoMarketV2NewTrustedVerifier struct {
	TrustedVerifier common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewTrustedVerifier is a free log retrieval operation binding the contract event 0x7f6e81c6bdf8127b60bd5ff0e2b79c517d3715457a3df709d8c042fdaffe406d.
//
// Solidity: event NewTrustedVerifier(address trustedVerifier)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterNewTrustedVerifier(opts *bind.FilterOpts) (*DefoMarketV2NewTrustedVerifierIterator, error) {

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "NewTrustedVerifier")
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2NewTrustedVerifierIterator{contract: _DefoMarketV2.contract, event: "NewTrustedVerifier", logs: logs, sub: sub}, nil
}

// WatchNewTrustedVerifier is a free log subscription operation binding the contract event 0x7f6e81c6bdf8127b60bd5ff0e2b79c517d3715457a3df709d8c042fdaffe406d.
//
// Solidity: event NewTrustedVerifier(address trustedVerifier)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchNewTrustedVerifier(opts *bind.WatchOpts, sink chan<- *DefoMarketV2NewTrustedVerifier) (event.Subscription, error) {

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "NewTrustedVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2NewTrustedVerifier)
				if err := _DefoMarketV2.contract.UnpackLog(event, "NewTrustedVerifier", log); err != nil {
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

// ParseNewTrustedVerifier is a log parse operation binding the contract event 0x7f6e81c6bdf8127b60bd5ff0e2b79c517d3715457a3df709d8c042fdaffe406d.
//
// Solidity: event NewTrustedVerifier(address trustedVerifier)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseNewTrustedVerifier(log types.Log) (*DefoMarketV2NewTrustedVerifier, error) {
	event := new(DefoMarketV2NewTrustedVerifier)
	if err := _DefoMarketV2.contract.UnpackLog(event, "NewTrustedVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator is returned from FilterEthscriptionsProtocolTransferEthscriptionForPreviousOwner and is used to iterate over the raw logs and unpacked data for EthscriptionsProtocolTransferEthscriptionForPreviousOwner events raised by the DefoMarketV2 contract.
type DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator struct {
	Event *DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner // Event containing the contract specifics and raw log

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
func (it *DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner)
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
		it.Event = new(DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner)
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
func (it *DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner represents a EthscriptionsProtocolTransferEthscriptionForPreviousOwner event raised by the DefoMarketV2 contract.
type DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner struct {
	PreviousOwner common.Address
	Recipient     common.Address
	Id            [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthscriptionsProtocolTransferEthscriptionForPreviousOwner is a free log retrieval operation binding the contract event 0xf1d95ed4d1680e6f665104f19c296ae52c1f64cd8114e84d55dc6349dbdafea3.
//
// Solidity: event ethscriptions_protocol_TransferEthscriptionForPreviousOwner(address indexed previousOwner, address indexed recipient, bytes32 indexed id)
func (_DefoMarketV2 *DefoMarketV2Filterer) FilterEthscriptionsProtocolTransferEthscriptionForPreviousOwner(opts *bind.FilterOpts, previousOwner []common.Address, recipient []common.Address, id [][32]byte) (*DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DefoMarketV2.contract.FilterLogs(opts, "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", previousOwnerRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator{contract: _DefoMarketV2.contract, event: "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", logs: logs, sub: sub}, nil
}

// WatchEthscriptionsProtocolTransferEthscriptionForPreviousOwner is a free log subscription operation binding the contract event 0xf1d95ed4d1680e6f665104f19c296ae52c1f64cd8114e84d55dc6349dbdafea3.
//
// Solidity: event ethscriptions_protocol_TransferEthscriptionForPreviousOwner(address indexed previousOwner, address indexed recipient, bytes32 indexed id)
func (_DefoMarketV2 *DefoMarketV2Filterer) WatchEthscriptionsProtocolTransferEthscriptionForPreviousOwner(opts *bind.WatchOpts, sink chan<- *DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner, previousOwner []common.Address, recipient []common.Address, id [][32]byte) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DefoMarketV2.contract.WatchLogs(opts, "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", previousOwnerRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner)
				if err := _DefoMarketV2.contract.UnpackLog(event, "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", log); err != nil {
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

// ParseEthscriptionsProtocolTransferEthscriptionForPreviousOwner is a log parse operation binding the contract event 0xf1d95ed4d1680e6f665104f19c296ae52c1f64cd8114e84d55dc6349dbdafea3.
//
// Solidity: event ethscriptions_protocol_TransferEthscriptionForPreviousOwner(address indexed previousOwner, address indexed recipient, bytes32 indexed id)
func (_DefoMarketV2 *DefoMarketV2Filterer) ParseEthscriptionsProtocolTransferEthscriptionForPreviousOwner(log types.Log) (*DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner, error) {
	event := new(DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner)
	if err := _DefoMarketV2.contract.UnpackLog(event, "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
