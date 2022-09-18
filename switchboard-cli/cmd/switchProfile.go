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

var switchProfileCmd = &cobra.Command{
	Use:   "switchProfile",
	Short: "switch enabled profile",
	Long:  `Switch enabled profile.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var profile types.Profile
		profileName := types.ProfileName{Name: args[0]}
		profileNameJson, err := json.Marshal(profileName)
		if err != nil {
			return err
		}
		endpoint := fmt.Sprintf("%s/switch_profile", config.Host)
		req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(profileNameJson))
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
			if err := json.Unmarshal(body, &profile); err != nil {
				return err
			} else {
				fmt.Printf("Switched profile to %v.\n", profile.Name)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(switchProfileCmd)
}
