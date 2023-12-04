package chain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"testing"
)

const (
	domainSign = "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
)

var TYPE_HASH = crypto.Keccak256([]byte(domainSign))

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
	"TradeOrder": {
		{
			Name: "signer",
			Type: "address",
		},
		{
			Name: "itemId",
			Type: "bytes32",
		},
		{
			Name: "price",
			Type: "uint256",
		},
	},
}

const primaryType = "TradeOrder"

var domainStandard = apitypes.TypedDataDomain{
	Name:              "DefoMarket",
	Version:           "1",
	ChainId:           math.NewHexOrDecimal256(31337),
	VerifyingContract: "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9",
	Salt:              "",
}

const SIGNER_PRIVER_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

var messageStandard = map[string]interface{}{
	"signer": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	"itemId": "0xdb9580c13554659ea1fb7530cdb5c3049f3af68789986433443483f7b9293880",
	"price":  big.NewInt(100),
}

var typedData = apitypes.TypedData{
	Types:       typesStandard,
	PrimaryType: primaryType,
	Domain:      domainStandard,
	Message:     messageStandard,
}

func TestSign712(t *testing.T) {
	t1 := hexutil.Encode(TYPE_HASH)
	fmt.Printf("t1=%s\n", t1)

	t2 := typedData.EncodeType("EIP712Domain")
	fmt.Printf("t2=%s\n", string(t2))

	t3 := typedData.TypeHash("EIP712Domain")
	fmt.Printf("t3=%s\n", t3)

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("domainSeparator=%s\n", domainSeparator)
}

func TestSign712Full(t *testing.T) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("domainSeparator=%s\n", domainSeparator)

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		t.Fatal(err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	signHash := crypto.Keccak256(rawData)
	t.Logf("sigHash=%s\n", hexutil.Encode(signHash))

	privateKey, err := crypto.HexToECDSA(SIGNER_PRIVER_KEY)
	if err != nil {
		t.Fatal(err)
	}

	signature, err := crypto.Sign(signHash, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("signature=%s", hexutil.Encode(signature))
}

func TestHashStruct(t *testing.T) {
	typeHash := typedData.TypeHash("TradeOrder")
	t.Logf("typeHash=%s\n", typeHash)

	hash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		t.Fatal(err)
	}
	mainHash := fmt.Sprintf("0x%s", common.Bytes2Hex(hash))
	t.Logf("main hash=%s", mainHash)

	hash, err = typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		t.Error(err)
	}
	domainHash := fmt.Sprintf("0x%s", common.Bytes2Hex(hash))
	t.Logf("domain hash=%s", domainHash)
}
