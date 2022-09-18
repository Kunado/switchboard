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
		profile, err := switchProfile(profileName)
		if err != nil {
			return err
		}
		fmt.Printf("Switched profile to %v.\n", profile.Name)
		return nil
	},
}

func switchProfile(profileName types.ProfileName) (profile types.Profile, err error) {
	profileNameJson, err := json.Marshal(profileName)
	if err != nil {
		return
	}
	endpoint := fmt.Sprintf("%s/switch_profile", config.Host)
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(profileNameJson))
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
	} else {
		body, _ := io.ReadAll(res.Body)
		if err = json.Unmarshal(body, &profile); err != nil {
			return
		}
		return profile, nil
	}
}

func init() {
	rootCmd.AddCommand(switchProfileCmd)
}
