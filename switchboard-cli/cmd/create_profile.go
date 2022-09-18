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

var createProfileCmd = &cobra.Command{
	Use:   "createProfile",
	Short: "create new profile",
	Long:  `Create new profile with its name`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName := types.ProfileName{Name: args[0]}
		profile, err := createProfile(profileName)
		if err != nil {
			return err
		}
		fmt.Printf("Successfully created a profile %v\n", profile.Name)
		return nil
	},
}

func createProfile(profileName types.ProfileName) (profile types.Profile, err error) {
	profileNameJson, err := json.Marshal(profileName)
	if err != nil {
		return
	}
	endpoint := fmt.Sprintf("%s/profiles", config.Host)
	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(profileNameJson))
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Something went wrong while requesting to %s\n. Status: %d\n", endpoint, res.StatusCode)
		return
	}
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &profile); err != nil {
		return
	}
	return profile, nil
}

func init() {
	rootCmd.AddCommand(createProfileCmd)
}
