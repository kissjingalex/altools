package ton2

import (
	"context"
	"github.com/tonkeeper/tonapi-go"
	"math/big"
)

const (
	TonApiURL        = "https://tonapi.io"
	TestnetTonApiURL = "https://testnet.tonapi.io"

	ApiKey = "AF3LFDMKHZPLEDIAAAAO3SGQ2OXXX27YNF7ZOEJPHQBZMCVRJ4YWRNE73UGDLMY55PMFDOQ"
)

type TonService struct {
	isTest bool
	client *tonapi.Client
}

func NewTonService(isTestNet bool) *TonService {
	serverUrl := tonapi.TonApiURL
	if isTestNet {
		serverUrl = tonapi.TestnetTonApiURL
	}
	client, err := tonapi.NewClient(serverUrl, tonapi.WithToken(ApiKey))
	if err != nil {
		panic("fail to create ton client")
	}

	return &TonService{
		isTest: isTestNet,
		client: client,
	}
}

type Balance struct {
	Decimals int
	Value    *big.Int
}

func (svc *TonService) GetBalance(addr string) (*Balance, error) {
	ctx := context.Background()
	account, err := svc.client.GetAccount(ctx, tonapi.GetAccountParams{AccountID: addr})
	if err != nil {
		return nil, err
	}
	return &Balance{
		Decimals: 9,
		Value:    big.NewInt(account.Balance),
	}, nil
}

type TxInfo struct {
	Hash     string
	Block    string
	Success  bool
	TotalFee *big.Int
}

func (svc *TonService) GetTransaction(txHash string) (*TxInfo, error) {
	ctx := context.Background()
	tx, err := svc.client.GetBlockchainTransaction(ctx, tonapi.GetBlockchainTransactionParams{TransactionID: txHash})
	if err != nil {
		return nil, err
	}

	return &TxInfo{
		Hash:     tx.Hash,
		Block:    tx.Block,
		Success:  tx.Success,
		TotalFee: big.NewInt(tx.TotalFees),
	}, nil
}

func (svc *TonService) SendMessage(payload []byte) error {
	ctx := context.Background()
	_, err := svc.client.SendMessage(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
