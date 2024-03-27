package tron

import (
	cjson "altools/plugins/json"
	nhttp "altools/plugins/net/http"
	"errors"
	"fmt"
	"net/http"
)

const (
	ApiHost      = "apilist.tronscanapi.com"
	ApiHost_Nile = "nileapi.tronscan.org"
	ApiKey       = "99341c07-7df2-48cd-b347-37e756e9f7c2"

	ApiUrl_TransactionInfo = "https://%s/api/transaction-info?hash=%s"
)

type TxInfoResp struct {
	ContractRet  string `json:"contractRet"`
	ContractType int64  `json:"contractType"`
	Confirmed    bool   `json:"confirmed"`
	Cost         TxCost `json:"cost"`
}

type TxCost struct {
	MultiSignFee       int64 `json:"multi_sign_fee"`
	Fee                int64 `json:"fee"`
	MemoFee            int64 `json:"memo_fee"`
	NetFee             int64 `json:"net_fee"`
	NetFeeCost         int64 `json:"net_fee_cost"`
	NetUsage           int64 `json:"net_usage"`
	EnergyPenaltyTotal int64 `json:"energy_penalty_total"`
	EnergyUsage        int64 `json:"energy_usage"`
	EnergyUsageTotal   int64 `json:"energy_usage_total"`
	EnergyFee          int64 `json:"energy_fee"`
	EnergyFeeCost      int64 `json:"energy_fee_cost"`
	OriginEnergyUsage  int64 `json:"origin_energy_usage"`
}

func GetTransactionInfo(txHash string) (*TxInfoResp, error) {
	if len(txHash) == 0 || len(txHash) == 0 {
		return nil, errors.New("invalid parameters")
	}

	url := fmt.Sprintf(ApiUrl_TransactionInfo, ApiHost, txHash)

	header := make(http.Header)
	header.Add("Content-type", "application/json;charset=utf-8")
	header.Add("accept", "application/json")
	header.Add("TRON-PRO-API-KEY", ApiKey)

	bs, err := nhttp.GetRaw(nil, url, header, nil, 20000, 1)
	if err != nil {
		fmt.Printf("fail to request to url=%s\n", url)
		return nil, err
	}
	fmt.Printf("res=%s\n", string(bs))

	data := &TxInfoResp{}
	err = cjson.JSON.Unmarshal(bs, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
