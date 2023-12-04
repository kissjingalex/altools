package contract

import (
	"altools/plugins/blockchain/util"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

const (
	DefoSign_Contract = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"
	SIGNER_PRIVER_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

var signMessageStandard = map[string]interface{}{
	"signer": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	"itemId": "0xdb9580c13554659ea1fb7530cdb5c3049f3af68789986433443483f7b9293880",
	"price":  big.NewInt(100),
}

func TestBase_Balance(t *testing.T) {
	ctx, err := prepareContext(localChainConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	balance, err := ctx.EthClient.BalanceAt(ctx.Ctx, ctx.UserAddr, nil)
	if err != nil {
		t.Fatal(err)
	}
	ethValue := util.ToEther(balance)
	fmt.Printf("balance(ether)=%s\n", ethValue.String())
}

func TestDefoSign_OrderHash(t *testing.T) {
	ctx, err := prepareContext(localChainConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	contractAddr := common.HexToAddress(DefoSign_Contract)
	defoSign, err := NewDefoSign(contractAddr, ctx.EthClient)
	if err != nil {
		t.Fatal(err)
	}

	callOpts := ctx.GetCallOpts()

	itemIdBytes, err := hexutil.Decode(signMessageStandard["itemId"].(string))
	if err != nil {
		t.Fatal(err)
	}
	itemId := [32]byte{}
	copy(itemId[:], itemIdBytes) //切片转数组

	order := TradeTypesTradeOrder{
		Signer: common.HexToAddress(signMessageStandard["signer"].(string)),
		ItemId: itemId,
		Price:  signMessageStandard["price"].(*big.Int),
	}

	hash, err := defoSign.GetOrderHash(callOpts, order)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("orderHash=%s", hexutil.Encode(hash[:]))

	signHash, err := defoSign.GetOrderTypedDataHash(callOpts, order)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("orderTypedDataHash=%s", hexutil.Encode(signHash[:])) //签名的hash

	//签证签名
	privateKey, err := crypto.HexToECDSA(SIGNER_PRIVER_KEY)
	if err != nil {
		t.Fatal(err)
	}

	signature, err := crypto.Sign(signHash[:], privateKey)
	if err != nil {
		t.Fatal(err)
	}
	r, s, v := decodeSignature(signature)
	order.R = r
	order.S = s
	order.V = v
	spew.Dump(order)

	digest, err := defoSign.VerifyOrder(callOpts, order)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("digest=%s", hexutil.Encode(digest[:]))
}

func TestDefiSign_GetDomainSeparator(t *testing.T) {
	ctx, err := prepareContext(localChainConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	contractAddr := common.HexToAddress(DefoSign_Contract)
	defoSign, err := NewDefoSign(contractAddr, ctx.EthClient)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := defoSign.GetDomainSeparator(ctx.GetCallOpts())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("domainSeparator=%s", hexutil.Encode(bytes[:]))
}
