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
		profiles, err := getProfiles()
		if err != nil {
			return err
		}
		for _, profile := range profiles {
			if profile.Enabled {
				fmt.Printf("* %v\n", profile.Name)
			} else {
				fmt.Printf("  %v\n", profile.Name)
			}
		}
		return nil
	},
}

func getProfiles() (profiles []types.Profile, err error) {
	endpoint := fmt.Sprintf("%s/profiles", config.Host)
	res, err := http.Get(endpoint)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Something went wrong while requesting to %s.\nStatus: %d\n", endpoint, res.StatusCode)
		return
	} else {
		body, _ := io.ReadAll(res.Body)
		if err = json.Unmarshal(body, &profiles); err != nil {
			return
		}
	}
	return profiles, nil
}

func init() {
	rootCmd.AddCommand(profilesCmd)
}
