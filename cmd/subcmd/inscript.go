package subcmd

import (
	"altools/pkg/inscript"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"os"
)

const (
	ENV_WALLET_ADDRESS     = "WALLET_ADDRESS"
	ENV_WALLET_PRIVATE_KEY = "WALLET_PRIVATE_KEY"
)

func init() {
	cmdKeys := &cobra.Command{
		Use:   "inscript",
		Short: "inscript related functions",
		Long:  "inscript related functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(inscriptSub()...)

	RootCmd.AddCommand(cmdKeys)
}

func inscriptSub() []*cobra.Command {

	cmdMint := &cobra.Command{
		Use:   "mint <chainType> <toAddress> <contentInHex>",
		Short: "mint an inscription for text type. available chain type: Goerli、Base_Goerli、Ethereum ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			privateKey := os.Getenv(ENV_WALLET_PRIVATE_KEY)
			if len(privateKey) == 0 {
				fmt.Println("please set env WALLET_PRIVATE_KEY")
				return nil
			}
			fmt.Printf("privateKey=%s\n", privateKey)

			chainType := args[0]
			toAddress := args[1]
			data := args[2]

			svc := inscript.NewInscriptService()
			rs, err := svc.Mint(chainType, privateKey, toAddress, data, false)
			if err != nil {
				fmt.Printf("Fail to mint, error = %v\n", err)
				return nil
			}

			fmt.Println("Mint successfully! result is :")
			spew.Dump(rs)

			return nil
		},
	}

	cmdMintMock := &cobra.Command{
		Use:   "mintMock <chainType> <toAddress> <contentInHex>",
		Short: "mock to mint an inscription for text type. available chain type: Goerli、Base_Goerli、Ethereum ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			privateKey := os.Getenv(ENV_WALLET_PRIVATE_KEY)
			if len(privateKey) == 0 {
				fmt.Println("please set env WALLET_PRIVATE_KEY")
				return nil
			}
			fmt.Printf("privateKey=%s\n", privateKey)

			chainType := args[0]
			toAddress := args[1]
			data := args[2]

			svc := inscript.NewInscriptService()
			rs, err := svc.Mint(chainType, privateKey, toAddress, data, true)
			if err != nil {
				fmt.Printf("Fail to mint, error = %v\n", err)
				return nil
			}

			fmt.Println("Mint(Mock) successfully! result is :")
			spew.Dump(rs)

			return nil
		},
	}

	return []*cobra.Command{cmdMint, cmdMintMock}
}
