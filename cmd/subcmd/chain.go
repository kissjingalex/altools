package subcmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cmdKeys := &cobra.Command{
		Use:   "chain",
		Short: "chain related functions",
		Long:  "chain related functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(chainSub()...)

	RootCmd.AddCommand(cmdKeys)
}

func chainSub() []*cobra.Command {
	cmdBalance := &cobra.Command{
		Use:   "balance <chain> <address>",
		Short: "fetch balance of address on chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return []*cobra.Command{cmdBalance}
}
