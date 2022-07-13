/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/techytoes/passVault/models"
	"github.com/techytoes/passVault/util"
)

// shrugCmd represents the shrug command
var shrugCmd = &cobra.Command{
	Use:   "shrug",
	Short: "Update credential information for a website/app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			return
		}
		config, err := util.LoadConfig(dirname)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Get flag values
		app, _ := cmd.Flags().GetString("app")
		description, _ := cmd.Flags().GetString("desc")
		userCredential, err := util.GetUsernamePassword()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Opening JSON file
		jsonText := util.ReadJson(fmt.Sprintf("%s/creds.json", dirname))

		// Unmarshalling existing content of the JSON file
		var credentials []models.Credential
		if err := json.Unmarshal([]byte(jsonText), &credentials); err != nil {
			fmt.Println(err)
			return
		}

		// Create encrypted password
		encPassword, err := util.Encrypt(userCredential.Password, []byte(config.EncryptKey))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create encrypted username
		encUsername, err := util.Encrypt(userCredential.Username, []byte(config.EncryptKey))
		if err != nil {
			fmt.Println(err)
			return
		}

		//finds the latest version for the app password
		newCredsVersion := 1
		for i := 0; i < len(credentials); i++ {
			if credentials[i].App == app {
				newCredsVersion += 1
			}
		}
		// Create new credential object
		newCredential := models.Credential{
			Username:    encUsername,
			Password:    encPassword,
			App:         app,
			Description: description,
			Created:     time.Now(),
			LastUsed:    time.Now(),
			Version:     int16(newCredsVersion),
		}
		credentials = append(credentials, newCredential)
		// now Marshal it
		result, _ := json.Marshal(credentials)

		// Overwrite the JSON file with the new data.
		util.OverwriteJson(fmt.Sprintf("%s/creds.json", dirname), result)
	},
}

func init() {
	rootCmd.AddCommand(shrugCmd)
	shrugCmd.Flags().String("app", "", "Name for the app/website")
	shrugCmd.Flags().String("desc", "", "any description for the credential info")
}
