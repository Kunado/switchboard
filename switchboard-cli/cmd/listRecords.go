/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"switchboard-server/types"

	"github.com/spf13/cobra"
)

// recordsCmd represents the records command
var recordsCmd = &cobra.Command{
	Use:   "listRecords",
	Short: "list records",
	Long:  `List records now server processable`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var records []types.Record
		endpoint := fmt.Sprintf("%s/records", config.Host)
		res, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Something went wrong. Status: %d\n", res.StatusCode)
		} else {
			body, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(body, &records); err != nil {
				return err
			} else {
				for _, record := range records {
					fmt.Printf("Host: %v, Value: %v\n", record.Host, record.Value)
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(recordsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
