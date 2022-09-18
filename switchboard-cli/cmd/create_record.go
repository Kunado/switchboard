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
}
