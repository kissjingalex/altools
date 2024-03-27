package tron

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestTronService_GetTransactionInfo(t *testing.T) {
	txHash := "acfff45c5b25e004d56b810f5bb1fbf3a9294eadee723715f616f94a52ea344b"

	rs, err := GetTransactionInfo(txHash)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(rs)
}
