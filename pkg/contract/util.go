package contract

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

// from types/transaction_signing.go
func decodeSignature(sig []byte) (r, s [32]byte, v uint8) {
	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}
	//r = new(big.Int).SetBytes(sig[:32])
	//s = new(big.Int).SetBytes(sig[32:64])
	//v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	v = sig[64] + 27

	return r, s, v
}
