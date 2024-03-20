package wallet

import (
	"altools/plugins/blockchain/hdwallet"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"io"
	"math/big"
	"testing"
)

func TestKey_bip39(t *testing.T) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	fmt.Println("entropy: ", hexutil.Encode(entropy), "len =", len(entropy)) // 32 byte
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")
	fmt.Println("seed: ", hexutil.Encode(seed), "len =", len(seed)) // 64 byte

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)
}

func TestKey_nostr(t *testing.T) {
	params := btcec.S256().Params()
	fmt.Println("params.BitSize = ", params.BitSize) // 256bit  32byte

	seed := make([]byte, params.BitSize/8+8)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("seed: ", hexutil.Encode(seed), "len =", len(seed)) // 40 byte

	one := new(big.Int).SetInt64(1)
	k := new(big.Int).SetBytes(seed)
	n := new(big.Int).Sub(params.N, one)
	k.Mod(k, n)
	k.Add(k, one)

	privateKey := hexutil.Encode(k.Bytes())
	fmt.Println("privateKey: ", privateKey, "len =", len(privateKey))
}

func TestKey_hdwallet(t *testing.T) {
	//target := "npub1v88gfe8f36hv07t67n6nvf0rjmuvcup8d47shk4z6urk3urlmvxqqk55w8"

	mnemonic := "found garlic dwarf parent hard true battle input payment right clump priority"
	entropy, err := bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("entropy: ", entropy, "len =", len(entropy)) // 12 -> 16 byte

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "123456")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("seed: ", hexutil.Encode(seed), "len =", len(seed)) //32 byte

	hash := sha1.Sum(seed)
	fmt.Println("hash: ", hexutil.Encode(hash[:20]), "len =", len(hash)) //20 byte

	//生成nostr privateKey
	newSeed := seed[:40]
	fmt.Println("newSeed: ", hexutil.Encode(newSeed), "len =", len(newSeed))

	params := btcec.S256().Params()
	one := new(big.Int).SetInt64(1)
	k := new(big.Int).SetBytes(seed)
	n := new(big.Int).Sub(params.N, one)
	k.Mod(k, n)
	k.Add(k, one)

	privateKey := hexutil.Encode(k.Bytes())
	fmt.Println("privateKey: ", privateKey, "len =", len(privateKey))
}

func TestKey_NoStr(t *testing.T) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	fmt.Println("entropy: ", hexutil.Encode(entropy), "len =", len(entropy)) // 32 byte
	mnemonic, _ := bip39.NewMnemonic(entropy)

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		t.Fatal(err)
	}

	path := fmt.Sprintf("m/44'/%d'/0'/0/%d", 1237, 0)

	dpath := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(dpath, false)
	if err != nil {
		t.Fatal(err)
	}

	address := account.Address.Hex()
	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("address: ", address, "len =", len(address))
	fmt.Println("privateKey: ", privateKey, "len =", len(privateKey))

	defoPrivKey, err := getEncodePrivateKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("defoPrivKey: ", defoPrivKey, "len =", len(defoPrivKey))
}

func getEncodePrivateKey(hexPrivateKey string) (string, error) {
	b, err := hex.DecodeString(hexPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key hex: %w", err)
	}
	bits5, err := bech32.ConvertBits(b, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode("dsec", bits5)
}
