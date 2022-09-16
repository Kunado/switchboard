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
	Run: func(cmd *cobra.Command, args []string) {
		var profile types.Profile
		profileName := types.ProfileName{Name: args[0]}
		profileNameJson, err := json.Marshal(profileName)
		if err != nil {
			fmt.Println(err)
		}
		res, err := http.Post("http://localhost:8080/profiles", "application/json", bytes.NewBuffer(profileNameJson))
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Println("Something went wrong.\n")
		} else {
			body, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(body, &profile); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Printf("Successfully created a profile %v\n", profile.Name)
			}
		}
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
