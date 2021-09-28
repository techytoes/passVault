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
		// Opening JSON file
		jsonText := helpers.ReadJson("creds.json")

		// Unmarshalling existing content of the JSON file
		var credentials []models.Credential
		if err := json.Unmarshal([]byte(jsonText), &credentials); err != nil {
			panic(err)
		}

		// Create new credential object
		newCredential := models.Credential{
			Email:       Email,
			Username:    UserName,
			Password:    password,
			App:         AppName,
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
	sniffCmd.Flags().StringVarP(&password, "password", "", "", "password for the app")
	sniffCmd.Flags().StringVarP(&description, "desc", "", "", "any description for the credential info")
}
