package chain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func signHash(data []byte) common.Hash {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	hash := crypto.Keccak256Hash([]byte(msg))
	return hash
}

// ValidateSignatureForEthereum 验证personal_sign的签名 https://goethereumbook.org/signature-verify/
func (svc *ChainService) ValidateSignatureForEthereum(data string, address string, signHex string) bool {
	sign, err := hexutil.Decode(signHex)
	if err != nil {
		return false
	}

	dataHash := signHash([]byte(data))
	//fmt.Printf("sign[64]=%d\n", sign[64])
	if sign[64] >= 27 {
		sign[64] -= 27
	}

	sigPublicKey, err := crypto.SigToPub(dataHash.Bytes(), sign)
	if err != nil {
		return false
	}

	signAddress := crypto.PubkeyToAddress(*sigPublicKey)

	return strings.ToLower(address) == strings.ToLower(signAddress.String())
}

// GenSignatureForText person_sign
func (svc *ChainService) GenSignatureForText(data string, privateKeyStr string) (string, error) {
	if strings.HasPrefix(privateKeyStr, "0x") {
		privateKeyStr = privateKeyStr[2:]
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", err
	}

	hash := signHash([]byte(data))
	fmt.Printf("hash=%s\n", hexutil.Encode(hash.Bytes()))

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(signature), nil
}
