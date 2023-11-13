package crypt

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/miguelmota/go-solidity-sha3"
)

type CryptoService struct {
}

var service = &CryptoService{}

func NewCryptoService() *CryptoService {
	return service
}

func (svc *CryptoService) GenSalt(size int) (s string, err error) {
	switch size {
	default:
		return "", errors.New("crypto/aes: invalid key size " + strconv.Itoa(size))
	case 16, 24, 32:
		break
	}

	challenge := make([]byte, size) // block size
	_, err = rand.Read(challenge)
	return hex.EncodeToString(challenge), err
}

func (svc *CryptoService) GenKey(size int) (s string, err error) {
	switch size {
	default:
		return "", errors.New("crypto/aes: invalid key size " + strconv.Itoa(size))
	case 16, 24, 32:
		break
	}

	challenge := make([]byte, size) // block size
	_, err = rand.Read(challenge)
	return string(challenge), nil
}

// Keccak https://github.com/miguelmota/go-solidity-sha3
// available types :"address", "bytes1", "uint8[]", "bytes32", "uint256", "address[]", "uint32"
func (svc *CryptoService) Keccak(dataType, content string) (s string, err error) {
	var value interface{}
	if dataType == "uint256" {
		n := new(big.Int)
		n.SetString(content, 10)
		fmt.Printf("n=%s\n", n.String())

		value = solsha3.Uint256(n)
	} else if dataType == "address" {
		value = solsha3.Address(content)
	}

	//types := []string{dataType}
	//values := []interface{}{value}

	hash := solsha3.SoliditySHA3(value)

	return hex.EncodeToString(hash), nil
}

//func (svc *CryptoService) Keccak(content string) (s string, err error) {
//	hash := crypto.Keccak256Hash([]byte(content))
//	return hex.EncodeToString(hash.Bytes()), nil
//}
