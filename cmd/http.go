package cmd

import "github.com/spf13/cobra"

var (
	httpCmd = &cobra.Command{
		Use:   "http",
		Short: "启动http服务",
		Long: `启动方式:
   ./hiGin http --config=./config/testing.toml 
   ./hiGin http -c=./config/testing.toml
	注: --config/-c:代表配置文件`,
		DisableSuggestions: true,
		Run:                httpServer,
	}
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

func httpServer(cmd *cobra.Command, args []string) {
	runServer()
}
