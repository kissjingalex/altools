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

// OrderTypesEthscriptionOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderTypesEthscriptionOrder struct {
	Signer                common.Address
	Creator               common.Address
	EthscriptionId        [32]byte
	Quantity              *big.Int
	Currency              common.Address
	Price                 *big.Int
	Nonce                 *big.Int
	StartTime             uint64
	EndTime               uint64
	ProtocolFeeDiscounted uint16
	CreatorFee            uint16
	Params                []byte
	V                     uint8
	R                     [32]byte
	S                     [32]byte
}

// DefoMarketMetaData contains all meta data concerning the DefoMarket contract.
var DefoMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrencyInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyOrderCancelList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthscriptionInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoncesInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNonceTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedSignatureInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"CancelAllOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"CancelMultipleOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"EthscriptionDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"EthscriptionOrderExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"EthscriptionWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"creatorFeeBps\",\"type\":\"uint96\"}],\"name\":\"NewCreatorFeeBps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currencyManager\",\"type\":\"address\"}],\"name\":\"NewCurrencyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"protocolFeeBps\",\"type\":\"uint96\"}],\"name\":\"NewProtocolFeeBps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"NewProtocolFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trustedVerifier\",\"type\":\"address\"}],\"name\":\"NewTrustedVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ethscriptions_protocol_TransferEthscriptionForPreviousOwner\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"cancelAllOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"cancelMultipleMakerOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"protocolFeeDiscounted\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"creatorFee\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.EthscriptionOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeEthscriptionOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"protocolFeeDiscounted\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"creatorFee\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.EthscriptionOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderTypedDataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMinOrderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethscriptionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"expiration\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"trustedSign\",\"type\":\"bytes\"}],\"name\":\"withdrawEthscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DefoMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use DefoMarketMetaData.ABI instead.
var DefoMarketABI = DefoMarketMetaData.ABI

// DefoMarket is an auto generated Go binding around an Ethereum contract.
type DefoMarket struct {
	DefoMarketCaller     // Read-only binding to the contract
	DefoMarketTransactor // Write-only binding to the contract
	DefoMarketFilterer   // Log filterer for contract events
}

// DefoMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefoMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefoMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefoMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefoMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefoMarketSession struct {
	Contract     *DefoMarket       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefoMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefoMarketCallerSession struct {
	Contract *DefoMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DefoMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefoMarketTransactorSession struct {
	Contract     *DefoMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DefoMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefoMarketRaw struct {
	Contract *DefoMarket // Generic contract binding to access the raw methods on
}

// DefoMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefoMarketCallerRaw struct {
	Contract *DefoMarketCaller // Generic read-only contract binding to access the raw methods on
}

// DefoMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefoMarketTransactorRaw struct {
	Contract *DefoMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefoMarket creates a new instance of DefoMarket, bound to a specific deployed contract.
func NewDefoMarket(address common.Address, backend bind.ContractBackend) (*DefoMarket, error) {
	contract, err := bindDefoMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefoMarket{DefoMarketCaller: DefoMarketCaller{contract: contract}, DefoMarketTransactor: DefoMarketTransactor{contract: contract}, DefoMarketFilterer: DefoMarketFilterer{contract: contract}}, nil
}

// NewDefoMarketCaller creates a new read-only instance of DefoMarket, bound to a specific deployed contract.
func NewDefoMarketCaller(address common.Address, caller bind.ContractCaller) (*DefoMarketCaller, error) {
	contract, err := bindDefoMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefoMarketCaller{contract: contract}, nil
}

// NewDefoMarketTransactor creates a new write-only instance of DefoMarket, bound to a specific deployed contract.
func NewDefoMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*DefoMarketTransactor, error) {
	contract, err := bindDefoMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefoMarketTransactor{contract: contract}, nil
}

// NewDefoMarketFilterer creates a new log filterer instance of DefoMarket, bound to a specific deployed contract.
func NewDefoMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*DefoMarketFilterer, error) {
	contract, err := bindDefoMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefoMarketFilterer{contract: contract}, nil
}

// bindDefoMarket binds a generic wrapper to an already deployed contract.
func bindDefoMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefoMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefoMarket *DefoMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefoMarket.Contract.DefoMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefoMarket *DefoMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoMarket.Contract.DefoMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefoMarket *DefoMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefoMarket.Contract.DefoMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefoMarket *DefoMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefoMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefoMarket *DefoMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefoMarket *DefoMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefoMarket.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DefoMarket *DefoMarketCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _DefoMarket.contract.Call(opts, &out, "eip712Domain")

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
func (_DefoMarket *DefoMarketSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DefoMarket.Contract.Eip712Domain(&_DefoMarket.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DefoMarket *DefoMarketCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DefoMarket.Contract.Eip712Domain(&_DefoMarket.CallOpts)
}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoMarket *DefoMarketCaller) GetDomainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DefoMarket.contract.Call(opts, &out, "getDomainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoMarket *DefoMarketSession) GetDomainName() (string, error) {
	return _DefoMarket.Contract.GetDomainName(&_DefoMarket.CallOpts)
}

// GetDomainName is a free data retrieval call binding the contract method 0x72403331.
//
// Solidity: function getDomainName() view returns(string)
func (_DefoMarket *DefoMarketCallerSession) GetDomainName() (string, error) {
	return _DefoMarket.Contract.GetDomainName(&_DefoMarket.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoMarket *DefoMarketCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DefoMarket.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoMarket *DefoMarketSession) GetDomainSeparator() ([32]byte, error) {
	return _DefoMarket.Contract.GetDomainSeparator(&_DefoMarket.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_DefoMarket *DefoMarketCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _DefoMarket.Contract.GetDomainSeparator(&_DefoMarket.CallOpts)
}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x00c1904c.
//
// Solidity: function getOrderTypedDataHash((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint16,bytes,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoMarket *DefoMarketCaller) GetOrderTypedDataHash(opts *bind.CallOpts, order OrderTypesEthscriptionOrder) ([32]byte, error) {
	var out []interface{}
	err := _DefoMarket.contract.Call(opts, &out, "getOrderTypedDataHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x00c1904c.
//
// Solidity: function getOrderTypedDataHash((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint16,bytes,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoMarket *DefoMarketSession) GetOrderTypedDataHash(order OrderTypesEthscriptionOrder) ([32]byte, error) {
	return _DefoMarket.Contract.GetOrderTypedDataHash(&_DefoMarket.CallOpts, order)
}

// GetOrderTypedDataHash is a free data retrieval call binding the contract method 0x00c1904c.
//
// Solidity: function getOrderTypedDataHash((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint16,bytes,uint8,bytes32,bytes32) order) view returns(bytes32)
func (_DefoMarket *DefoMarketCallerSession) GetOrderTypedDataHash(order OrderTypesEthscriptionOrder) ([32]byte, error) {
	return _DefoMarket.Contract.GetOrderTypedDataHash(&_DefoMarket.CallOpts, order)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_DefoMarket *DefoMarketCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefoMarket.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_DefoMarket *DefoMarketSession) GetVerifier() (common.Address, error) {
	return _DefoMarket.Contract.GetVerifier(&_DefoMarket.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_DefoMarket *DefoMarketCallerSession) GetVerifier() (common.Address, error) {
	return _DefoMarket.Contract.GetVerifier(&_DefoMarket.CallOpts)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_DefoMarket *DefoMarketCaller) UserMinOrderNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DefoMarket.contract.Call(opts, &out, "userMinOrderNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_DefoMarket *DefoMarketSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _DefoMarket.Contract.UserMinOrderNonce(&_DefoMarket.CallOpts, arg0)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_DefoMarket *DefoMarketCallerSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _DefoMarket.Contract.UserMinOrderNonce(&_DefoMarket.CallOpts, arg0)
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0x028afabf.
//
// Solidity: function cancelAllOrders() returns()
func (_DefoMarket *DefoMarketTransactor) CancelAllOrders(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoMarket.contract.Transact(opts, "cancelAllOrders")
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0x028afabf.
//
// Solidity: function cancelAllOrders() returns()
func (_DefoMarket *DefoMarketSession) CancelAllOrders() (*types.Transaction, error) {
	return _DefoMarket.Contract.CancelAllOrders(&_DefoMarket.TransactOpts)
}

// CancelAllOrders is a paid mutator transaction binding the contract method 0x028afabf.
//
// Solidity: function cancelAllOrders() returns()
func (_DefoMarket *DefoMarketTransactorSession) CancelAllOrders() (*types.Transaction, error) {
	return _DefoMarket.Contract.CancelAllOrders(&_DefoMarket.TransactOpts)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_DefoMarket *DefoMarketTransactor) CancelMultipleMakerOrders(opts *bind.TransactOpts, orderNonces []*big.Int) (*types.Transaction, error) {
	return _DefoMarket.contract.Transact(opts, "cancelMultipleMakerOrders", orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_DefoMarket *DefoMarketSession) CancelMultipleMakerOrders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _DefoMarket.Contract.CancelMultipleMakerOrders(&_DefoMarket.TransactOpts, orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_DefoMarket *DefoMarketTransactorSession) CancelMultipleMakerOrders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _DefoMarket.Contract.CancelMultipleMakerOrders(&_DefoMarket.TransactOpts, orderNonces)
}

// ExecuteEthscriptionOrder is a paid mutator transaction binding the contract method 0xd92a1740.
//
// Solidity: function executeEthscriptionOrder((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint16,bytes,uint8,bytes32,bytes32) order, address recipient, bytes ) payable returns()
func (_DefoMarket *DefoMarketTransactor) ExecuteEthscriptionOrder(opts *bind.TransactOpts, order OrderTypesEthscriptionOrder, recipient common.Address, arg2 []byte) (*types.Transaction, error) {
	return _DefoMarket.contract.Transact(opts, "executeEthscriptionOrder", order, recipient, arg2)
}

// ExecuteEthscriptionOrder is a paid mutator transaction binding the contract method 0xd92a1740.
//
// Solidity: function executeEthscriptionOrder((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint16,bytes,uint8,bytes32,bytes32) order, address recipient, bytes ) payable returns()
func (_DefoMarket *DefoMarketSession) ExecuteEthscriptionOrder(order OrderTypesEthscriptionOrder, recipient common.Address, arg2 []byte) (*types.Transaction, error) {
	return _DefoMarket.Contract.ExecuteEthscriptionOrder(&_DefoMarket.TransactOpts, order, recipient, arg2)
}

// ExecuteEthscriptionOrder is a paid mutator transaction binding the contract method 0xd92a1740.
//
// Solidity: function executeEthscriptionOrder((address,address,bytes32,uint256,address,uint256,uint256,uint64,uint64,uint16,uint16,bytes,uint8,bytes32,bytes32) order, address recipient, bytes ) payable returns()
func (_DefoMarket *DefoMarketTransactorSession) ExecuteEthscriptionOrder(order OrderTypesEthscriptionOrder, recipient common.Address, arg2 []byte) (*types.Transaction, error) {
	return _DefoMarket.Contract.ExecuteEthscriptionOrder(&_DefoMarket.TransactOpts, order, recipient, arg2)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_DefoMarket *DefoMarketTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefoMarket.contract.Transact(opts, "withdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_DefoMarket *DefoMarketSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefoMarket.Contract.WithdrawETH(&_DefoMarket.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_DefoMarket *DefoMarketTransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DefoMarket.Contract.WithdrawETH(&_DefoMarket.TransactOpts, to, amount)
}

// WithdrawEthscription is a paid mutator transaction binding the contract method 0x924b71d6.
//
// Solidity: function withdrawEthscription(bytes32 ethscriptionId, uint64 expiration, bytes trustedSign) returns()
func (_DefoMarket *DefoMarketTransactor) WithdrawEthscription(opts *bind.TransactOpts, ethscriptionId [32]byte, expiration uint64, trustedSign []byte) (*types.Transaction, error) {
	return _DefoMarket.contract.Transact(opts, "withdrawEthscription", ethscriptionId, expiration, trustedSign)
}

// WithdrawEthscription is a paid mutator transaction binding the contract method 0x924b71d6.
//
// Solidity: function withdrawEthscription(bytes32 ethscriptionId, uint64 expiration, bytes trustedSign) returns()
func (_DefoMarket *DefoMarketSession) WithdrawEthscription(ethscriptionId [32]byte, expiration uint64, trustedSign []byte) (*types.Transaction, error) {
	return _DefoMarket.Contract.WithdrawEthscription(&_DefoMarket.TransactOpts, ethscriptionId, expiration, trustedSign)
}

// WithdrawEthscription is a paid mutator transaction binding the contract method 0x924b71d6.
//
// Solidity: function withdrawEthscription(bytes32 ethscriptionId, uint64 expiration, bytes trustedSign) returns()
func (_DefoMarket *DefoMarketTransactorSession) WithdrawEthscription(ethscriptionId [32]byte, expiration uint64, trustedSign []byte) (*types.Transaction, error) {
	return _DefoMarket.Contract.WithdrawEthscription(&_DefoMarket.TransactOpts, ethscriptionId, expiration, trustedSign)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DefoMarket *DefoMarketTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DefoMarket.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DefoMarket *DefoMarketSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DefoMarket.Contract.Fallback(&_DefoMarket.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DefoMarket *DefoMarketTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DefoMarket.Contract.Fallback(&_DefoMarket.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefoMarket *DefoMarketTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefoMarket.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefoMarket *DefoMarketSession) Receive() (*types.Transaction, error) {
	return _DefoMarket.Contract.Receive(&_DefoMarket.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefoMarket *DefoMarketTransactorSession) Receive() (*types.Transaction, error) {
	return _DefoMarket.Contract.Receive(&_DefoMarket.TransactOpts)
}

// DefoMarketCancelAllOrdersIterator is returned from FilterCancelAllOrders and is used to iterate over the raw logs and unpacked data for CancelAllOrders events raised by the DefoMarket contract.
type DefoMarketCancelAllOrdersIterator struct {
	Event *DefoMarketCancelAllOrders // Event containing the contract specifics and raw log

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
func (it *DefoMarketCancelAllOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketCancelAllOrders)
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
		it.Event = new(DefoMarketCancelAllOrders)
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
func (it *DefoMarketCancelAllOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketCancelAllOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketCancelAllOrders represents a CancelAllOrders event raised by the DefoMarket contract.
type DefoMarketCancelAllOrders struct {
	User        common.Address
	NewMinNonce *big.Int
	Timestamp   uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelAllOrders is a free log retrieval operation binding the contract event 0x692159ac25967247c250ff3ad11f2c390430d4ac38944f63da07860e76feec8b.
//
// Solidity: event CancelAllOrders(address user, uint256 newMinNonce, uint64 timestamp)
func (_DefoMarket *DefoMarketFilterer) FilterCancelAllOrders(opts *bind.FilterOpts) (*DefoMarketCancelAllOrdersIterator, error) {

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "CancelAllOrders")
	if err != nil {
		return nil, err
	}
	return &DefoMarketCancelAllOrdersIterator{contract: _DefoMarket.contract, event: "CancelAllOrders", logs: logs, sub: sub}, nil
}

// WatchCancelAllOrders is a free log subscription operation binding the contract event 0x692159ac25967247c250ff3ad11f2c390430d4ac38944f63da07860e76feec8b.
//
// Solidity: event CancelAllOrders(address user, uint256 newMinNonce, uint64 timestamp)
func (_DefoMarket *DefoMarketFilterer) WatchCancelAllOrders(opts *bind.WatchOpts, sink chan<- *DefoMarketCancelAllOrders) (event.Subscription, error) {

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "CancelAllOrders")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketCancelAllOrders)
				if err := _DefoMarket.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseCancelAllOrders(log types.Log) (*DefoMarketCancelAllOrders, error) {
	event := new(DefoMarketCancelAllOrders)
	if err := _DefoMarket.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketCancelMultipleOrdersIterator is returned from FilterCancelMultipleOrders and is used to iterate over the raw logs and unpacked data for CancelMultipleOrders events raised by the DefoMarket contract.
type DefoMarketCancelMultipleOrdersIterator struct {
	Event *DefoMarketCancelMultipleOrders // Event containing the contract specifics and raw log

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
func (it *DefoMarketCancelMultipleOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketCancelMultipleOrders)
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
		it.Event = new(DefoMarketCancelMultipleOrders)
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
func (it *DefoMarketCancelMultipleOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketCancelMultipleOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketCancelMultipleOrders represents a CancelMultipleOrders event raised by the DefoMarket contract.
type DefoMarketCancelMultipleOrders struct {
	User        common.Address
	OrderNonces []*big.Int
	Timestamp   uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelMultipleOrders is a free log retrieval operation binding the contract event 0x447a449ff88d1c270a413a1edc839b011295e3f954a7fb37f4bd15d61e2b842b.
//
// Solidity: event CancelMultipleOrders(address user, uint256[] orderNonces, uint64 timestamp)
func (_DefoMarket *DefoMarketFilterer) FilterCancelMultipleOrders(opts *bind.FilterOpts) (*DefoMarketCancelMultipleOrdersIterator, error) {

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "CancelMultipleOrders")
	if err != nil {
		return nil, err
	}
	return &DefoMarketCancelMultipleOrdersIterator{contract: _DefoMarket.contract, event: "CancelMultipleOrders", logs: logs, sub: sub}, nil
}

// WatchCancelMultipleOrders is a free log subscription operation binding the contract event 0x447a449ff88d1c270a413a1edc839b011295e3f954a7fb37f4bd15d61e2b842b.
//
// Solidity: event CancelMultipleOrders(address user, uint256[] orderNonces, uint64 timestamp)
func (_DefoMarket *DefoMarketFilterer) WatchCancelMultipleOrders(opts *bind.WatchOpts, sink chan<- *DefoMarketCancelMultipleOrders) (event.Subscription, error) {

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "CancelMultipleOrders")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketCancelMultipleOrders)
				if err := _DefoMarket.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseCancelMultipleOrders(log types.Log) (*DefoMarketCancelMultipleOrders, error) {
	event := new(DefoMarketCancelMultipleOrders)
	if err := _DefoMarket.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the DefoMarket contract.
type DefoMarketEIP712DomainChangedIterator struct {
	Event *DefoMarketEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *DefoMarketEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketEIP712DomainChanged)
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
		it.Event = new(DefoMarketEIP712DomainChanged)
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
func (it *DefoMarketEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketEIP712DomainChanged represents a EIP712DomainChanged event raised by the DefoMarket contract.
type DefoMarketEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DefoMarket *DefoMarketFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*DefoMarketEIP712DomainChangedIterator, error) {

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &DefoMarketEIP712DomainChangedIterator{contract: _DefoMarket.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DefoMarket *DefoMarketFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *DefoMarketEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketEIP712DomainChanged)
				if err := _DefoMarket.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseEIP712DomainChanged(log types.Log) (*DefoMarketEIP712DomainChanged, error) {
	event := new(DefoMarketEIP712DomainChanged)
	if err := _DefoMarket.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketEthscriptionDepositedIterator is returned from FilterEthscriptionDeposited and is used to iterate over the raw logs and unpacked data for EthscriptionDeposited events raised by the DefoMarket contract.
type DefoMarketEthscriptionDepositedIterator struct {
	Event *DefoMarketEthscriptionDeposited // Event containing the contract specifics and raw log

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
func (it *DefoMarketEthscriptionDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketEthscriptionDeposited)
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
		it.Event = new(DefoMarketEthscriptionDeposited)
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
func (it *DefoMarketEthscriptionDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketEthscriptionDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketEthscriptionDeposited represents a EthscriptionDeposited event raised by the DefoMarket contract.
type DefoMarketEthscriptionDeposited struct {
	Owner          common.Address
	EthscriptionId [32]byte
	Timestamp      uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEthscriptionDeposited is a free log retrieval operation binding the contract event 0xb7a0cccb198549b78d6c922c373fd26013f05b1f206cf19fcc3e91b52bb96815.
//
// Solidity: event EthscriptionDeposited(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarket *DefoMarketFilterer) FilterEthscriptionDeposited(opts *bind.FilterOpts, owner []common.Address, ethscriptionId [][32]byte) (*DefoMarketEthscriptionDepositedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var ethscriptionIdRule []interface{}
	for _, ethscriptionIdItem := range ethscriptionId {
		ethscriptionIdRule = append(ethscriptionIdRule, ethscriptionIdItem)
	}

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "EthscriptionDeposited", ownerRule, ethscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketEthscriptionDepositedIterator{contract: _DefoMarket.contract, event: "EthscriptionDeposited", logs: logs, sub: sub}, nil
}

// WatchEthscriptionDeposited is a free log subscription operation binding the contract event 0xb7a0cccb198549b78d6c922c373fd26013f05b1f206cf19fcc3e91b52bb96815.
//
// Solidity: event EthscriptionDeposited(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarket *DefoMarketFilterer) WatchEthscriptionDeposited(opts *bind.WatchOpts, sink chan<- *DefoMarketEthscriptionDeposited, owner []common.Address, ethscriptionId [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var ethscriptionIdRule []interface{}
	for _, ethscriptionIdItem := range ethscriptionId {
		ethscriptionIdRule = append(ethscriptionIdRule, ethscriptionIdItem)
	}

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "EthscriptionDeposited", ownerRule, ethscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketEthscriptionDeposited)
				if err := _DefoMarket.contract.UnpackLog(event, "EthscriptionDeposited", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseEthscriptionDeposited(log types.Log) (*DefoMarketEthscriptionDeposited, error) {
	event := new(DefoMarketEthscriptionDeposited)
	if err := _DefoMarket.contract.UnpackLog(event, "EthscriptionDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketEthscriptionOrderExecutedIterator is returned from FilterEthscriptionOrderExecuted and is used to iterate over the raw logs and unpacked data for EthscriptionOrderExecuted events raised by the DefoMarket contract.
type DefoMarketEthscriptionOrderExecutedIterator struct {
	Event *DefoMarketEthscriptionOrderExecuted // Event containing the contract specifics and raw log

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
func (it *DefoMarketEthscriptionOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketEthscriptionOrderExecuted)
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
		it.Event = new(DefoMarketEthscriptionOrderExecuted)
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
func (it *DefoMarketEthscriptionOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketEthscriptionOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketEthscriptionOrderExecuted represents a EthscriptionOrderExecuted event raised by the DefoMarket contract.
type DefoMarketEthscriptionOrderExecuted struct {
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
func (_DefoMarket *DefoMarketFilterer) FilterEthscriptionOrderExecuted(opts *bind.FilterOpts, orderHash [][32]byte) (*DefoMarketEthscriptionOrderExecutedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "EthscriptionOrderExecuted", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketEthscriptionOrderExecutedIterator{contract: _DefoMarket.contract, event: "EthscriptionOrderExecuted", logs: logs, sub: sub}, nil
}

// WatchEthscriptionOrderExecuted is a free log subscription operation binding the contract event 0x93a6900c7e12c8592eb245abc171ff4709f08fcba2e290729cd16cfe71380260.
//
// Solidity: event EthscriptionOrderExecuted(bytes32 indexed orderHash, uint256 orderNonce, bytes32 ethscriptionId, uint256 quantity, address seller, address buyer, address currency, uint256 price, uint64 endTime)
func (_DefoMarket *DefoMarketFilterer) WatchEthscriptionOrderExecuted(opts *bind.WatchOpts, sink chan<- *DefoMarketEthscriptionOrderExecuted, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "EthscriptionOrderExecuted", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketEthscriptionOrderExecuted)
				if err := _DefoMarket.contract.UnpackLog(event, "EthscriptionOrderExecuted", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseEthscriptionOrderExecuted(log types.Log) (*DefoMarketEthscriptionOrderExecuted, error) {
	event := new(DefoMarketEthscriptionOrderExecuted)
	if err := _DefoMarket.contract.UnpackLog(event, "EthscriptionOrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketEthscriptionWithdrawnIterator is returned from FilterEthscriptionWithdrawn and is used to iterate over the raw logs and unpacked data for EthscriptionWithdrawn events raised by the DefoMarket contract.
type DefoMarketEthscriptionWithdrawnIterator struct {
	Event *DefoMarketEthscriptionWithdrawn // Event containing the contract specifics and raw log

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
func (it *DefoMarketEthscriptionWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketEthscriptionWithdrawn)
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
		it.Event = new(DefoMarketEthscriptionWithdrawn)
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
func (it *DefoMarketEthscriptionWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketEthscriptionWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketEthscriptionWithdrawn represents a EthscriptionWithdrawn event raised by the DefoMarket contract.
type DefoMarketEthscriptionWithdrawn struct {
	Owner          common.Address
	EthscriptionId [32]byte
	Timestamp      uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEthscriptionWithdrawn is a free log retrieval operation binding the contract event 0x2de4767ccba4af9cda81fafe1bed205b39a99c1b5a9a38bfdd2d73f9c7d3a1b9.
//
// Solidity: event EthscriptionWithdrawn(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarket *DefoMarketFilterer) FilterEthscriptionWithdrawn(opts *bind.FilterOpts, owner []common.Address, ethscriptionId [][32]byte) (*DefoMarketEthscriptionWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var ethscriptionIdRule []interface{}
	for _, ethscriptionIdItem := range ethscriptionId {
		ethscriptionIdRule = append(ethscriptionIdRule, ethscriptionIdItem)
	}

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "EthscriptionWithdrawn", ownerRule, ethscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketEthscriptionWithdrawnIterator{contract: _DefoMarket.contract, event: "EthscriptionWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthscriptionWithdrawn is a free log subscription operation binding the contract event 0x2de4767ccba4af9cda81fafe1bed205b39a99c1b5a9a38bfdd2d73f9c7d3a1b9.
//
// Solidity: event EthscriptionWithdrawn(address indexed owner, bytes32 indexed ethscriptionId, uint64 timestamp)
func (_DefoMarket *DefoMarketFilterer) WatchEthscriptionWithdrawn(opts *bind.WatchOpts, sink chan<- *DefoMarketEthscriptionWithdrawn, owner []common.Address, ethscriptionId [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var ethscriptionIdRule []interface{}
	for _, ethscriptionIdItem := range ethscriptionId {
		ethscriptionIdRule = append(ethscriptionIdRule, ethscriptionIdItem)
	}

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "EthscriptionWithdrawn", ownerRule, ethscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketEthscriptionWithdrawn)
				if err := _DefoMarket.contract.UnpackLog(event, "EthscriptionWithdrawn", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseEthscriptionWithdrawn(log types.Log) (*DefoMarketEthscriptionWithdrawn, error) {
	event := new(DefoMarketEthscriptionWithdrawn)
	if err := _DefoMarket.contract.UnpackLog(event, "EthscriptionWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketNewCreatorFeeBpsIterator is returned from FilterNewCreatorFeeBps and is used to iterate over the raw logs and unpacked data for NewCreatorFeeBps events raised by the DefoMarket contract.
type DefoMarketNewCreatorFeeBpsIterator struct {
	Event *DefoMarketNewCreatorFeeBps // Event containing the contract specifics and raw log

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
func (it *DefoMarketNewCreatorFeeBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketNewCreatorFeeBps)
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
		it.Event = new(DefoMarketNewCreatorFeeBps)
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
func (it *DefoMarketNewCreatorFeeBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketNewCreatorFeeBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketNewCreatorFeeBps represents a NewCreatorFeeBps event raised by the DefoMarket contract.
type DefoMarketNewCreatorFeeBps struct {
	CreatorFeeBps *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewCreatorFeeBps is a free log retrieval operation binding the contract event 0xfe3cf77440bf81f0127725a2a60bdb8a2d3244d16fefcd74e890e42233be3a3a.
//
// Solidity: event NewCreatorFeeBps(uint96 creatorFeeBps)
func (_DefoMarket *DefoMarketFilterer) FilterNewCreatorFeeBps(opts *bind.FilterOpts) (*DefoMarketNewCreatorFeeBpsIterator, error) {

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "NewCreatorFeeBps")
	if err != nil {
		return nil, err
	}
	return &DefoMarketNewCreatorFeeBpsIterator{contract: _DefoMarket.contract, event: "NewCreatorFeeBps", logs: logs, sub: sub}, nil
}

// WatchNewCreatorFeeBps is a free log subscription operation binding the contract event 0xfe3cf77440bf81f0127725a2a60bdb8a2d3244d16fefcd74e890e42233be3a3a.
//
// Solidity: event NewCreatorFeeBps(uint96 creatorFeeBps)
func (_DefoMarket *DefoMarketFilterer) WatchNewCreatorFeeBps(opts *bind.WatchOpts, sink chan<- *DefoMarketNewCreatorFeeBps) (event.Subscription, error) {

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "NewCreatorFeeBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketNewCreatorFeeBps)
				if err := _DefoMarket.contract.UnpackLog(event, "NewCreatorFeeBps", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseNewCreatorFeeBps(log types.Log) (*DefoMarketNewCreatorFeeBps, error) {
	event := new(DefoMarketNewCreatorFeeBps)
	if err := _DefoMarket.contract.UnpackLog(event, "NewCreatorFeeBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketNewCurrencyManagerIterator is returned from FilterNewCurrencyManager and is used to iterate over the raw logs and unpacked data for NewCurrencyManager events raised by the DefoMarket contract.
type DefoMarketNewCurrencyManagerIterator struct {
	Event *DefoMarketNewCurrencyManager // Event containing the contract specifics and raw log

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
func (it *DefoMarketNewCurrencyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketNewCurrencyManager)
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
		it.Event = new(DefoMarketNewCurrencyManager)
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
func (it *DefoMarketNewCurrencyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketNewCurrencyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketNewCurrencyManager represents a NewCurrencyManager event raised by the DefoMarket contract.
type DefoMarketNewCurrencyManager struct {
	CurrencyManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCurrencyManager is a free log retrieval operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_DefoMarket *DefoMarketFilterer) FilterNewCurrencyManager(opts *bind.FilterOpts, currencyManager []common.Address) (*DefoMarketNewCurrencyManagerIterator, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketNewCurrencyManagerIterator{contract: _DefoMarket.contract, event: "NewCurrencyManager", logs: logs, sub: sub}, nil
}

// WatchNewCurrencyManager is a free log subscription operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_DefoMarket *DefoMarketFilterer) WatchNewCurrencyManager(opts *bind.WatchOpts, sink chan<- *DefoMarketNewCurrencyManager, currencyManager []common.Address) (event.Subscription, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketNewCurrencyManager)
				if err := _DefoMarket.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseNewCurrencyManager(log types.Log) (*DefoMarketNewCurrencyManager, error) {
	event := new(DefoMarketNewCurrencyManager)
	if err := _DefoMarket.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketNewProtocolFeeBpsIterator is returned from FilterNewProtocolFeeBps and is used to iterate over the raw logs and unpacked data for NewProtocolFeeBps events raised by the DefoMarket contract.
type DefoMarketNewProtocolFeeBpsIterator struct {
	Event *DefoMarketNewProtocolFeeBps // Event containing the contract specifics and raw log

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
func (it *DefoMarketNewProtocolFeeBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketNewProtocolFeeBps)
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
		it.Event = new(DefoMarketNewProtocolFeeBps)
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
func (it *DefoMarketNewProtocolFeeBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketNewProtocolFeeBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketNewProtocolFeeBps represents a NewProtocolFeeBps event raised by the DefoMarket contract.
type DefoMarketNewProtocolFeeBps struct {
	ProtocolFeeBps *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeBps is a free log retrieval operation binding the contract event 0xdc3c6f144c8a3540eb29e62ebf4b80d38985663b5c6e2255152a55a0dbb2ee44.
//
// Solidity: event NewProtocolFeeBps(uint96 protocolFeeBps)
func (_DefoMarket *DefoMarketFilterer) FilterNewProtocolFeeBps(opts *bind.FilterOpts) (*DefoMarketNewProtocolFeeBpsIterator, error) {

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "NewProtocolFeeBps")
	if err != nil {
		return nil, err
	}
	return &DefoMarketNewProtocolFeeBpsIterator{contract: _DefoMarket.contract, event: "NewProtocolFeeBps", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeBps is a free log subscription operation binding the contract event 0xdc3c6f144c8a3540eb29e62ebf4b80d38985663b5c6e2255152a55a0dbb2ee44.
//
// Solidity: event NewProtocolFeeBps(uint96 protocolFeeBps)
func (_DefoMarket *DefoMarketFilterer) WatchNewProtocolFeeBps(opts *bind.WatchOpts, sink chan<- *DefoMarketNewProtocolFeeBps) (event.Subscription, error) {

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "NewProtocolFeeBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketNewProtocolFeeBps)
				if err := _DefoMarket.contract.UnpackLog(event, "NewProtocolFeeBps", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseNewProtocolFeeBps(log types.Log) (*DefoMarketNewProtocolFeeBps, error) {
	event := new(DefoMarketNewProtocolFeeBps)
	if err := _DefoMarket.contract.UnpackLog(event, "NewProtocolFeeBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketNewProtocolFeeRecipientIterator is returned from FilterNewProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for NewProtocolFeeRecipient events raised by the DefoMarket contract.
type DefoMarketNewProtocolFeeRecipientIterator struct {
	Event *DefoMarketNewProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *DefoMarketNewProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketNewProtocolFeeRecipient)
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
		it.Event = new(DefoMarketNewProtocolFeeRecipient)
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
func (it *DefoMarketNewProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketNewProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketNewProtocolFeeRecipient represents a NewProtocolFeeRecipient event raised by the DefoMarket contract.
type DefoMarketNewProtocolFeeRecipient struct {
	ProtocolFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address protocolFeeRecipient)
func (_DefoMarket *DefoMarketFilterer) FilterNewProtocolFeeRecipient(opts *bind.FilterOpts) (*DefoMarketNewProtocolFeeRecipientIterator, error) {

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "NewProtocolFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &DefoMarketNewProtocolFeeRecipientIterator{contract: _DefoMarket.contract, event: "NewProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeRecipient is a free log subscription operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address protocolFeeRecipient)
func (_DefoMarket *DefoMarketFilterer) WatchNewProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *DefoMarketNewProtocolFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "NewProtocolFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketNewProtocolFeeRecipient)
				if err := _DefoMarket.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseNewProtocolFeeRecipient(log types.Log) (*DefoMarketNewProtocolFeeRecipient, error) {
	event := new(DefoMarketNewProtocolFeeRecipient)
	if err := _DefoMarket.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketNewTrustedVerifierIterator is returned from FilterNewTrustedVerifier and is used to iterate over the raw logs and unpacked data for NewTrustedVerifier events raised by the DefoMarket contract.
type DefoMarketNewTrustedVerifierIterator struct {
	Event *DefoMarketNewTrustedVerifier // Event containing the contract specifics and raw log

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
func (it *DefoMarketNewTrustedVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketNewTrustedVerifier)
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
		it.Event = new(DefoMarketNewTrustedVerifier)
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
func (it *DefoMarketNewTrustedVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketNewTrustedVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketNewTrustedVerifier represents a NewTrustedVerifier event raised by the DefoMarket contract.
type DefoMarketNewTrustedVerifier struct {
	TrustedVerifier common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewTrustedVerifier is a free log retrieval operation binding the contract event 0x7f6e81c6bdf8127b60bd5ff0e2b79c517d3715457a3df709d8c042fdaffe406d.
//
// Solidity: event NewTrustedVerifier(address trustedVerifier)
func (_DefoMarket *DefoMarketFilterer) FilterNewTrustedVerifier(opts *bind.FilterOpts) (*DefoMarketNewTrustedVerifierIterator, error) {

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "NewTrustedVerifier")
	if err != nil {
		return nil, err
	}
	return &DefoMarketNewTrustedVerifierIterator{contract: _DefoMarket.contract, event: "NewTrustedVerifier", logs: logs, sub: sub}, nil
}

// WatchNewTrustedVerifier is a free log subscription operation binding the contract event 0x7f6e81c6bdf8127b60bd5ff0e2b79c517d3715457a3df709d8c042fdaffe406d.
//
// Solidity: event NewTrustedVerifier(address trustedVerifier)
func (_DefoMarket *DefoMarketFilterer) WatchNewTrustedVerifier(opts *bind.WatchOpts, sink chan<- *DefoMarketNewTrustedVerifier) (event.Subscription, error) {

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "NewTrustedVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketNewTrustedVerifier)
				if err := _DefoMarket.contract.UnpackLog(event, "NewTrustedVerifier", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseNewTrustedVerifier(log types.Log) (*DefoMarketNewTrustedVerifier, error) {
	event := new(DefoMarketNewTrustedVerifier)
	if err := _DefoMarket.contract.UnpackLog(event, "NewTrustedVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator is returned from FilterEthscriptionsProtocolTransferEthscriptionForPreviousOwner and is used to iterate over the raw logs and unpacked data for EthscriptionsProtocolTransferEthscriptionForPreviousOwner events raised by the DefoMarket contract.
type DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator struct {
	Event *DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner // Event containing the contract specifics and raw log

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
func (it *DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner)
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
		it.Event = new(DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner)
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
func (it *DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner represents a EthscriptionsProtocolTransferEthscriptionForPreviousOwner event raised by the DefoMarket contract.
type DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner struct {
	PreviousOwner common.Address
	Recipient     common.Address
	Id            [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthscriptionsProtocolTransferEthscriptionForPreviousOwner is a free log retrieval operation binding the contract event 0xf1d95ed4d1680e6f665104f19c296ae52c1f64cd8114e84d55dc6349dbdafea3.
//
// Solidity: event ethscriptions_protocol_TransferEthscriptionForPreviousOwner(address indexed previousOwner, address indexed recipient, bytes32 indexed id)
func (_DefoMarket *DefoMarketFilterer) FilterEthscriptionsProtocolTransferEthscriptionForPreviousOwner(opts *bind.FilterOpts, previousOwner []common.Address, recipient []common.Address, id [][32]byte) (*DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator, error) {

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

	logs, sub, err := _DefoMarket.contract.FilterLogs(opts, "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", previousOwnerRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwnerIterator{contract: _DefoMarket.contract, event: "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", logs: logs, sub: sub}, nil
}

// WatchEthscriptionsProtocolTransferEthscriptionForPreviousOwner is a free log subscription operation binding the contract event 0xf1d95ed4d1680e6f665104f19c296ae52c1f64cd8114e84d55dc6349dbdafea3.
//
// Solidity: event ethscriptions_protocol_TransferEthscriptionForPreviousOwner(address indexed previousOwner, address indexed recipient, bytes32 indexed id)
func (_DefoMarket *DefoMarketFilterer) WatchEthscriptionsProtocolTransferEthscriptionForPreviousOwner(opts *bind.WatchOpts, sink chan<- *DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner, previousOwner []common.Address, recipient []common.Address, id [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _DefoMarket.contract.WatchLogs(opts, "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", previousOwnerRule, recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner)
				if err := _DefoMarket.contract.UnpackLog(event, "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", log); err != nil {
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
func (_DefoMarket *DefoMarketFilterer) ParseEthscriptionsProtocolTransferEthscriptionForPreviousOwner(log types.Log) (*DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner, error) {
	event := new(DefoMarketEthscriptionsProtocolTransferEthscriptionForPreviousOwner)
	if err := _DefoMarket.contract.UnpackLog(event, "ethscriptions_protocol_TransferEthscriptionForPreviousOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
