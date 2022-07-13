/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/techytoes/passVault/models"
	"github.com/techytoes/passVault/util"
)

// swallowCmd represents the swallow command
var swallowCmd = &cobra.Command{
	Use:   "swallow",
	Short: "Delete credential information for a website/app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		// Get flag values
		app, _ := cmd.Flags().GetString("app")

		// Opening JSON file
		jsonText := util.ReadJson(fmt.Sprintf("%s/creds.json", dirname))

		// Unmarshalling existing content of the JSON file
		var credentials []models.Credential
		if err := json.Unmarshal([]byte(jsonText), &credentials); err != nil {
			panic(err)
		}

		var newCredentials []models.Credential
		for i := 0; i < len(credentials); i++ {
			if credentials[i].App != app {
				newCredentials = append(newCredentials, credentials[i])
			}
		}

		// now Marshal it
		result, _ := json.Marshal(newCredentials)

		// Overwrite the JSON file with the new data.
		util.OverwriteJson(fmt.Sprintf("%s/creds.json", dirname), result)
	},
}

func init() {
	rootCmd.AddCommand(swallowCmd)
	swallowCmd.Flags().String("app", "", "Name for the app/website")
}
