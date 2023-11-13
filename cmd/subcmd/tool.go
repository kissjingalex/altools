package subcmd

import "github.com/spf13/cobra"

func init() {
	cmdKeys := &cobra.Command{
		Use:   "tool",
		Short: "show usage of some tools",
		Long:  "show usage of some tools",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(toolSub()...)

	RootCmd.AddCommand(cmdKeys)
}

func toolSub() []*cobra.Command {

	return []*cobra.Command{}
}

const (
	usageCast = `
env:
ETH_RPC_URL=https://xxx

cmd:
cast rpc eth_blockNumber <json request>
cast block <block number, int64>
cast balance <account>

cast tx <txHash>
cast receipt <hash>
cast receipt <hash> —json | jq

cast pretty-calldata 0x000000
cast 4byte <method selector>   (0x + 8位）
cast keccak “<method>”
cast sig “transfer(address,uint256)”

转换
cast --to-hex
cast --to-dec

cast --to-wei
cast --to-uint  如cast --to-uint 100000 ether　将10000转成ether的单位。

cast --to-bytes32
cast --to-ascii

cast --from-wei
cast --format-bytes32-string`
)
