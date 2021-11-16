/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/techytoes/passVault/models"
	"github.com/techytoes/passVault/util"

	"github.com/spf13/cobra"
)

// sniffCmd represents the sniff command
var sniffCmd = &cobra.Command{
	Use:   "sniff",
	Short: "Save credential information for a website/app",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		config, err := util.LoadConfig(dirname)
		if err != nil {
			panic(err)
		}
		// Get flag values
		app, _ := cmd.Flags().GetString("app")
		description, _ := cmd.Flags().GetString("desc")
		username, password := getUsernamePassword()

		// Opening JSON file
		jsonText := util.ReadJson(fmt.Sprintf("%s/creds.json", dirname))

		// Unmarshalling existing content of the JSON file
		var credentials []models.Credential
		if err := json.Unmarshal([]byte(jsonText), &credentials); err != nil {
			panic(err)
		}

		// Create encrypted password
		encPassword, err := util.Encrypt([]byte(password), []byte(config.EncryptKey))
		if err != nil {
			panic(err)
		}

		// Create encrypted username
		encUsername, err := util.Encrypt([]byte(username), []byte(config.EncryptKey))
		if err != nil {
			panic(err)
		}

		// Create new credential object
		newCredential := models.Credential{
			Username:    encUsername,
			Password:    encPassword,
			App:         app,
			Description: description,
			Created:     time.Now(),
			LastUsed:    time.Now(),
		}
		credentials = append(credentials, newCredential)
		// now Marshal it
		result, _ := json.Marshal(credentials)

		// Overwrite the JSON file with the new data.
		util.OverwriteJson(fmt.Sprintf("%s/creds.json", dirname), result)
	},
}

func getUsernamePassword() (string, string) {
	usernamePromptContent := util.PromptContent{
		ErrorMsg: "Please provide a valid username.",
		Label:    "What is the username for this application?",
	}
	username := util.PromptGetInput(usernamePromptContent)

	passwordPromptContent := util.PromptContent{
		ErrorMsg: "Please provide a valid password.",
		Label:    "What is the password for this application?",
	}
	password := util.PromptGetInput(passwordPromptContent)
	return username, password
}

func init() {
	rootCmd.AddCommand(sniffCmd)
	sniffCmd.Flags().String("app", "", "Name for the app/website")
	sniffCmd.Flags().String("desc", "", "any description for the credential info")
}
