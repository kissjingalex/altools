package crypt

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strconv"
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
