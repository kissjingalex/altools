package subcmd

import (
	"altools/pkg/tron"
	"altools/pkg/wallet"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

func init() {
	cmdKeys := &cobra.Command{
		Use:   "wallet",
		Short: "wallet related functions",
		Long:  "wallet related functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(walletSub()...)

	RootCmd.AddCommand(cmdKeys)
}

const (
	defaultPath     = "m/44'/60'/0'/0/0" //60
	defaultTronPath = "m/44'/195'/0'/0/0"
)

func walletSub() []*cobra.Command {
	cmdNewKey := &cobra.Command{
		Use:   "key",
		Short: "gen a new single key",
		RunE: func(cmd *cobra.Command, args []string) error {
			svc := wallet.NewWalletService()
			key, err := svc.GenKey()
			if err != nil {
				fmt.Printf("Fail to gen a single key, error=%v\n", err)
				return nil
			}

			address := hex.EncodeToString(key.Address[:])
			pubKey := key.PrivateKey.PublicKey
			publicKey := hex.EncodeToString(crypto.FromECDSAPub(&pubKey))
			privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

			fmt.Printf("address       = %s\n", address)
			fmt.Printf("publicKey     = %s\n", publicKey)
			fmt.Printf("privateKey    = %s\n", privateKey)

			return nil
		},
	}

	cmdGen := &cobra.Command{
		Use:   "gen <path> <mnemonic>",
		Short: "gen a new wallet by path. use random mnemonic if not given",
		Args:  cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := defaultPath
			if len(args) > 0 {
				path = args[0]
			}

			mnemonic := ""
			if len(args) > 1 {
				mnemonic = args[1]
			}

			svc := wallet.NewWalletService()
			acc, err := svc.GenAccount(path, mnemonic)
			if err != nil {
				fmt.Printf("Fail to gen a new wallet account, error=%v\n", err)
				return nil
			}

			fmt.Printf("Mnemonic   = %s\n", acc.Mnemonic)
			fmt.Printf("Path       = %s\n", acc.Path)
			fmt.Printf("Address    = %s\n", acc.Address)
			fmt.Printf("PrivateKey = %s\n", acc.PrivateKey)

			return nil
		},
	}

	cmdGenTron := &cobra.Command{
		Use:   "genTron <mnemonic>",
		Short: "genTron a new wallet for tron. use random mnemonic if not given",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := defaultTronPath

			mnemonic := ""
			if len(args) > 0 {
				mnemonic = args[1]
			}

			svc := wallet.NewWalletService()
			acc, err := svc.GenAccount(path, mnemonic)
			if err != nil {
				fmt.Printf("Fail to gen a new wallet account, error=%v\n", err)
				return nil
			}

			tronAddress, err := tron.ToTronAddress(acc.Address)
			if err != nil {
				fmt.Printf("Fail gen convert to tron address, error=%v\n", err)
				return nil
			}

			fmt.Printf("Mnemonic    = %s\n", acc.Mnemonic)
			fmt.Printf("Path        = %s\n", acc.Path)
			fmt.Printf("Address     = %s\n", acc.Address)
			fmt.Printf("tronAddress = %s\n", tronAddress)
			fmt.Printf("PrivateKey  = %s\n", acc.PrivateKey)

			return nil
		},
	}

	return []*cobra.Command{cmdGen, cmdNewKey, cmdGenTron}
}
