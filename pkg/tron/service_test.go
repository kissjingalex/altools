package tron

import (
	"encoding/json"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"math/big"
	"testing"
)

type TestAccount struct {
	Address    string
	PrivateKey string
	Mnemonic   string
}

var testSender = &TestAccount{
	Address:    "TVJkvDov1zJwE4KgjbVRXBSGp55qv9ULjF",
	PrivateKey: "2967ab54e58d0a3d2f53a134eb6f90c44612bde56510ae745bf401fd13fc547e",
	Mnemonic:   "clinic cube hard medal recycle engage pride wheat abandon gorilla stick meat matrix exclude fence car coconut develop naive tuna donkey twenty motion brief",
}

var testUser = &TestAccount{
	Address:    "TDBCbsoB3RnuRQsuqz1LKa2mKKkk62jU57",
	PrivateKey: "9d165f6f4d5e5256371a124c526697978f60f29b0d3948c8af9d5a477f5b3818",
	Mnemonic:   "exist ball front gallery relax vehicle average divide cool industry crucial grant rival account stool side wash key hire hover bone entry envelope animal",
}

func TestAddress(t *testing.T) {
	s := NewTronService()
	//ethAddr := "0x418aa06d692d0f98f8d3c24cf9901371b021ef555b"
	ethAddr := "0x8aa06d692d0f98f8d3c24cf9901371b021ef555b"
	tronAddr, err := s.ToTronAddress(ethAddr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("tronAddr = %s", tronAddr)

	ethAddr, err = s.ToEthAddress(tronAddr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ethAddr = %s", ethAddr)
}

func TestTronService_GetAccontInfo(t *testing.T) {
	addr := "TRkyRH7YVH59MRxKCpHoTG3M2dGo3Cxccb"
	svc := NewTronService()
	rs, err := svc.GetAccountInfo(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rs = %v", rs)
}

func TestTronService_GetUsdtBalance(t *testing.T) {
	addr := "TRkyRH7YVH59MRxKCpHoTG3M2dGo3Cxccb"
	svc := NewTronService()
	rs, err := svc.GetUsdtBalance(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rs = %v", rs)
}

func TestTronService_TransferUsdt(t *testing.T) {
	svc := NewTronService()

	sender := &TxSender{
		Address:    testSender.Address,
		PrivateKey: testSender.PrivateKey,
	}
	rs, err := svc.TransferUsdt(sender, testUser.Address, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}

	asJSON, _ := json.Marshal(rs)
	fmt.Println(common.JSONPrettyFormat(string(asJSON)))
	//t.Logf("rs = %v", rs)
}

func TestTronService_TransferTrx(t *testing.T) {
	svc := NewTronService()

	sender := &TxSender{
		Address:    testSender.Address,
		PrivateKey: testSender.PrivateKey,
	}
	rs, err := svc.TransferTrx(sender, testUser.Address, big.NewInt(10*1000000))
	if err != nil {
		t.Fatal(err)
	}

	asJSON, _ := json.Marshal(rs)
	fmt.Println(common.JSONPrettyFormat(string(asJSON)))
}
