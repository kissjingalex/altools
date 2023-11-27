package inscript

import (
	"altools/pkg/chain"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

type InscriptService struct {
}

var service = &InscriptService{}

func NewInscriptService() *InscriptService {
	return service
}

type MintResult struct {
	FromAddress string
	ToAddress   string
	Data        string

	TxHash string
	Gas    uint64
	GasFee *big.Int

	Signature *chain.TxSignature
}

func (svc *InscriptService) Mint(chainType string, senderPrivateKey string, toAddress string, data string) (*MintResult, error) {
	dataBytes, err := hexutil.Decode(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("mint content=%s\n", string(dataBytes))

	chainSvc := chain.NewChainService()
	tr, err := chainSvc.Transfer(chainType, senderPrivateKey, &chain.TransferOpt{
		ToAddress: toAddress,
		Value:     big.NewInt(0),
		Data:      dataBytes,
	})

	if err != nil {
		return nil, err
	}

	return &MintResult{
		FromAddress: tr.FromAddress,
		ToAddress:   tr.ToAddress,
		Data:        data,
		TxHash:      tr.TxHash,
		Gas:         tr.Gas,
		GasFee:      tr.GasFee,
		Signature:   tr.Signature,
	}, nil
}
