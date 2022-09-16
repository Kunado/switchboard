/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"switchboard-server/types"

	"github.com/spf13/cobra"
)

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "listProfiles",
	Short: "list profiles",
	Long:  `List all profiles registered.`,
	Run: func(cmd *cobra.Command, args []string) {
		var profiles []types.Profile
		res, err := http.Get("http://localhost:8080/profiles")
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Println("something went wrong.\n")
		} else {
			body, _ := io.ReadAll(res.Body)
			if err := json.Unmarshal(body, &profiles); err != nil {
				fmt.Println(err)
				return
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
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
