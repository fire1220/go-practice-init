package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "command",
		Short: "command：代表编译后的二进制文件",
	}
	rootFlags    *pflag.FlagSet
	flagsVarName flagsVarNameType
)

const (
	flagsConfig = "config"
)

type flagsVarNameType struct {
	ConfigFile string
}

func getDefaultConfigFile() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "testing"
	}
	return fmt.Sprintf("./config/%s.toml", appEnv)
}

func init() {
	rootFlags = rootCmd.PersistentFlags()
	rootFlags.StringVarP(&flagsVarName.ConfigFile, flagsConfig, "c", getDefaultConfigFile(), "config file (default is ./conf/testing.toml)")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
