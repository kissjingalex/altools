package subcmd

import (
	"altools/pkg/tron"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmdKeys := &cobra.Command{
		Use:   "tron",
		Short: "tron related functions",
		Long:  "tron related functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(tronSub()...)

	RootCmd.AddCommand(cmdKeys)
}

func tronSub() []*cobra.Command {
	cmdToTronAddr := &cobra.Command{
		Use:   "toTronAddr <eth_address>",
		Short: "transform tron address to ethereum address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]

			svc := tron.NewTronService()
			targetAddr, err := svc.ToTronAddress(addr)
			if err != nil {
				fmt.Printf("fail to transform to ethereum address, error = %v\n", err)
				return nil
			}
			fmt.Printf("Tron Address = %s\n", targetAddr)
			return nil
		},
	}

	cmdToEthAddr := &cobra.Command{
		Use:   "toEthAddr <eth_address>",
		Short: "transform ethereum address to tron address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]

			svc := tron.NewTronService()
			targetAddr, err := svc.ToEthAddress(addr)
			if err != nil {
				fmt.Printf("fail to transform to tron address, error = %v\n", err)
				return nil
			}
			fmt.Printf("Ethereum Address = %s\n", targetAddr)
			return nil
		},
	}

	cmdAccount := &cobra.Command{
		Use:   "account <address>",
		Short: "get account info for user tron address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]

			svc := tron.NewTronService()
			acc, err := svc.GetAccountInfo(addr)
			if err != nil {
				fmt.Printf("fail to get account info for tron address, error = %v\n", err)
				return nil
			}
			fmt.Printf("Address   = %s\n", acc.Address)
			fmt.Printf("Type      = %d\n", acc.Type)
			fmt.Printf("Balance   = %s\n", acc.Balance)
			fmt.Printf("Allowance = %s\n", acc.Allowance)
			return nil
		},
	}

	cmdUsdt := &cobra.Command{
		Use:   "usdt <address>",
		Short: "query usdt balance for user tron address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]

			svc := tron.NewTronService()
			rs, err := svc.GetUsdtBalance(addr)
			if err != nil {
				fmt.Printf("fail to get account info for tron address, error = %v\n", err)
				return nil
			}
			fmt.Printf("Address = %s\n", addr)
			fmt.Printf("Value   = %s\n", rs.Value.String())
			fmt.Printf("Amount  = %s\n", rs.Amount.String())
			return nil
		},
	}

	return []*cobra.Command{cmdToTronAddr, cmdToEthAddr, cmdAccount, cmdUsdt}
}
