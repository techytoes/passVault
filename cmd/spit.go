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

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/techytoes/passVault/models"
	"github.com/techytoes/passVault/util"
)

// spitCmd represents the spit command
var spitCmd = &cobra.Command{
	Use:   "spit",
	Short: "Returns info regarding the particular credential",
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
		app, _ := cmd.Flags().GetString("app")
		username, _ := cmd.Flags().GetString("username")

		// Opening JSON files
		jsonText := util.ReadJson(fmt.Sprintf("%s/creds.json", dirname))

		// Unmarshalling existing content of the JSON file
		var credentials []models.Credential
		if err := json.Unmarshal([]byte(jsonText), &credentials); err != nil {
			panic(err)
		}

		if username == "" {
			printAllCredForApp(app, credentials, config)
		} else {
			printCredForAppUsername(app, username, credentials, config)
		}
	},
}

func printCredForAppUsername(app string, username string, credentials []models.Credential, config util.Config) {
	for i := 0; i < len(credentials); i++ {
		decUsername, _ := util.Decrypt(credentials[i].Username, []byte(config.EncryptKey))
		if credentials[i].App == app && string(decUsername[:]) == username {
			decPassword, _ := util.Decrypt(credentials[i].Password, []byte(config.EncryptKey))
			fmt.Println("App-Name:", app)
			fmt.Println("Username:", username)
			clipboard.WriteAll(string(decPassword[:]))
		}
	}
}

func printAllCredForApp(app string, credentials []models.Credential, config util.Config) {
	for i := 0; i < len(credentials); i++ {
		if credentials[i].App == app {
			decUsername, _ := util.Decrypt(credentials[i].Username, []byte(config.EncryptKey))
			fmt.Println("App-Name:", app)
			fmt.Println("Username:", string(decUsername[:]))
		}
	}
}

func init() {
	rootCmd.AddCommand(spitCmd)
	spitCmd.Flags().String("app", "", "application/website for the credential")
	spitCmd.Flags().String("username", "", "username for the app")
}
