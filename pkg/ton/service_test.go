package ton

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

type TestAccount struct {
	Address    string
	PrivateKey string
	Mnemonic   string
}

var testSender = &TestAccount{
	Address:    "UQDjpLi6XzPP3cJdK1PY5NEQA3CEKpo4S8eT1WP6r-JG2Z7j",
	PrivateKey: "0xd63b0342887ab9be63277a86ede7f8847227c8d8e4865e2f1b28748db639d678d2bce3d730cbfc91a692444614668868a3a51240d147cc6fc4972781c508e1da",
	Mnemonic:   "usual fury grant armed collect already jazz increase service timber smoke camera riot tree carpet junior staff twist alpha galaxy armor else indicate isolate",
}

var testUser = &TestAccount{
	Address:    "UQBaEAp6HF605KE3FM20Mx0IbkXJwJGC2BQxeGfAfulunyhg",
	PrivateKey: "0x7fdd211ad50deb8b6e54c21e2143a79793f075d2e12214a6f0a6151ce2def1fa3cf52676c5c957204a1198e93cc90180c35bce8bbd3e4b98ced45017209ba34a",
	Mnemonic:   "impose inflict timber goat uphold trend rare they control cost foot trip give timber bottom token bike cake certain enroll aunt struggle forget organ",
}

func TestTon(t *testing.T) {
	d1 := decimal.NewFromInt(3)
	d2 := decimal.NewFromInt(1000000)
	d3 := d1.Mul(d2)
	t.Logf(d3.String())
}

func TestTonService_ParseAddress(t *testing.T) {
	addr := "UQBOpzbphvGUuFf_gFSG4szRGgIgOcrfotk6_vmo2Y33RI6c"
	info := ParseAddress(addr)
	spew.Dump(info)
}

func TestTonService_CreateWallet(t *testing.T) {
	svc := NewTonService(true)
	w, err := svc.CreateWallet("")
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(w) // UQDT66bcJ1bwafVSIeQtrMjA_J6T8k5FvELV44QOyz2DSAq5
}

func TestTonService_GetWalletFromMnemonic(t *testing.T) {
	mnemonic := "push device caught chronic amused glove capable stadium almost dizzy popular morning liberty saddle apology element repeat update smart yellow decorate machine accuse toast"
	svc := NewTonService(true)
	w, err := svc.GetWalletFromMnemonic(mnemonic)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(w)
}

func TestTonService_GetWalletFromPrivateKey(t *testing.T) {
	privateKey := "0x654cb8f95dda667e4a74ea95e730e60f17014c569874675412960df8f10a6b85b9dd6891cab6c36f8546ae74eb2fb3fd58c73664b84c7cf6ce7eab7b6f5a2bf1"
	svc := NewTonService(true)
	w, err := svc.GetWalletFromPrivateKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(w)
}

func TestTonService_GetBalance(t *testing.T) {
	addr := "UQDjpLi6XzPP3cJdK1PY5NEQA3CEKpo4S8eT1WP6r-JG2Z7j"
	//addr := "UQBOpzbphvGUuFf_gFSG4szRGgIgOcrfotk6_vmo2Y33RI6c"
	svc := NewTonService(true)
	rs, err := svc.GetBalance(addr)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(rs)
}

func TestTonService_TransferTon(t *testing.T) {
	sender := &TxSender{
		Address:    testSender.Address,
		PrivateKey: testSender.PrivateKey,
	}
	svc := NewTonService(true)
	rs, err := svc.TransferTon(sender, testUser.Address, big.NewInt(100000000))
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(rs)
}
