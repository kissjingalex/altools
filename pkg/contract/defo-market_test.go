package contract

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"testing"
	"time"
)

// 类型定义
// WithdrawEthscription(bytes32 ethscriptionId,address Recipient,uint64 Expiration)
// "EthscriptionOrder(address signer,address creator,bytes32 ethscriptionId,uint256 quantity,address currency,uint256 price,uint256 nonce,uint64 startTime,uint64 endTime,uint16 protocolFeeDiscounted,uint16 creatorFee,bytes params)"
var typesStandard = apitypes.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"WithdrawEthscription": {
		{
			Name: "ethscriptionId",
			Type: "bytes32",
		},
		{
			Name: "recipient",
			Type: "address",
		},
		{
			Name: "expiration",
			Type: "uint64",
		},
	},
	"EthscriptionOrder": {
		{
			Name: "signer",
			Type: "address",
		},
		{
			Name: "creator",
			Type: "address",
		},
		{
			Name: "ethscriptionId",
			Type: "bytes32",
		},
		{
			Name: "quantity", // nft =1 , ft > 1
			Type: "uint256",
		},
		{
			Name: "currency",
			Type: "address",
		},
		{
			Name: "price",
			Type: "uint256",
		},
		{
			Name: "nonce",
			Type: "uint256",
		},
		{
			Name: "startTime",
			Type: "uint64",
		},
		{
			Name: "endTime",
			Type: "uint64",
		},
		{
			Name: "protocolFeeDiscounted", // x /10000
			Type: "uint16",
		},
		{
			Name: "creatorFee",
			Type: "uint16",
		},
		{
			Name: "params",
			Type: "bytes",
		},
	},
}

// 允许切换chain
var targetChainConfig = baseGoerliChainConfig // localChainConfig
var targetUserGroup = baseGoerliUserGroup     // localUserGroup

const (
	CURRENCY_ADDRESS = "0x0000000000000000000000000000000000000000"

	// local chain
	//ETHSCRIPTION_ID  = "0xdb9580c13554659ea1fb7530cdb5c3049f3af68789986433443483f7b9293880"
	//CHAIN_ID         = 31337
	//CONTRACT_ADDRESS = "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853" // deployed contract

	// base_goerli
	ETHSCRIPTION_ID  = "0x223d3ab70c75c58b2199d902e0a218ec7d0487b4b86c85a5ccb25dae71224ef4"
	CHAIN_ID         = 84531
	CONTRACT_ADDRESS = "0x608807d1489aD49373d02855bbcf10f6C1F240C3" // deployed contract
)

var domainStandard = apitypes.TypedDataDomain{
	Name:              "DefoMarket",
	Version:           "1",
	ChainId:           math.NewHexOrDecimal256(CHAIN_ID),
	VerifyingContract: CONTRACT_ADDRESS,
	Salt:              "",
}

var marketMessageStandard = map[string]interface{}{}

var typedData = apitypes.TypedData{
	Types:       typesStandard,
	PrimaryType: "EthscriptionOrder",
	Domain:      domainStandard,
	Message:     marketMessageStandard,
}

func toBytes32(bs []byte) [32]byte {
	rs := [32]byte{}
	copy(rs[:], bs[:32])
	return rs
}

func createOrder() (OrderTypesEthscriptionOrder, map[string]interface{}) {
	signer := targetUserGroup.Seller // 卖家签名这个订单
	userAddress := signer.Address
	userAddr := common.HexToAddress(userAddress)

	ethscriptionId, _ := hexutil.Decode(ETHSCRIPTION_ID)

	startTime := time.Now().Unix() - 20*60 //最好不要取当前时间
	endTime := startTime + 60*60           // 1小时以内

	order := OrderTypesEthscriptionOrder{
		Signer:                userAddr,
		Creator:               userAddr, //这里假设creator也是signer，不影响订单
		EthscriptionId:        toBytes32(ethscriptionId),
		Quantity:              big.NewInt(1),
		Currency:              common.HexToAddress(CURRENCY_ADDRESS),
		Price:                 big.NewInt(100000000000), // in wei, 100gwei
		Nonce:                 big.NewInt(1),            //order id
		StartTime:             uint64(startTime),
		EndTime:               uint64(endTime),
		ProtocolFeeDiscounted: uint16(100),
		CreatorFee:            0,
		Params:                nil,
		V:                     0,
		R:                     [32]byte{},
		S:                     [32]byte{},
	}

	orderMap := map[string]interface{}{
		"signer":                order.Signer.Hex(),
		"creator":               order.Creator.Hex(),
		"ethscriptionId":        order.EthscriptionId,
		"quantity":              order.Quantity,
		"currency":              order.Currency.Hex(),
		"price":                 order.Price,
		"nonce":                 order.Nonce,
		"startTime":             big.NewInt(int64(order.StartTime)),
		"endTime":               big.NewInt(int64(order.EndTime)),
		"protocolFeeDiscounted": big.NewInt(int64(order.ProtocolFeeDiscounted)),
		"creatorFee":            big.NewInt(int64(order.CreatorFee)),
		"params":                order.Params,
	}

	return order, orderMap
}

func TestMarket_SignLocal(t *testing.T) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("domainSeparator=%s\n", domainSeparator)

	// typedData
	_, orderMap := createOrder()

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, orderMap)
	if err != nil {
		t.Fatal(err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	signHash := crypto.Keccak256(rawData)
	t.Logf("sigHash=%s\n", hexutil.Encode(signHash)) //digest

	//signature
	signerPrivateKey := targetChainConfig.PrivateKey

	privateKey, err := crypto.HexToECDSA(signerPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	signature, err := crypto.Sign(signHash, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("signature=%s", hexutil.Encode(signature))
}

func TestMarket_ExecuteOrder(t *testing.T) {
	txContext, err := prepareContext(targetChainConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer txContext.Close()

	contractAddr := common.HexToAddress(CONTRACT_ADDRESS)
	defoMarket, err := NewDefoMarket(contractAddr, txContext.EthClient)
	if err != nil {
		t.Fatal(err)
	}

	callOpts := txContext.GetCallOpts()

	// domain separator
	domainSeparator, err := defoMarket.GetDomainSeparator(callOpts)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("domainSeparator=%s", hexutil.Encode(domainSeparator[:]))

	// sign hash
	order, orderMap := createOrder() //注意这里有时间因素，因此必须是同一个订单

	// local digest  (verified)
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, orderMap)
	if err != nil {
		t.Fatal(err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator[:]), string(typedDataHash)))
	signHash := crypto.Keccak256(rawData)
	t.Logf("local signHash=%s\n", hexutil.Encode(signHash)) //digest

	// chain digest
	//digestOnChain, err := defoMarket.GetOrderTypedDataHash(callOpts, order)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Logf("chain signHash=%s", hexutil.Encode(digestOnChain[:]))

	// local sign ==================== sign by Seller
	signer := targetUserGroup.Seller
	privateKey, err := crypto.HexToECDSA(signer.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	signature, err := crypto.Sign(signHash, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("signature=%s", hexutil.Encode(signature))

	// complete order detail with signature
	r, s, v := decodeSignature(signature)
	order.R = r
	order.S = s
	order.V = v
	spew.Dump(order)

	// execute order, for now, only one user
	// notice: it is presumed to be executed by buyer
	buyer := targetUserGroup.Buyer
	buyerPrivateKey, err := crypto.HexToECDSA(buyer.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	txOpts, err := txContext.GetTransactOpts(buyerPrivateKey, order.Price)
	if err != nil {
		t.Fatal(err)
	}

	recipient := common.HexToAddress(targetUserGroup.Buyer.Address) //买家接受
	var args2 []byte
	txExec, err := defoMarket.ExecuteEthscriptionOrder(txOpts, order, recipient, args2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("\n=====================tx debug=====================")
	spew.Dump(txExec)
	fmt.Println("=====================tx debug end=====================")

	txContext.ShowTxDetail(txExec)
}

type WithdrawEthscription struct {
	EthscriptionId [32]byte
	Recipient      common.Address
	Expiration     uint64
}

func createWithdrawEthscription() (WithdrawEthscription, map[string]interface{}) {
	seller := targetUserGroup.Seller
	sellerAddr := common.HexToAddress(seller.Address)
	ethscriptionId, _ := hexutil.Decode(ETHSCRIPTION_ID)

	expiration := time.Now().Unix() + 60*60

	we := WithdrawEthscription{
		EthscriptionId: toBytes32(ethscriptionId),
		Recipient:      sellerAddr,
		Expiration:     uint64(expiration),
	}

	weMap := map[string]interface{}{
		"ethscriptionId": we.EthscriptionId,
		"recipient":      we.Recipient.Hex(),
		"expiration":     big.NewInt(int64(we.Expiration)),
	}

	return we, weMap
}

func TestMarket_SighWithdraw(t *testing.T) {
	txContext, err := prepareContext(targetChainConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer txContext.Close()

	contractAddr := common.HexToAddress(domainStandard.VerifyingContract)
	defoMarket, err := NewDefoMarket(contractAddr, txContext.EthClient)
	if err != nil {
		t.Fatal(err)
	}

	callOpts := txContext.GetCallOpts()

	// domain separator
	domainSeparator, err := defoMarket.GetDomainSeparator(callOpts)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("domainSeparator=%s", hexutil.Encode(domainSeparator[:]))

	// local sign
	we, weMap := createWithdrawEthscription()
	dataType := "WithdrawEthscription" //数据类型

	typedDataHash, err := typedData.HashStruct(dataType, weMap)
	if err != nil {
		t.Fatal(err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator[:]), string(typedDataHash)))
	signHash := crypto.Keccak256(rawData)
	t.Logf("local signHash=%s\n", hexutil.Encode(signHash)) //digest

	// signed by contract owner
	signer := targetUserGroup.Admin
	privateKey, err := crypto.HexToECDSA(signer.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	signature, err := crypto.Sign(signHash, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	//注意调整签名
	signature[64] = signature[64] + 27

	t.Logf("signature=%s", hexutil.Encode(signature))

	// executed by seller
	seller := targetUserGroup.Seller
	sellerPrivateKey, err := crypto.HexToECDSA(seller.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	txOpts, err := txContext.GetTransactOpts(sellerPrivateKey, big.NewInt(0)) // no payment
	if err != nil {
		t.Fatal(err)
	}

	txExec, err := defoMarket.WithdrawEthscription(txOpts, we.EthscriptionId, we.Expiration, signature)
	if err != nil {
		t.Fatal(err)
	}

	txContext.ShowTxDetail(txExec)
}
