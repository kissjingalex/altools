package wallet

import (
	"altools/plugins/blockchain/hdwallet"
	"github.com/tyler-smith/go-bip39"
)

type WalletService struct {
}

var service = &WalletService{}

func NewWalletService() *WalletService {
	return service
}

func (svc *WalletService) GenAccount(path string) (*Account, error) {
	mnemonic := svc.GenMnemonic()

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	dpath := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(dpath, false)
	if err != nil {
		return nil, err
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return nil, err
	}

	return &Account{
		Mnemonic:   mnemonic,
		Address:    account.Address.Hex(),
		Path:       path,
		PrivateKey: privateKey,
	}, nil
}

func (svc *WalletService) GenMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	return mnemonic
}
