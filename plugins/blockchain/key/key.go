package key

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetPrivateKeyString(priv *ecdsa.PrivateKey) string {
	return hexutil.Encode(crypto.FromECDSA(priv))
}

func ToPrivateKeyFromString(privKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privKey)
}
