package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"switchboard-server/types"

	"github.com/spf13/cobra"
)

var profilesCmd = &cobra.Command{
	Use:   "listProfiles",
	Short: "list profiles",
	Long:  `List all profiles registered.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var profiles []types.Profile
		var endpoint string
		endpoint = fmt.Sprintf("%s/profiles", config.Host)
		res, err := http.Get(endpoint)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Something went wrong. Status: %d\n", res.StatusCode)
		} else {
			body, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(body, &profiles); err != nil {
				return err
			} else {
				for _, profile := range profiles {
					if profile.Enabled {
						fmt.Printf("* %v\n", profile.Name)
					} else {
						fmt.Printf("  %v\n", profile.Name)
					}
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)
}
