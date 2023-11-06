package subcmd

import (
	"altools/pkg/interact"
	"github.com/spf13/cobra"
)

func init() {
	cmdKeys := &cobra.Command{
		Use:   "interact",
		Short: "interact related functions",
		Long:  "interact related functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(interactSub()...)

	RootCmd.AddCommand(cmdKeys)
}

func interactSub() []*cobra.Command {
	cmdStart := &cobra.Command{
		Use:   "start",
		Short: "start to interact",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			svc := interact.NewInteractService()
			svc.Start()

			return nil
		},
	}

	return []*cobra.Command{cmdStart}
}
