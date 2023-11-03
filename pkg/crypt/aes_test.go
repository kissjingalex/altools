package crypt

import (
	"fmt"
	"testing"
)

func TestCryptoService_AESEncrypt(t *testing.T) {
	s := "aaaaa"
	bs := []byte(s)
	fmt.Printf("len=%d", len(bs))
}
