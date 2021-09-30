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
	"passVault/models"
	"passVault/util"

	"github.com/spf13/cobra"
)


// spitCmd represents the spit command
var spitCmd = &cobra.Command{
	Use:   "spit",
	Short: "Returns info regarding the particular credential",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := util.LoadConfig(".")
		if err != nil {
			panic(err)
		}

		app, _ := cmd.Flags().GetString("app")
		//username, _ := cmd.Flags().GetString("username")
		// Opening JSON files
		jsonText := util.ReadJson("creds.json")

		// Unmarshalling existing content of the JSON file
		var credentials []models.Credential
		if err := json.Unmarshal([]byte(jsonText), &credentials); err != nil {
			panic(err)
		}

		for i:=0; i<len(credentials); i++ {
			if credentials[i].App == app {
				fmt.Println(credentials[i].App, credentials[i].Email, credentials[i].Username)
				plaintext, err := util.Decrypt(credentials[i].Password, []byte(config.EncryptKey))
				if err != nil {
					panic(err)
				}
				pass := string(plaintext[:])
				fmt.Println(pass)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(spitCmd)
	spitCmd.Flags().String("app", "", "application/website for the credential")
	spitCmd.Flags().String("username", "", "username for the app")
}
