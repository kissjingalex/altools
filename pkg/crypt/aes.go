package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// AESEncrypt 加密数据
func (svc *CryptoService) AESEncrypt(data, key, iv []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data len=0")
	}

	// 设置密码
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	content := PKCS5Padding(data, block.BlockSize())
	encrypted := make([]byte, len(content))

	if len(iv) != block.BlockSize() {
		return nil, errors.New("cipher CBC: IV length must equal block size") //一般是16
	}

	aesDecrypt := cipher.NewCBCEncrypter(block, iv)
	aesDecrypt.CryptBlocks(encrypted, content)
	return encrypted, nil
}

// AESDecrypt 解密数据
func (svc *CryptoService) AESDecrypt(src, key, iv []byte) (data []byte, err error) {
	if len(src) == 0 {
		return nil, errors.New("data len=0")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, errors.New("cipher CBC: IV length must equal block size")
	}

	decrypted := make([]byte, len(src))
	aesDecrypt := cipher.NewCBCDecrypter(block, iv)
	aesDecrypt.CryptBlocks(decrypted, src)
	return PKCS5Trimming(decrypted), nil
}

// PKCS5Padding PKCS5包装
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS5Trimming 解包装
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
