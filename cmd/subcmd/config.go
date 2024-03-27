package subcmd

import "github.com/fatih/color"

var (
	g = color.New(color.FgGreen).SprintFunc()
)

// Directories
var (
	// ConfigDir is the directory to store config file
	ConfigDir string
	// DefaultConfigFile is the default config file name
	DefaultConfigFile string
)

// Error strings
var (
	// ErrConfigNotMatch indicates error for no config matchs
	ErrConfigNotMatch = "no config matchs"
	// ErrEmptyEndpoint indicates error for empty endpoint
	ErrEmptyEndpoint = "no endpoint has been set"
)

// Config defines the config schema
type Config struct {
	Verbose bool `yaml:"verbose"`
	TestNet bool `yaml:"testnet"`
}

// ReadConfig represents the current config read from local
var GlobalConfig *Config
