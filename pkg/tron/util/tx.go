package util

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/protobuf/proto"
)

func SignTx(privateKey string, tx *core.Transaction) (*core.Transaction, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}

	// data hash
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	signature, err := crypto.Sign(hash, privKey)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)

	return tx, nil
}
