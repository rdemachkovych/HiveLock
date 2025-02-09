package cmd

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Store a secret",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: hivelock set [key] [value]")
			return
		}
		key, value := args[0], args[1]
		body := fmt.Sprintf(`{"key": "%s", "value": "%s"}`, key, value)
		resp, err := http.Post(apiURL+"/secrets", "application/json", bytes.NewBuffer([]byte(body)))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Println("Secret stored successfully.")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
