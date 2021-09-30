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
	"time"

	"github.com/spf13/cobra"
	"passVault/helpers"
	"passVault/models"
)

var (
	password    string
	description string
)

// sniffCmd represents the sniff command
var sniffCmd = &cobra.Command{
	Use:   "sniff",
	Short: "Fetch credential info related to an application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		// Get flag values
		app, _ := cmd.Flags().GetString("app")
		description, _ := cmd.Flags().GetString("desc")
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		// Opening JSON file
		jsonText := helpers.ReadJson("creds.json")

		// Unmarshalling existing content of the JSON file
		var credentials []models.Credential
		if err := json.Unmarshal([]byte(jsonText), &credentials); err != nil {
			panic(err)
		}
		if len(EncKey) != 32 {
			EncKey = "the-key-has-to-be-32-bytes-long!"
			EncKey = "chotu-shanti-suraj-28-1315-betul"
		}
		// Create encrypted password
		encPassword, err := helpers.Encrypt([]byte(password), []byte(EncKey))
		if err != nil {
			panic(err)
		}

		// Create new credential object
		newCredential := models.Credential{
			Email:       email,
			Username:    username,
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
		helpers.OverwriteJson("creds.json", result)
	},
}

func init() {
	rootCmd.AddCommand(sniffCmd)
	sniffCmd.Flags().String("password", "", "password for the app")
	sniffCmd.Flags().String("desc", "", "any description for the credential info")
	sniffCmd.Flags().String("app", "", "application/website for the credential")
	sniffCmd.Flags().String("email","", "email for the app")
	sniffCmd.Flags().String("username", "", "username for the app")
}
