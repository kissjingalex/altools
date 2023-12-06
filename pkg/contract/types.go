package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var DefaultGasPrice = big.NewInt(20000000)

type ChainConfig struct {
	GateWay      string
	WsGateWay    string
	ContractAddr string
	UserAddr     string
	PrivateKey   string
	ChainId      *big.Int
}

type TxContext struct {
	EthClient    *ethclient.Client
	PrivateKey   *ecdsa.PrivateKey
	PublicKey    *ecdsa.PublicKey
	ChainId      *big.Int
	ContractAddr common.Address
	UserAddr     common.Address
	Ctx          context.Context
	Nonce        uint64
	GasPrice     *big.Int
}

func PrepareContext(config *ChainConfig) (*TxContext, error) {
	var client *ethclient.Client
	if config.WsGateWay != "" {
		var err error
		client, err = ethclient.Dial(config.WsGateWay)
		if err != nil {
			return nil, err
		}
		fmt.Printf("websocket gateway=%s\n", config.WsGateWay)
	} else {
		var err error
		client, err = ethclient.Dial(config.GateWay)
		if err != nil {
			return nil, err
		}
		fmt.Printf("gateway=%s\n", config.GateWay)
	}

	ctx := context.Background()
	contractAddr := common.HexToAddress(config.ContractAddr)
	fmt.Printf("contractAddr=%s\n", contractAddr)

	//signer
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	userAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("userAddr=%s\n", userAddr)

	// nonce 和 gasPrice
	nonce, err := client.PendingNonceAt(ctx, userAddr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("nonce=%d\n", nonce)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("suggestGasPrice=%s\n", gasPrice.String())
	if gasPrice.Cmp(DefaultGasPrice) > 0 {
		gasPrice = DefaultGasPrice
		fmt.Println("gasPrice is too big, reset to default(2000000)")
	}

	return &TxContext{
		EthClient:    client,
		PrivateKey:   privateKey,
		PublicKey:    publicKeyECDSA,
		ChainId:      config.ChainId,
		ContractAddr: contractAddr,
		UserAddr:     userAddr,
		Ctx:          ctx,
		Nonce:        nonce,
		GasPrice:     gasPrice,
	}, nil
}

func (c *TxContext) GetCallOpts() *bind.CallOpts {
	return &bind.CallOpts{
		Pending:     false,
		From:        c.UserAddr,
		BlockNumber: nil,
		Context:     c.Ctx,
	}
}

func (c *TxContext) GetSigner(value *big.Int) (*bind.TransactOpts, error) {
	chainId := c.ChainId
	auth, err := bind.NewKeyedTransactorWithChainID(c.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(c.Nonce))
	auth.Value = value             // in wei
	auth.GasLimit = uint64(300000) // in units, such 122558
	//auth.GasPrice = c.GasPrice
	auth.GasPrice = big.NewInt(1000000000) // 1gwei //本地anvil的suggestedGasPrice太小，需要设置大一点

	return auth, nil
}

// GetTransactOpts 注意这里的UserAddress已经更改，因此需要重新更新nonce
func (c *TxContext) GetTransactOpts(privateKey *ecdsa.PrivateKey, value *big.Int) (*bind.TransactOpts, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	txSenderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("txSenderAddress=%s\n", txSenderAddress)

	nonce, err := c.EthClient.PendingNonceAt(c.Ctx, txSenderAddress)
	if err != nil {
		return nil, err
	}
	fmt.Printf("txSender nonce=%d\n", nonce)

	chainId := c.ChainId
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value             // in wei
	auth.GasLimit = uint64(300000) // in units, such 122558
	//auth.GasPrice = c.GasPrice
	auth.GasPrice = big.NewInt(1000000000) // 1gwei //本地anvil的suggestedGasPrice太小，需要设置大一点

	return auth, nil
}

func (c *TxContext) ShowTxDetail(tx *types.Transaction) {
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("tx gasLimit: %d\n", tx.Gas())
	fmt.Printf("tx gasPrice: %s\n", tx.GasPrice().String())
	fmt.Printf("tx value: %s\n", tx.Value().String())
	fmt.Printf("tx cost: %s\n", tx.Cost().String())
	fmt.Printf("tx blob gasFeeCap: %s\n", tx.BlobGasFeeCap().String())
}

func (c *TxContext) Close() {
	if c.EthClient != nil {
		c.EthClient.Close()
	}
}
