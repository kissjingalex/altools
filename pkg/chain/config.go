package chain

import (
	"fmt"
	"math/big"
)

const (
	DefaultChain = "Ethereum"

	pathFormat = "m/44'/%d'/0'/0/%d"
)

type WalletChainConfig struct {
	Chain    string
	ChainId  *big.Int
	CoinType int32
	Index    int32
	EndPoint string
}

// endpoint可以查询 https://chainlist.org/
// WCConfigList coinType https://github.com/satoshilabs/slips/blob/master/slip-0044.md
var WCConfigList = []*WalletChainConfig{
	{
		Chain:    "Base",
		ChainId:  big.NewInt(8453),
		CoinType: 60,
		Index:    0,
		EndPoint: "https://mainnet.base.org",
	},
	{
		Chain:    "Base_Goerli",
		ChainId:  big.NewInt(84531),
		CoinType: 60,
		Index:    0, // 1,
		EndPoint: "https://goerli.base.org",
	},
	{
		Chain:    "Ethereum",
		ChainId:  big.NewInt(1),
		CoinType: 60,
		Index:    0, // 2,
		EndPoint: "https://eth.llamarpc.com",
	},
	{
		Chain:    "Goerli",
		ChainId:  big.NewInt(5),
		CoinType: 60,
		Index:    0, // 2,
		EndPoint: "https://api.zan.top/node/v1/eth/goerli/public",
	},
	{
		Chain:    "Bitcoin",
		CoinType: 0,
		Index:    0,
	},
	{
		Chain:    "TRON",
		CoinType: 195,
		Index:    0,
	},
}

var WCConfigMap = make(map[string]*WalletChainConfig)

func init() {
	for _, config := range WCConfigList {
		WCConfigMap[config.Chain] = config
	}
}

func GetWalletChainConfig(chain string) *WalletChainConfig {
	return WCConfigMap[chain]
}

func (c *WalletChainConfig) GenWalletPath() string {
	return fmt.Sprintf(pathFormat, c.CoinType, c.Index)
}
