package subcmd

import (
	"altools/pkg/wallet"
	"fmt"
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

	cmdKeys.AddCommand(keysSub()...)

	RootCmd.AddCommand(cmdKeys)
}

const (
	defaultPath = "m/44'/60'/0'/0/0" //60
)

func keysSub() []*cobra.Command {
	cmdGen := &cobra.Command{
		Use:   "gen <path>",
		Short: "gen a new wallet by path",
		//Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := defaultPath
			if len(args) > 0 {
				path = args[0]
			}

			svc := wallet.NewWalletService()
			acc, err := svc.GenAccount(path)
			if err != nil {
				fmt.Printf("Fail to gen a new wallet account, error=%v\n", err)
				return nil
			}

			fmt.Printf("Mnemonic   = %s\n", acc.Mnemonic)
			fmt.Printf("Address    = %s\n", acc.Address)
			fmt.Printf("PrivateKey = %s\n", acc.PrivateKey)

			return nil
		},
	}

	return []*cobra.Command{cmdGen}
}
