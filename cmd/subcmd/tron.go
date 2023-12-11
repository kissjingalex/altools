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

	return []*cobra.Command{cmdToTronAddr, cmdToEthAddr}
}
