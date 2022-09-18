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

	"github.com/spf13/cobra"
)

// createProfileCmd represents the createProfile command
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
		res, err := http.Post("http://localhost:8080/profiles", "application/json", bytes.NewBuffer(profileNameJson))
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createProfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createProfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
