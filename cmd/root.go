package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var apiURL = "http://localhost:8080"

var rootCmd = &cobra.Command{
	Use:   "hivelock",
	Short: "HiveLock CLI Tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HiveLock CLI - Use 'hivelock server' to start the API server")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
