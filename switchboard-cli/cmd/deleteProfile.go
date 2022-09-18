/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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

// deleteProfileCmd represents the deleteProfile command
var deleteProfileCmd = &cobra.Command{
	Use:   "deleteProfile",
	Short: "delete a profile",
	Long:  `Delete a profile`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var profiles []types.Profile
		profileName := types.ProfileName{Name: args[0]}
		profileNameJson, err := json.Marshal(profileName)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/profiles", bytes.NewBuffer(profileNameJson))
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
			if err := json.Unmarshal(body, &profiles); err != nil {
				return err
			} else {
				fmt.Printf("Successfully deleted a profile: %v\n", profileName.Name)
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
	rootCmd.AddCommand(deleteProfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteProfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteProfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
