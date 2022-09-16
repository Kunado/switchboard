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

// switchProfileCmd represents the switchProfile command
var switchProfileCmd = &cobra.Command{
	Use:   "switchProfile",
	Short: "switch enabled profile",
	Long:  `Switch enabled profile.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var profile types.Profile
		profileName := types.ProfileName{Name: args[0]}
		profileNameJson, err := json.Marshal(profileName)
		if err != nil {
			fmt.Println(err)
		}
		req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/switch_profile", bytes.NewBuffer(profileNameJson))
		if err != nil {
			fmt.Println(err)
		}
		client := &http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Println("something went wrong.\n")
		} else {
			body, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(body, &profile); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Printf("Switched profile to %v.\n", profile.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(switchProfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchProfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchProfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
