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
		recordBuilder := types.RecordBuilder{Host: args[0], Value: args[1], ProfileName: args[2]}
		record, err := createRecord(recordBuilder)
		if err != nil {
			return err
		}
		fmt.Printf("Successfully created a record Host: %s, Value: %s\n", record.Host, record.Value)
		return nil
	},
}

func createRecord(recordBuilder types.RecordBuilder) (record types.Record, err error) {
	recordBuilderJson, err := json.Marshal(recordBuilder)
	if err != nil {
		return
	}
	endpoint := fmt.Sprintf("%s/records", config.Host)
	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(recordBuilderJson))
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Something went wrong while requesting to %s.\nStatus: %d\n", endpoint, res.StatusCode)
		return
	} else {
		body, _ := io.ReadAll(res.Body)
		if err = json.Unmarshal(body, &record); err != nil {
			return
		}
		return record, nil
	}
}

func init() {
	rootCmd.AddCommand(createRecordCmd)
}
