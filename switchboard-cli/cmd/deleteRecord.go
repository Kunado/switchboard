/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"switchboard-server/types"
	"time"

	"github.com/spf13/cobra"
)

// deleteRecordCmd represents the deleteRecord command
var deleteRecordCmd = &cobra.Command{
	Use:   "deleteRecord",
	Short: "delete a record",
	Long:  `Delete a record with its value.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var records []types.Record
		recordValue := types.RecordValue{Value: args[0]}
		recordValueJson, err := json.Marshal(recordValue)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/records", bytes.NewBuffer(recordValueJson))
		if err != nil {
			return err
		}
		client := &http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
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
				fmt.Printf("Successfully deleted a record: %v\n", recordValue.Value)
				for _, record := range records {
					fmt.Printf("Host: %v, Value: %v\n", record.Host, record.Value)
				}
			}
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(deleteRecordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteRecordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteRecordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
