package subcmd

import (
	"altools/pkg/ton"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xssnick/tonutils-go/tlb"
)

func init() {
	cmdKeys := &cobra.Command{
		Use:   "ton",
		Short: "ton related functions",
		Long:  "ton related functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(tonSub()...)

	RootCmd.AddCommand(cmdKeys)
}

func tonSub() []*cobra.Command {
	cmdAddress := &cobra.Command{
		Use:   "address <address>",
		Short: "parse ton address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]

			info := ton.ParseAddress(addr, false)
			PrintAsJson(info)
			return nil
		},
	}

	cmdRawAddress := &cobra.Command{
		Use:   "rawAddress <address>",
		Short: "parse ton raw address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]

			info := ton.ParseAddress(addr, true)
			PrintAsJson(info)
			return nil
		},
	}

	cmdWallet := &cobra.Command{
		Use:   "wallet <privateKeyInHex>",
		Short: "create a new wallet if no private key provided, or restore from private key.",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var privateKey string
			if len(args) > 0 {
				privateKey = args[0]
			}

			svc := ton.NewTonService(isTestnet)
			var w *ton.Wallet
			var err error
			if len(privateKey) > 0 {
				w, err = svc.GetWalletFromPrivateKey(privateKey)
			} else {
				w, err = svc.CreateWallet("")
			}
			if err != nil {
				fmt.Printf("Fail to create wallet, error=%v\n", err)
				return nil
			}

			fmt.Printf("Address    = %s\n", w.Address)
			fmt.Printf("PrivateKey = %s\n", w.PrivateKey)
			fmt.Printf("Mnemonic   = %s\n", w.Mnemonic)
			return nil
		},
	}

	cmdWalletM := &cobra.Command{
		Use:   "walletM <Mnemonic>",
		Short: "restore wallet from mnemonic",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			mnemonic := args[0]

			svc := ton.NewTonService(isTestnet)
			w, err := svc.GetWalletFromMnemonic(mnemonic)
			if err != nil {
				fmt.Printf("Fail to restore wallet, error=%v\n", err)
				return nil
			}

			fmt.Printf("Address    = %s\n", w.Address)
			fmt.Printf("PrivateKey = %s\n", w.PrivateKey)
			fmt.Printf("Mnemonic   = %s\n", w.Mnemonic)
			return nil
		},
	}

	cmdBalance := &cobra.Command{
		Use:   "balance <address>",
		Short: "get balance(ton) for an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]

			svc := ton.NewTonService(isTestnet)
			rs, err := svc.GetBalance(addr)
			if err != nil {
				fmt.Printf("Fail to get balance for address(%s), error=%v\n", addr, err)
				return nil
			}
			PrintAsJson(rs)
			return nil
		},
	}

	cmdSend := &cobra.Command{
		Use:   "send <privateKeyInHex> <receiveAddr> <amount>",
		Short: "send some ton of specific amount(float) to target user address",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			privateKey := args[0]
			targetAddr := args[1]
			amount := args[2]

			svc := ton.NewTonService(isTestnet)
			w, err := svc.GetWalletFromPrivateKey(privateKey)
			if err != nil {
				fmt.Printf("Fail to get wallet from private key, error=%v\n", err)
				return nil
			}

			coins := tlb.MustFromTON(amount)

			rs, err := svc.TransferTon(&ton.TxSender{
				Address:    w.Address,
				PrivateKey: w.PrivateKey,
			}, targetAddr, coins.Nano())
			if err != nil {
				fmt.Printf("Fail to transfer ton to user, error=%v\n", err)
				return nil
			}

			PrintAsJson(rs)
			return nil
		},
	}

	return []*cobra.Command{cmdAddress, cmdRawAddress, cmdWallet, cmdWalletM, cmdBalance, cmdSend}
}
