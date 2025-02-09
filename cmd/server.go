package cmd

import (
	"hivelock/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the API server",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
