package chain

import (
	"altools/plugins/blockchain/util"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type ChainService struct {
}

var service = &ChainService{}

func NewChainService() *ChainService {
	return service
}

// 支持多链
var clientMap = make(map[string]*ethclient.Client)

// 交易相关的常量
var gwei = big.NewInt(1000000000)

var DefaultGasPrice = big.NewInt(20000000) // 0.02 Gwei
var DefaultGasLimit = uint64(300000)

var DefaultGasPriceForTransfer = big.NewInt(2000000) // 0.002 Gwei
var DefaultGasLimitForTransfer = uint64(30000)       // 21000  27329（转多钱钱包）

func (svc *ChainService) Release() {
	if len(clientMap) == 0 {
		return
	}

	for _, client := range clientMap {
		client.Close()
	}
}

func (svc *ChainService) GetEthClient(chain string) (client *ethclient.Client, err error) {
	wc := GetWalletChainConfig(chain)
	if wc == nil {
		return nil, errors.New("not supported chain")
	}

	client = clientMap[chain]
	if client == nil {
		client, err = ethclient.Dial(wc.EndPoint)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("fail to connect endpoint=%s", wc.EndPoint))
		}
		clientMap[chain] = client
	}

	return client, nil
}

// CheckAddress 解析地址是否合法，全小写的地址转成正式地址
func (svc *ChainService) CheckAddress(address string) (string, error) {
	if len(address) < 0 {
		return "", errors.New("empty string")
	}

	isValid := util.IsValidAddress(address)
	if !isValid {
		return "", errors.New("invalid address")
	}

	addr := common.HexToAddress(address)
	return addr.Hex(), nil
}

func (svc *ChainService) GetBalance(address string, chain string) (*big.Int, error) {
	client, err := svc.GetEthClient(chain)
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(address)
	at, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return nil, err
	}

	return at, nil
}

func (svc *ChainService) GetPendingNonce(address string, chain string) (uint64, error) {
	client, err := svc.GetEthClient(chain)
	if err != nil {
		return 0, err
	}

	addr := common.HexToAddress(address)
	n, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (svc *ChainService) GetSuggestGasPrice(chain string) (*big.Int, error) {
	client, err := svc.GetEthClient(chain)
	if err != nil {
		return nil, err
	}

	n, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	return n, nil
}
