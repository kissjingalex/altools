package ton

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
	"testing"
)

func TestTon(t *testing.T) {
	d1 := decimal.NewFromInt(3)
	d2 := decimal.NewFromInt(1000000)
	d3 := d1.Mul(d2)
	t.Logf(d3.String())
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
	mnemonic := "lady midnight guard obvious hill budget car write tunnel lion harvest chicken rural brief human cereal use scatter service rapid typical honey cancel fruit"
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
	//addr := "EQAEirkgAaB3ZBMBbpLCseMrx5CojWp_u-T4VrNlSAeUQucq"
	addr := "UQBbLKS4Omj37NfpAQvznvUPq_HOfXCEex_PkLUcznKAH_kc"
	svc := NewTonService(false)
	rs, err := svc.GetBalance(addr)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(rs)
}
