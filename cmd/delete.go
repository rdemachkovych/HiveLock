package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [key]",
	Short: "Delete a secret",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: hivelock delete [key]")
			return
		}
		key := args[0]
		req, err := http.NewRequest("DELETE", apiURL+"/secrets/"+key, nil)
		if err != nil {
			log.Fatal(err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Println("Secret deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
