package subcmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

func init() {
	cmdKeys := &cobra.Command{
		Use:   "cast",
		Short: "cast related functions",
		Long:  "cast related functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(castSub()...)

	RootCmd.AddCommand(cmdKeys)
}

func castSub() []*cobra.Command {
	cmdToHex := &cobra.Command{
		Use:   "toHex <content>",
		Short: "string to hex",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			content := args[0]

			rs := hexutil.Encode([]byte(content))
			fmt.Printf("hex=%s\n", rs)

			return nil
		},
	}

	cmdFromHex := &cobra.Command{
		Use:   "fromHex <content>",
		Short: "hex to string",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			content := args[0]

			rs, err := hexutil.Decode(content)
			if err != nil {
				fmt.Printf("Fail to cast, error=%v", err)
				return nil
			}

			fmt.Printf("data=%s\n", string(rs))

			return nil
		},
	}

	return []*cobra.Command{cmdToHex, cmdFromHex}
}
