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

	"github.com/spf13/cobra"
)

// createRecordCmd represents the createRecord command
var createRecordCmd = &cobra.Command{
	Use:   "createRecord",
	Short: "create a new cname record",
	Long:  `Create a new cname record`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var record types.Record
		recordBuilder := types.RecordBuilder{Host: args[0], Value: args[1], ProfileName: args[2]}
		recordBuilderJson, err := json.Marshal(recordBuilder)
		if err != nil {
			return err
		}
		endpoint := fmt.Sprintf("%s/records", config.Host)
		res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(recordBuilderJson))
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Something went wrong. Status: %d\n", res.StatusCode)
		} else {
			body, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(body, &record); err != nil {
				return err
			} else {
				fmt.Printf("Successfully created a record Host: %v, Value: %v, ProfileName: %v\n", recordBuilder.Host, recordBuilder.Value, recordBuilder.ProfileName)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createRecordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createRecordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createRecordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
