package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hi-gin",
		Short: "hi-gin 代表编译后的二进制文件",
	}
	rootFlag *pflag.FlagSet
	cfgFile  string
)

func init() {
	rootFlag = rootCmd.PersistentFlags()
	rootFlag.StringVarP(&cfgFile, "config", "c", "", "配置文件路径")
}

func Execute() {
	rootCmd.Execute()
}
