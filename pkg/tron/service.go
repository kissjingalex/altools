package tron

import (
	"altools/pkg/tron/util"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

type TronService struct {
}

var service = &TronService{}

func NewTronService() *TronService {
	return service
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
