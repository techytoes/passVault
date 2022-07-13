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

type UserCredential struct {
	username string
	password string
}

// sniffCmd represents the sniff command
var sniffCmd = &cobra.Command{
	Use:   "sniff",
	Short: "Save credential information for a website/app",
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
		userCredential, err := getUsernamePassword()
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
		encPassword, err := util.Encrypt([]byte(userCredential.password), []byte(config.EncryptKey))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create encrypted username
		encUsername, err := util.Encrypt([]byte(userCredential.username), []byte(config.EncryptKey))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create new credential object
		newCredential := models.Credential{
			Username:    encUsername,
			Password:    encPassword,
			App:         app,
			Description: description,
			Created:     time.Now(),
			LastUsed:    time.Now(),
			Version:     1,
		}
		credentials = append(credentials, newCredential)
		// now Marshal it
		result, _ := json.Marshal(credentials)

		// Overwrite the JSON file with the new data.
		util.OverwriteJson(fmt.Sprintf("%s/creds.json", dirname), result)
	},
}

func getUsernamePassword() (UserCredential, error) {
	usernamePromptContent := util.PromptContent{
		ErrorMsg: "Please provide a valid username.",
		Label:    "What is the username for this application?",
	}
	username := util.PromptGetInput(usernamePromptContent, false)

	passwordPromptContent := util.PromptContent{
		ErrorMsg: "Please provide a valid password.",
		Label:    "Please provide a valid password.",
	}
	password := util.PromptGetInput(passwordPromptContent, true)

	passwordPromptAgainContent := util.PromptContent{
		ErrorMsg: "Please provide password again",
		Label:    "Please provide password again",
	}
	passwordAgain := util.PromptGetInput(passwordPromptAgainContent, true)
	if password != passwordAgain {
		return UserCredential{}, fmt.Errorf("passwords don't match!! please try again")
	}
	return UserCredential{username, password}, nil
}

func init() {
	rootCmd.AddCommand(sniffCmd)
	sniffCmd.Flags().String("app", "", "Name for the app/website")
	sniffCmd.Flags().String("desc", "", "any description for the credential info")
}
