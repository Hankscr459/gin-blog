package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.2"
var rootCmd = &cobra.Command{
	Use:     "Gin_Blog",
	Version: version,
	Short:   "Mongodb Gin Web Server",
	Long: `
		Mongodb Gin Web Server Cli Web Dev
		server: --p YourPORT // default = env.Port
		apiDocs: -d          // default = false
	`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
