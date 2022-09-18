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
}
