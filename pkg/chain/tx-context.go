package chain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type TxContext struct {
	Ctx      context.Context
	Client   *ethclient.Client
	ChainId  *big.Int
	Nonce    uint64
	GasPrice *big.Int
	GasLimit uint64
}

// 为转移eth准备tx的context，注意gasLimit设置比较低
func (s *ChainService) prepareTxContext(ctx context.Context, chain string, senderAddr *common.Address) (*TxContext, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	client, err := s.GetEthClient(chain)
	if err != nil {
		return nil, err
	}

	chainConfig := GetWalletChainConfig(chain)
	chainId := chainConfig.ChainId

	// nonce 和 gasPrice
	nonce, err := client.PendingNonceAt(ctx, *senderAddr)
	if err != nil {
		fmt.Printf("Fail to get nonce. error=%v", err)
	}
	//fmt.Printf("nonce=%d\n", nonce)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Printf("Fail to get gasPrice. error=%v", err)
	}
	//fmt.Printf("suggestGasPrice=%s\n", gasPrice.String())
	//if gasPrice.Cmp(DefaultGasPrice) > 0 {
	//	gasPrice = DefaultGasPrice
	//	//log.Info("ChainService").Msgf("gasPrice is too big, reset to default(2000000)")
	//}

	return &TxContext{
		Ctx:      ctx,
		Client:   client,
		ChainId:  chainId,
		Nonce:    nonce,
		GasPrice: gasPrice,
		GasLimit: DefaultGasLimit,
	}, nil
}

// GetSigner 用于执行合约
func (tc *TxContext) GetSigner(privateKeyStr string, value *big.Int) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	chainId := tc.ChainId
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(tc.Nonce))
	auth.Value = value              // in wei
	auth.GasLimit = DefaultGasLimit // in units, such 122558
	auth.GasPrice = tc.GasPrice

	return auth, nil
}
