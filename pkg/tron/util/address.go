package util

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

func Has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func ToTronAddressBytes(ethAddr string) (address.Address, error) {
	bs, err := hexutil.Decode(ethAddr)
	if err != nil {
		return nil, err
	}
	addressTron := make([]byte, 0)
	addressTron = append(addressTron, address.TronBytePrefix)
	addressTron = append(addressTron, bs...)

	return addressTron, nil
}
