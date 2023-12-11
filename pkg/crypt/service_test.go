package crypt

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestKeccak(t *testing.T) {
	content := "NewContract(address)"
	bs := crypto.Keccak256([]byte(content))

	t.Log(hexutil.Encode(bs))
}
