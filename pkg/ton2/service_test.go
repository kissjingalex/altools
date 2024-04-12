package ton2

import (
	"github.com/davecgh/go-spew/spew"
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

func TestTonService_GetBalance(t *testing.T) {
	addr := testSender.Address
	svc := NewTonService(true)
	rs, err := svc.GetBalance(addr)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(rs)
}

func TestTonService_GetTransaction(t *testing.T) {
	txHash := "U7m2J5-xN3wDLxq84LRK4McEqm1NG5gjQVidF3uvmto="
	svc := NewTonService(true)
	rs, err := svc.GetTransaction(txHash)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(rs)
}
