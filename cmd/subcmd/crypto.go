package subcmd

import (
	"altools/pkg/crypt"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	cmdKeys := &cobra.Command{
		Use:   "crypt",
		Short: "crypt related functions",
		Long:  "crypt related functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	cmdKeys.AddCommand(cryptoSub()...)

	RootCmd.AddCommand(cmdKeys)
}

func cryptoSub() []*cobra.Command {
	cmdSalt := &cobra.Command{
		Use:   "salt <size>",
		Short: "gen a salt string in hex, default size is 16",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			size := 16
			var err error
			if len(args) > 0 {
				size, err = strconv.Atoi(args[0])
				if err != nil {
					fmt.Printf("size must be a number, error = %v\n", err)
					return nil
				}
			}

			svc := crypt.NewCryptoService()
			salt, err := svc.GenSalt(size)
			if err != nil {
				fmt.Printf("fail to get balance, error = %v\n", err)
				return nil
			}
			fmt.Printf("salt = %s\n", salt)
			return nil
		},
	}

	cmdKey := &cobra.Command{
		Use:   "key <size>",
		Short: "gen a key string in hex, default size is 16",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			size := 16
			var err error
			if len(args) > 0 {
				size, err = strconv.Atoi(args[0])
				if err != nil {
					fmt.Printf("size must be a number, error = %v\n", err)
					return nil
				}
			}

			svc := crypt.NewCryptoService()
			key, err := svc.GenKey(size)
			if err != nil {
				fmt.Printf("fail to get balance, error = %v\n", err)
				return nil
			}
			fmt.Printf("key = %s\n", key)
			return nil
		},
	}

	cmdAesEnc := &cobra.Command{
		Use:   "aesEnc <data> <key> <ivHex>",
		Short: "encrypt data by aes",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			dataBytes := []byte(args[0])
			keyBytes := []byte(args[1])
			ivBytes, err := hex.DecodeString(args[2])
			if err != nil {
				fmt.Printf("fail to decode iv string, error = %v\n", err)
				return nil
			}

			svc := crypt.NewCryptoService()
			bs, err := svc.AESEncrypt(dataBytes, keyBytes, ivBytes)
			if err != nil {
				fmt.Printf("fail to encrypt data, error = %v\n", err)
				return nil
			}

			fmt.Printf("data    = %s\n", args[0])
			fmt.Printf("encrypt = %s\n", hex.EncodeToString(bs))

			return nil
		},
	}

	cmdAesDec := &cobra.Command{
		Use:   "aesDec <dataHex> <key> <ivHex>",
		Short: "decrypt data by aes",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			dataBytes, err := hex.DecodeString(args[0])
			if err != nil {
				fmt.Printf("fail to decode data string, error = %v\n", err)
				return nil
			}
			keyBytes := []byte(args[1])
			ivBytes, err := hex.DecodeString(args[2])
			if err != nil {
				fmt.Printf("fail to decode iv string, error = %v\n", err)
				return nil
			}

			svc := crypt.NewCryptoService()
			bs, err := svc.AESDecrypt(dataBytes, keyBytes, ivBytes)
			if err != nil {
				fmt.Printf("fail to encrypt data, error = %v\n", err)
				return nil
			}

			fmt.Printf("encrypt = %s\n", args[0])
			fmt.Printf("data    = %s\n", string(bs))
			return nil
		},
	}

	return []*cobra.Command{cmdSalt, cmdKey, cmdAesEnc, cmdAesDec}
}
