package chain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type TransferOpt struct {
	ToAddress string
	Value     *big.Int
	Data      []byte
}

type TransferPercentOpt struct {
	ToAddress string
	Percent   int // 0 - 100
}

type TransferResult struct {
	FromAddress string
	ToAddress   string
	Value       *big.Int
	TxHash      string
	Gas         uint64
	GasFee      *big.Int

	Signature *TxSignature
	SignBytes string
}

type TxSignature struct {
	V *big.Int
	R *big.Int
	S *big.Int
}

func (s *ChainService) Transfer(chainType string, senderPrivateKey string, opt *TransferOpt, isMock bool) (*TransferResult, error) {
	senderPriKey, err := crypto.HexToECDSA(senderPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := senderPriKey.PublicKey
	senderAddr := crypto.PubkeyToAddress(publicKey)
	ctx := context.Background()

	txContext, err := s.prepareTxContext(ctx, chainType, &senderAddr)
	if err != nil {
		return nil, err
	}

	//调整gasLimit
	txContext.GasLimit = DefaultGasLimitForTransfer

	toAddr := common.HexToAddress(opt.ToAddress)

	//查询Gas费用
	gas, err := txContext.Client.EstimateGas(txContext.Ctx, ethereum.CallMsg{
		From:  senderAddr,
		To:    &toAddr,
		Data:  opt.Data,
		Value: opt.Value,
	})
	gasFee := new(big.Int)
	gasFee.Mul(txContext.GasPrice, big.NewInt(int64(gas)))

	// create tx
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    txContext.Nonce,
		GasPrice: txContext.GasPrice,
		Gas:      txContext.GasLimit,
		To:       &toAddr,
		Data:     opt.Data,
		Value:    opt.Value,
	})

	signer := types.NewEIP155Signer(txContext.ChainId)
	signedTx, err := types.SignTx(tx, signer, senderPriKey)
	if err != nil {
		return nil, err
	}

	//获取签名bytes
	h := signer.Hash(tx)
	sigBytes, err := crypto.Sign(h[:], senderPriKey)

	vt, rt, st := signedTx.RawSignatureValues()
	sig := &TxSignature{
		V: vt,
		R: rt,
		S: st,
	}

	if !isMock {
		err = txContext.Client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			return nil, err
		}
	}

	return &TransferResult{
		FromAddress: senderAddr.Hex(),
		ToAddress:   opt.ToAddress,
		Value:       opt.Value,
		TxHash:      signedTx.Hash().Hex(),
		Gas:         gas,
		GasFee:      gasFee,
		Signature:   sig,
		SignBytes:   hexutil.Encode(sigBytes),
	}, nil
}
