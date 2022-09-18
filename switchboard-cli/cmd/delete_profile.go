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

var deleteProfileCmd = &cobra.Command{
	Use:   "deleteProfile",
	Short: "delete a profile",
	Long:  `Delete a profile`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName := types.ProfileName{Name: args[0]}
		profiles, err := deleteProfile(profileName)
		if err != nil {
			return err
		}
		fmt.Printf("Successfully deleted a profile: %v\n", profileName.Name)
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

func deleteProfile(profileName types.ProfileName) (profiles []types.Profile, err error) {
	profileNameJson, err := json.Marshal(profileName)
	if err != nil {
		return
	}
	endpoint := fmt.Sprintf("%s/profiles", config.Host)
	req, err := http.NewRequest(http.MethodDelete, endpoint, bytes.NewBuffer(profileNameJson))
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
		err = fmt.Errorf("Something went wrong. Status: %d\n", res.StatusCode)
		return
	}
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &profiles); err != nil {
		return
	}
	return profiles, nil
}

func init() {
	rootCmd.AddCommand(deleteProfileCmd)
}
