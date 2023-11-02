package subcmd

import (
	"altools/pkg/chain"
	"altools/plugins/blockchain/util"
	"fmt"
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
	cmdCheckAddress := &cobra.Command{
		Use:   "checkAddress <address>",
		Short: "check if address is valid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			address := args[0]

			svc := chain.NewChainService()
			addr, err := svc.CheckAddress(address)
			if err != nil {
				fmt.Printf("invalid address, error = %v\n", err)
				return nil
			}

			fmt.Printf("address = %v\n", addr)

			return nil
		},
	}

	cmdPersonalSign := &cobra.Command{
		Use:   "personalSign <content> <privateKey>",
		Short: "generate personal_sign text",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			content := args[0]
			privateKeyStr := args[1]

			svc := chain.NewChainService()
			text, err := svc.GenSignatureForText(content, privateKeyStr)
			if err != nil {
				fmt.Printf("fail to generate signature, error = %v\n", err)
				return nil
			}
			fmt.Printf("signature = %s\n", text)

			return nil
		},
	}

	cmdValidatePersonalSign := &cobra.Command{
		Use:   "validatePersonalSign <content> <address> <signHex>",
		Short: "validate a personal_sign text",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			content := args[0]
			address := args[1]
			signHex := args[2]

			svc := chain.NewChainService()
			isValid := svc.ValidateSignatureForEthereum(content, address, signHex)
			fmt.Printf("isValid = %v\n", isValid)

			return nil
		},
	}

	cmdBalance := &cobra.Command{
		Use:   "balance <address> <chain>",
		Short: "fetch balance of address on chain. if chain is empty, default is Ethereum",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			address := args[0]
			chainName := chain.DefaultChain
			if len(args) > 1 {
				chainName = args[1]
			}

			svc := chain.NewChainService()
			balance, err := svc.GetBalance(address, chainName)
			if err != nil {
				fmt.Printf("fail to get balance, error = %v\n", err)
				return nil
			}
			fmt.Printf("balance = %s\n", balance.String())
			fmt.Printf("balance(ether) = %s\n", util.ToEtherDecimal(balance).String())
			return nil
		},
	}

	cmdNonce := &cobra.Command{
		Use:   "nonce <address> <chain>",
		Short: "fetch pending nonce of address on chain. if chain is empty, default is Ethereum",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			address := args[0]
			chainName := chain.DefaultChain
			if len(args) > 1 {
				chainName = args[1]
			}

			svc := chain.NewChainService()
			nonce, err := svc.GetPendingNonce(address, chainName)
			if err != nil {
				fmt.Printf("fail to get nonce, error = %v\n", err)
				return nil
			}
			fmt.Printf("nonce = %d\n", nonce)

			return nil
		},
	}

	cmdGasPrice := &cobra.Command{
		Use:   "gasPrice <chain>",
		Short: "fetch suggest gas price of chain. if chain is empty, default is Ethereum",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			chainName := chain.DefaultChain
			if len(args) > 0 {
				chainName = args[0]
			}

			svc := chain.NewChainService()
			price, err := svc.GetSuggestGasPrice(chainName)
			if err != nil {
				fmt.Printf("fail to get gas price, error = %v\n", err)
				return nil
			}
			fmt.Printf("SuggestGasPrice = %s\n", price.String())
			fmt.Printf("SuggestGasPrice(ether) = %s\n", util.ToEtherDecimal(price).String())
			return nil
		},
	}

	return []*cobra.Command{cmdBalance, cmdNonce, cmdGasPrice, cmdCheckAddress, cmdPersonalSign, cmdValidatePersonalSign}
}
