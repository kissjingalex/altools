package subcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
)

// cmd flag values
var (
	verbose bool
)

var (
	RootCmd = &cobra.Command{
		Use:          "altools",
		Short:        "alex tools",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Long: fmt.Sprintf(`
CLI interface to alex tools

%s`, g("type 'altools --help' for details")),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
)

func init() {
	initConfig()

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", GlobalConfig.Verbose, "verbose")
}

func initConfig() {
	ConfigDir = os.Getenv("HOME") + "/.config/altools"
	if err := os.MkdirAll(ConfigDir, 0700); err != nil {
		panic(err.Error())
	}

	DefaultConfigFile = ConfigDir + "/config.default"
	var err error
	GlobalConfig, err = LoadConfig()
	if err != nil {
		GlobalConfig.Verbose = false
	}
}

// LoadConfig loads config file in yaml format
func LoadConfig() (*Config, error) {
	in, err := os.ReadFile(DefaultConfigFile)
	readConfig := &Config{}
	if err == nil {
		if err := yaml.Unmarshal(in, readConfig); err != nil {
			return readConfig, err
		}
	}
	return readConfig, err
}

// Execute 将所有子命令添加到root命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 调用一次。
func Execute() {
	RootCmd.SilenceErrors = true
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
