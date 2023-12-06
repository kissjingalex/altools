package v2

import (
	"altools/pkg/contract"
	cjson "altools/plugins/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"log"
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
			Name: "seller",
			Type: "address",
		},
		{
			Name: "buyer",
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
	},
}

// 允许切换chain
var targetChainConfig = contract.BaseGoerliChainConfig // BaseGoerliChainConfig // LocalChainConfig
var targetUserGroup = contract.BaseGoerliUserGroup     // BaseGoerliUserGroup     // LocalUserGroup

const (
	CURRENCY_ADDRESS = "0x0000000000000000000000000000000000000000"

	// local chain
	//ETHSCRIPTION_ID  = "0xdb9580c13554659ea1fb7530cdb5c3049f3af68789986433443483f7b9293880"
	//CHAIN_ID         = 31337
	//CONTRACT_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3" // deployed contract

	// base_goerli
	ETHSCRIPTION_ID  = "0x7aaed8765db0419096e2d81b22bc06a83b9205aed3448d2fb346b01a1fe7c9d7" // "0x223d3ab70c75c58b2199d902e0a218ec7d0487b4b86c85a5ccb25dae71224ef4"
	CHAIN_ID         = 84531
	CONTRACT_ADDRESS = "0x59F4A5239a2DF011D4c274eE7Ba5C1490d010E4A" // deployed contract
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

func createOrder() (OrderTypesV2EthscriptionOrder, map[string]interface{}) {
	seller := targetUserGroup.Seller // 卖家签名这个订单
	buyer := targetUserGroup.Buyer

	sellerAddr := common.HexToAddress(seller.Address)
	buyerAddr := common.HexToAddress(buyer.Address)

	ethscriptionId, _ := hexutil.Decode(ETHSCRIPTION_ID)

	startTime := time.Now().Unix() - 20*60 //最好不要取当前时间
	endTime := startTime + 60*60           // 1小时以内

	order := OrderTypesV2EthscriptionOrder{
		Seller:                sellerAddr,
		Buyer:                 buyerAddr,
		EthscriptionId:        toBytes32(ethscriptionId),
		Quantity:              big.NewInt(1),
		Currency:              common.HexToAddress(CURRENCY_ADDRESS),
		Price:                 big.NewInt(100000000000), // in wei, 100gwei
		Nonce:                 big.NewInt(2),            // if order nonce is same, it is unable to repeat execution of the order.
		StartTime:             uint64(startTime),
		EndTime:               uint64(endTime),
		ProtocolFeeDiscounted: uint16(100),
		V:                     0,
		R:                     [32]byte{},
		S:                     [32]byte{},
	}

	orderMap := map[string]interface{}{
		"seller":                order.Seller.Hex(),
		"buyer":                 order.Buyer.Hex(),
		"ethscriptionId":        order.EthscriptionId,
		"quantity":              order.Quantity,
		"currency":              order.Currency.Hex(),
		"price":                 order.Price,
		"nonce":                 order.Nonce,
		"startTime":             big.NewInt(int64(order.StartTime)),
		"endTime":               big.NewInt(int64(order.EndTime)),
		"protocolFeeDiscounted": big.NewInt(int64(order.ProtocolFeeDiscounted)),
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
	txContext, err := contract.PrepareContext(targetChainConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer txContext.Close()

	contractAddr := common.HexToAddress(CONTRACT_ADDRESS)
	defoMarket, err := NewDefoMarketV2(contractAddr, txContext.EthClient)
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
	digestOnChain, err := defoMarket.GetOrderTypedDataHash(callOpts, order)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("chain signHash=%s", hexutil.Encode(digestOnChain[:]))

	// local sign ==================== sign by Admin
	signer := targetUserGroup.Admin
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
	r, s, v := contract.DecodeSignature(signature)
	order.R = r
	order.S = s
	order.V = v

	fmt.Println("\n=====================order debug=====================")
	spew.Dump(order)
	fmt.Println("=====================order debug end=====================")

	// execute order, for now, only one user
	// notice: it is presumed to be executed by admin
	// it is no matter who send the transaction，the key is the typed data signature.
	caller := targetUserGroup.Admin
	callerPrivateKey, err := crypto.HexToECDSA(caller.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	txOpts, err := txContext.GetTransactOpts(callerPrivateKey, order.Price)
	if err != nil {
		t.Fatal(err)
	}

	txExec, err := defoMarket.ExecuteEthscriptionOrder(txOpts, order)
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
	txContext, err := contract.PrepareContext(targetChainConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer txContext.Close()

	contractAddr := common.HexToAddress(domainStandard.VerifyingContract)
	defoMarket, err := NewDefoMarketV2(contractAddr, txContext.EthClient)
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

	// executed by admin
	caller := targetUserGroup.Admin
	callerPrivateKey, err := crypto.HexToECDSA(caller.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	txOpts, err := txContext.GetTransactOpts(callerPrivateKey, big.NewInt(0)) // no payment
	if err != nil {
		t.Fatal(err)
	}

	//recieve by seller
	seller := targetUserGroup.Seller
	recipient := common.HexToAddress(seller.Address)

	txExec, err := defoMarket.WithdrawEthscription(txOpts, we.EthscriptionId, recipient, we.Expiration, signature)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("\n=====================tx debug=====================")
	spew.Dump(txExec)
	fmt.Println("=====================tx debug end=====================")

	txContext.ShowTxDetail(txExec)
}

func TestMarket_sub(t *testing.T) {
	client, err := ethclient.Dial(targetChainConfig.WsGateWay)
	if err != nil {
		panic(err)
	}

	contractAddr := common.HexToAddress(CONTRACT_ADDRESS)
	contract, err := NewDefoMarketV2(contractAddr, client)

	tradeLog := make(chan *DefoMarketV2EthscriptionsProtocolTransferEthscriptionForPreviousOwner)
	sub, err := contract.WatchEthscriptionsProtocolTransferEthscriptionForPreviousOwner(nil, tradeLog, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err0 := <-sub.Err():
			if err0 != nil {
				fmt.Printf("err: %v\n", err0)
			}
		case eventLog := <-tradeLog:
			bs, _ := cjson.JSON.Marshal(eventLog)
			fmt.Printf("event Trade: %s\n", string(bs))
		}
	}
}
