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
		var profile types.Profile
		profileName := types.ProfileName{Name: args[0]}
		profileNameJson, err := json.Marshal(profileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		endpoint := fmt.Sprintf("%s/profiles", config.Host)
		res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(profileNameJson))
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
				fmt.Printf("Successfully created a profile %v\n", profile.Name)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createProfileCmd)
}
