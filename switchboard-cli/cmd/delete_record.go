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

var deleteRecordCmd = &cobra.Command{
	Use:   "deleteRecord",
	Short: "delete a record",
	Long:  `Delete a record with its value.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		recordValue := types.RecordValue{Value: args[0]}
		records, err := deleteRecord(recordValue)
		if err != nil {
			return err
		}
		fmt.Printf("Successfully deleted a record: %v\n", recordValue.Value)
		for _, record := range records {
			fmt.Printf("Host: %s, Value: %s\n", record.Host, record.Value)
		}
		return err
	},
}

func deleteRecord(recordValue types.RecordValue) (records []types.Record, err error) {
	recordValueJson, err := json.Marshal(recordValue)
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("%s/records", config.Host)
	req, err := http.NewRequest(http.MethodDelete, endpoint, bytes.NewBuffer(recordValueJson))
	if err != nil {
		return
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Something wnet wrong while requesting to %s.\nStatus: %d\n", endpoint, res.StatusCode)
		return
	} else {
		body, _ := io.ReadAll(res.Body)
		if err = json.Unmarshal(body, &records); err != nil {
			return
		}
	}
	return records, nil
}

func init() {
	rootCmd.AddCommand(deleteRecordCmd)
}
