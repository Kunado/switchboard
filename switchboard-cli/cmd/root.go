package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "switchboard",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

type Config struct {
	Host string
}

var config Config

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&config.Host, "host", "http://localhost", "host to send requests to.\nYou can set host with host flag like --host=http://host.com,\nor definition of HOST=http://host.com in ~/.switchboardrc file.")
	viper.BindPFlag("Host", rootCmd.PersistentFlags().Lookup("host"))

	cobra.OnInitialize(func() {
		homePath := os.Getenv("HOME")
		_, err := os.Stat(fmt.Sprintf("%s/.switchboardrc", homePath))
		if err == nil {
			viper.SetConfigName(".switchboardrc")
			viper.SetConfigType("env")
			viper.AddConfigPath("$HOME")

			if err := viper.ReadInConfig(); err != nil {
				fmt.Printf("An error occurred while loading config file..\n")
				fmt.Println(err)
				os.Exit(1)
			}

			if err := viper.Unmarshal(&config); err != nil {
				fmt.Printf("An error occurred while unmarshaling config file.\n")
				fmt.Println(err)
				os.Exit(1)
			}
		}
	})
}
