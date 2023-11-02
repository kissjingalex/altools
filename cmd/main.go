package main

import (
	"altools/cmd/subcmd"
	"altools/pkg/chain"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	version = "1.0"
)

func main() {
	// HACK Force usage of go implementation rather than the C based one. Do the right way, see the
	// notes one line 66,67 of https://golang.org/src/net/net.go that say can make the decision at
	// build time.
	os.Setenv("GODEBUG", "netdns=go") // do not know why

	subcmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show version",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintf(os.Stderr,
				"altools. %v version %s\n",
				path.Base(os.Args[0]), version)
			if err != nil {
				return err
			}
			os.Exit(0)
			return nil
		},
	})

	subcmd.Execute()

	release()
}

// 释放资源
func release() {
	chainSvc := chain.NewChainService()
	chainSvc.Release()
}
