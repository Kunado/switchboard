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
	"switchboard/db"

	"github.com/spf13/cobra"
)

// createRecordCmd represents the createRecord command
var createRecordCmd = &cobra.Command{
	Use:   "createRecord",
	Short: "create a new cname record",
	Long:  `Create a new cname record`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var record db.Record
		recordBuilder := db.RecordBuilder{Host: args[0], Value: args[1], ProfileName: args[2]}
		recordBuilderJson, err := json.Marshal(recordBuilder)
		if err != nil {
			fmt.Println(err)
		}
		res, err := http.Post("http://localhost:8080/records", "application/json", bytes.NewBuffer(recordBuilderJson))
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Println("Something went wrong.\n")
		} else {
			body, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(body, &record); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Printf("Successfully created a record Host: %v, Value: %v, ProfileName: %v\n", recordBuilder.Host, recordBuilder.Value, recordBuilder.ProfileName)
			}
		}
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
