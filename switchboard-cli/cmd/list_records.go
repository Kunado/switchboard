package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"switchboard-server/types"

	"github.com/spf13/cobra"
)

var recordsCmd = &cobra.Command{
	Use:   "listRecords",
	Short: "list records",
	Long:  `List records now server processable`,
	RunE: func(cmd *cobra.Command, args []string) error {
		records, err := getRecords()
		if err != nil {
			return err
		}
		for _, record := range records {
			fmt.Printf("Host: %v, Value: %v\n", record.Host, record.Value)
		}
		return nil
	},
}

func getRecords() (records []types.Record, err error) {
	endpoint := fmt.Sprintf("%s/records", config.Host)
	res, err := http.Get(endpoint)
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
	rootCmd.AddCommand(recordsCmd)
}
