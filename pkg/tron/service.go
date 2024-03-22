package tron

import (
	"altools/pkg/tron/util"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common/decimals"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"math/big"
)

const (
	nodeOfficial = "grpc.trongrid.io:50051"
	nodeNile     = "grpc.nile.trongrid.io:50051"

	apiKey = "c4f3fefc-1a73-46d1-9165-238c00c2cf0f"

	connOptsWithTLS = false

	usdtContractOnNile = "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj"

	defaultFeeLimit int64 = 50 * 1000000 // 1trx = 1e6 sun
)

type TronService struct {
	client *client.GrpcClient
}

func NewTronService() *TronService {
	service := &TronService{}
	service.client = newGrpcClient()
	return service
}

func newGrpcClient() *client.GrpcClient {
	conn := client.NewGrpcClient(nodeNile)

	opts := make([]grpc.DialOption, 0)
	if connOptsWithTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// set API
	conn.SetAPIKey(apiKey)

	if err := conn.Start(opts...); err != nil {
		panic(err)
		return nil
	}

	fmt.Printf("Success to dial to node[%s]\n", conn.Address)
	return conn
}

func (svc *TronService) ToTronAddress(addr string) (string, error) {
	if len(addr) == 0 {
		return "", errors.New("empty address")
	}
	tronAddr, err := util.ToTronAddressBytes(addr)
	if err != nil {
		return "", err
	}
	return tronAddr.String(), nil
}

func (svc *TronService) ToEthAddress(addr string) (string, error) {
	if len(addr) == 0 {
		return "", errors.New("empty address")
	}
	tronAddr, err := address.Base58ToAddress(addr)
	if err != nil {
		return "", err
	}
	ethAddr := hexutil.Encode(tronAddr.Bytes()[1:])
	return common.HexToAddress(ethAddr).Hex(), nil
}

type BalanceResult struct {
	TokenDecimals *big.Int
	Symbol        string
	Value         *big.Int
	Amount        *big.Float
}

func (svc *TronService) GetUsdtBalance(userAddr string) (*BalanceResult, error) {
	contractAddr := usdtContractOnNile

	// get contract decimals if any
	tokenDecimals, err := svc.client.TRC20GetDecimals(contractAddr)
	if err != nil {
		tokenDecimals = big.NewInt(0)
	}

	// get contract decimals if any
	symbol, err := svc.client.TRC20GetSymbol(contractAddr)
	if err != nil {
		symbol = ""
	}

	value, err := svc.client.TRC20ContractBalance(userAddr, contractAddr)
	if err != nil {
		return nil, err
	}

	amount := decimals.RemoveDecimals(value, tokenDecimals.Int64())

	return &BalanceResult{
		TokenDecimals: tokenDecimals,
		Symbol:        symbol,
		Value:         value,
		Amount:        amount,
	}, nil
}

type TransferReceipt struct {
	Fee               int64
	EnergyFee         int64
	EnergyUsage       int64
	OriginEnergyUsage int64
	EnergyUsageTotal  int64
	NetFee            int64
	NetUsage          int64
}

type TransferResult struct {
	TxHash          string
	BlockNumber     int64
	Message         string
	ContractAddress string
	Success         bool
	ResMessage      string
	Receipt         *TransferReceipt
}

func (svc *TronService) TransferTrx(sender *TxSender, userAddr string, value *big.Int) (*TransferResult, error) {
	tx, err := svc.client.Transfer(sender.Address, userAddr, value.Int64())
	if err != nil {
		return nil, err
	}

	controller := NewTronController(svc.client, sender, tx.Transaction)
	if err = controller.ExecuteTransaction(); err != nil {
		return nil, err
	}

	rs := assembleTransferResult(controller, tx, "")

	return rs, nil
}

func (svc *TronService) TransferUsdt(sender *TxSender, userAddr string, value *big.Int) (*TransferResult, error) {
	contractAddr := usdtContractOnNile

	tx, err := svc.client.TRC20Send(sender.Address, userAddr, contractAddr, value, defaultFeeLimit)
	if err != nil {
		return nil, err
	}

	controller := NewTronController(svc.client, sender, tx.Transaction)
	if err = controller.ExecuteTransaction(); err != nil {
		return nil, err
	}

	rs := assembleTransferResult(controller, tx, contractAddr)

	return rs, nil
}

func assembleTransferResult(controller *Controller, tx *api.TransactionExtention, contractAddr string) *TransferResult {
	rs := &TransferResult{
		TxHash:          hexutil.Encode(tx.GetTxid()),
		BlockNumber:     0,
		Message:         string(controller.Result.Message),
		ContractAddress: contractAddr,
		Success:         controller.GetResultError() == nil,
		ResMessage:      "",
		Receipt:         nil,
	}

	receipt := controller.Receipt
	if receipt != nil {
		rs.BlockNumber = receipt.BlockNumber
		rs.ResMessage = string(controller.Receipt.ResMessage)

		txReceipt := &TransferReceipt{
			Fee:               receipt.Fee,
			EnergyFee:         receipt.Receipt.EnergyFee,
			EnergyUsage:       receipt.Receipt.EnergyUsage,
			OriginEnergyUsage: receipt.Receipt.OriginEnergyUsage,
			EnergyUsageTotal:  receipt.Receipt.EnergyUsageTotal,
			NetFee:            receipt.Receipt.NetFee,
			NetUsage:          receipt.Receipt.NetUsage,
		}
		rs.Receipt = txReceipt
	}

	return rs
}
