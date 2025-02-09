package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Retrieve a secret",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: hivelock get [key]")
			return
		}
		key := args[0]
		resp, err := http.Get(apiURL + "/secrets/" + key)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Retrieved secret:", string(body))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
