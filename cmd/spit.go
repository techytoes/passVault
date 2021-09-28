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
	"passVault/helpers"
	"passVault/models"

	"github.com/spf13/cobra"
)


// spitCmd represents the spit command
var spitCmd = &cobra.Command{
	Use:   "spit",
	Short: "A brief description of your command",
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

		for i:=0; i<len(credentials); i++ {
			if credentials[i].App == AppName {
				fmt.Println(credentials[i].App, credentials[i].Email, credentials[i].Username, credentials[i].Password)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(spitCmd)
	sniffCmd.Flags().StringVarP(&AppName, "app", "", "", "application/website for the credential")
	sniffCmd.Flags().StringVarP(&UserName, "username", "", "", "username for the app")
}
