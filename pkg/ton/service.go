package ton

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	address "github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"golang.org/x/crypto/ed25519"
	"math/big"
	"strings"
)

const (
	configMainNet = "https://ton.org/global.config.json"
	configTestNet = "https://ton-blockchain.github.io/testnet-global.config.json"
)

type TonService struct {
	isTest     bool
	connClient *liteclient.ConnectionPool
	apiClient  *ton.APIClient
	ctx        context.Context
}

func NewTonService(isTestNet bool) *TonService {
	client := liteclient.NewConnectionPool()

	configUrl := configMainNet
	if isTestNet {
		configUrl = configTestNet
	}
	fmt.Printf("TonService client config url=%s\n", configUrl)

	// get config
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), configUrl)
	if err != nil {
		fmt.Printf("fail to get config from url=%s\n", configUrl)
		panic(err)
	}

	// to bound all requests of operation to single node
	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		fmt.Printf("fail to add connections from url=%s\n", configUrl)
		panic(err)
	}

	apiClient := ton.NewAPIClient(client)
	apiClient.WithRetry(10)
	apiClient.SetTrustedBlockFromConfig(cfg)

	// bound all requests to single ton node
	ctx := client.StickyContext(context.Background())

	return &TonService{
		isTest:     isTestNet,
		connClient: client,
		apiClient:  apiClient,
		ctx:        ctx,
	}
}

type TonWallet struct {
	Address    string
	PrivateKey string
	Type       int
	Mnemonic   string
}

func (svc *TonService) CreateWallet(password string) (*TonWallet, error) {
	var seed []string
	hasPsw := len(password) > 0
	if hasPsw {
		seed = wallet.NewSeedWithPassword(password)
	} else {
		seed = wallet.NewSeed()
	}

	var w *wallet.Wallet
	var err error
	if hasPsw {
		w, err = wallet.FromSeedWithPassword(svc.apiClient, seed, password, wallet.V4R2)
	} else {
		w, err = wallet.FromSeed(svc.apiClient, seed, wallet.V4R2)
	}

	if err != nil {
		fmt.Printf("fail to create new ton wallet, error=%v\n", err)
		return nil, err
	}

	return assembleTonWallet(w, strings.Join(seed, " ")), nil
}

func (svc *TonService) GetWalletFromMnemonic(mnemonic string) (*TonWallet, error) {
	words := strings.Split(mnemonic, " ")

	w, err := wallet.FromSeed(svc.apiClient, words, wallet.V4R2)

	if err != nil {
		fmt.Printf("fail to create new ton wallet, error=%v\n", err)
		return nil, err
	}

	return assembleTonWallet(w, mnemonic), nil
}

func (svc *TonService) GetWalletFromPrivateKey(privateKey string) (*TonWallet, error) {
	bs, err := hexutil.Decode(privateKey)
	if err != nil {
		fmt.Printf("invalid private key (%s)\n", privateKey)
		return nil, err
	}

	priv := ed25519.PrivateKey(bs)
	w, err := wallet.FromPrivateKey(svc.apiClient, priv, wallet.V4R2)
	if err != nil {
		fmt.Printf("fail to create new ton wallet, error=%v\n", err)
		return nil, err
	}

	return assembleTonWallet(w, ""), nil
}

func assembleTonWallet(w *wallet.Wallet, mnemonic string) *TonWallet {
	address := w.WalletAddress()
	privateKey := hexutil.Encode(w.PrivateKey())
	return &TonWallet{
		Address:    address.String(),
		PrivateKey: privateKey,
		Type:       int(address.Type()),
		Mnemonic:   mnemonic,
	}
}

type TonBalance struct {
	Decimals int
	Value    *big.Int
}

func (svc *TonService) GetBalance(addr string) (*TonBalance, error) {
	block, err := svc.apiClient.CurrentMasterchainInfo(svc.ctx)
	if err != nil {
		fmt.Printf("fail to get master chain info, error=%v\n", err)
		return nil, err
	}

	tonAddr := address.MustParseAddr(addr)
	acc, err := svc.apiClient.WaitForBlock(block.SeqNo).GetAccount(svc.ctx, block, tonAddr)
	if err != nil {
		fmt.Printf("fail to get account info, error=%v\n", err)
		return nil, err
	}

	if !acc.IsActive {
		return nil, errors.New("account is not active")
	}

	balance := acc.State.Balance
	return &TonBalance{
		Decimals: 9,
		Value:    balance.Nano(),
	}, nil
}
