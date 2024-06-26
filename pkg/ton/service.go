package ton

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"golang.org/x/crypto/ed25519"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	//configMainNet = "https://ton.org/global.config.json"
	configMainNet = "https://ton-blockchain.github.io/global.config.json"
	configTestNet = "https://ton-blockchain.github.io/testnet-global.config.json"
)

type TonService struct {
	isTest     bool
	connClient *liteclient.ConnectionPool
	apiClient  *ton.APIClient
	ctx        context.Context
}

func getConfigFromUrl(isTestNet bool) *liteclient.GlobalConfig {
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

	return cfg
}

func getConfigFromLocal(isTestNet bool) *liteclient.GlobalConfig {
	configData := clientConfigOfficial
	if isTestNet {
		configData = clientConfigTestnet
	}

	config := &liteclient.GlobalConfig{}
	if err := json.Unmarshal([]byte(configData), config); err != nil {
		return nil
	}
	return config
}

func NewTonService(isTestNet bool) *TonService {
	client := liteclient.NewConnectionPool()

	cfg := getConfigFromLocal(isTestNet)

	// to bound all requests of operation to single node
	err := client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		fmt.Printf("fail to add connections")
		panic(err)
	}

	apiClient := ton.NewAPIClient(client)
	apiClient.WithRetry(10)
	apiClient.SetTrustedBlockFromConfig(cfg)

	// bound all requests to single ton node
	//ctx := client.StickyContext(context.Background())
	ctx := context.Background()

	return &TonService{
		isTest:     isTestNet,
		connClient: client,
		apiClient:  apiClient,
		ctx:        ctx,
	}
}

type AddressInfo struct {
	Type         int   // 2 standard
	Chain        int32 // 0 base chain
	Data         string
	IsTestNet    bool
	IsBounceable bool
	BAddress     string
	UBAddress    string
}

func CheckConfig() error {
	config := &liteclient.GlobalConfig{}

	configData, err := os.ReadFile("./config.json")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(configData, config); err != nil {
		return err
	}

	for _, server := range config.Liteservers {
		ipStr := intToIP4(server.IP)
		fmt.Printf("liteServer(%d) = %s\n", server.IP, ipStr)
	}

	return nil
}

func intToIP4(ipInt int64) string {
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func ParseAddress(addr string, isRaw bool) *AddressInfo {
	var tonAddr *address.Address
	if isRaw {
		tonAddr = address.MustParseRawAddr(addr)
	} else {
		tonAddr = address.MustParseAddr(addr)
	}

	addrType := tonAddr.Type()
	chain := tonAddr.Workchain()
	data := hex.EncodeToString(tonAddr.Data())
	info := &AddressInfo{
		Type:         int(addrType),
		Chain:        chain,
		Data:         data,
		IsTestNet:    tonAddr.IsTestnetOnly(),
		IsBounceable: tonAddr.IsBounceable(),
		BAddress:     "",
		UBAddress:    "",
	}

	if info.IsBounceable {
		info.BAddress = tonAddr.String()

		tonAddr.SetBounce(false)
		info.UBAddress = tonAddr.String()
	} else {
		info.UBAddress = tonAddr.String()

		tonAddr.SetBounce(true)
		info.BAddress = tonAddr.String()
	}

	return info
}

type Wallet struct {
	Address    string
	PrivateKey string
	Type       int
	Mnemonic   string
}

func (svc *TonService) CreateWallet(password string) (*Wallet, error) {
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
		//w, err = wallet.FromSeed(svc.apiClient, seed, wallet.V4R2)
		w, err = wallet.FromSeed(nil, seed, wallet.V4R2)
	}

	if err != nil {
		fmt.Printf("fail to create new ton wallet, error=%v\n", err)
		return nil, err
	}

	return assembleTonWallet(w, strings.Join(seed, " ")), nil
}

func (svc *TonService) GetWalletFromMnemonic(mnemonic string) (*Wallet, error) {
	words := strings.Split(mnemonic, " ")

	w, err := wallet.FromSeed(svc.apiClient, words, wallet.V4R2)

	if err != nil {
		fmt.Printf("fail to create new ton wallet, error=%v\n", err)
		return nil, err
	}

	return assembleTonWallet(w, mnemonic), nil
}

func (svc *TonService) GetWalletFromPrivateKey(privateKey string) (*Wallet, error) {
	w, err := svc.getRawWalletFromPrivateKey(privateKey)
	if err != nil {
		fmt.Printf("fail to get raw ton wallet, error=%v\n", err)
		return nil, err
	}

	return assembleTonWallet(w, ""), nil
}

func (svc *TonService) getRawWalletFromPrivateKey(privateKey string) (*wallet.Wallet, error) {
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

	return w, nil
}

func assembleTonWallet(w *wallet.Wallet, mnemonic string) *Wallet {
	addr := w.WalletAddress()
	privateKey := hexutil.Encode(w.PrivateKey())
	return &Wallet{
		Address:    addr.String(),
		PrivateKey: privateKey,
		Type:       int(addr.Type()),
		Mnemonic:   mnemonic,
	}
}

type Balance struct {
	Decimals int
	Value    *big.Int
}

func (svc *TonService) GetBalance(addr string) (*Balance, error) {
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
	return &Balance{
		Decimals: 9,
		Value:    balance.Nano(),
	}, nil
}

type TxSender struct {
	Address    string
	PrivateKey string
}

type TransferResult struct {
	TxHash     string
	BlockSeqNo uint32
	Balance    *big.Int
}

func (svc *TonService) TransferTon(sender *TxSender, userAddr string, value *big.Int) (*TransferResult, error) {
	w, err := svc.getRawWalletFromPrivateKey(sender.PrivateKey)
	if err != nil {
		return nil, err
	}

	addr := address.MustParseAddr(userAddr)

	// if destination wallet is not initialized (or you don't care)
	// you should set bounce to false to not get money back.
	// If bounce is true, money will be returned in case of not initialized destination wallet or smart-contract error
	bounce := false

	transfer, err := w.BuildTransfer(addr, tlb.MustFromNano(value, 9), bounce, "")
	if err != nil {
		fmt.Printf("Transfer error, %v\n", err)
		return nil, err
	}

	tx, block, err := w.SendWaitTransaction(svc.ctx, transfer)
	if err != nil {
		fmt.Printf("SendWaitTransaction error, %v\n", err)
		return nil, err
	}
	txHash := base64.StdEncoding.EncodeToString(tx.Hash)
	blockSeqNo := block.SeqNo

	balance, err := w.GetBalance(svc.ctx, block)
	if err != nil {
		fmt.Printf("GetBalance error, %v\n", err)
		return nil, err
	}

	return &TransferResult{
		TxHash:     txHash,
		BlockSeqNo: blockSeqNo,
		Balance:    balance.Nano(),
	}, nil
}

type TransferBody struct {
	InMsgHash string
	Body      []byte
}

func (svc *TonService) BuildTransferBody(sender *TxSender, userAddr string, value *big.Int) (*TransferBody, error) {
	w, err := svc.getRawWalletFromPrivateKey(sender.PrivateKey)
	if err != nil {
		return nil, err
	}

	addr := address.MustParseAddr(userAddr)

	// if destination wallet is not initialized (or you don't care)
	// you should set bounce to false to not get money back.
	// If bounce is true, money will be returned in case of not initialized destination wallet or smart-contract error
	bounce := false

	transfer, err := w.BuildTransfer(addr, tlb.MustFromNano(value, 9), bounce, "")
	if err != nil {
		fmt.Printf("Transfer error, %v\n", err)
		return nil, err
	}

	ext, err := w.BuildExternalMessageForMany(svc.ctx, []*wallet.Message{transfer})
	if err != nil {
		fmt.Printf("Transfer error, %v\n", err)
		return nil, err
	}
	inMsgHash := ext.Body.Hash()

	req, err := tlb.ToCell(ext)
	if err != nil {
		fmt.Printf("Transfer error, %v\n", err)
		return nil, err
	}

	body := req.ToBOCWithFlags(false)

	return &TransferBody{
		InMsgHash: base64.StdEncoding.EncodeToString(inMsgHash),
		Body:      body,
	}, nil
}
