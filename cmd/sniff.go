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
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"passVault/models"
	"passVault/util"
)

type promptContent struct {
	errorMsg string
	label    string
}

// sniffCmd represents the sniff command
var sniffCmd = &cobra.Command{
	Use:   "sniff",
	Short: "Save credential information for a website/app",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		config, err := util.LoadConfig(".")
		if err != nil {
			panic(err)
		}
		// Get flag values
		app, _ := cmd.Flags().GetString("app")
		description, _ := cmd.Flags().GetString("desc")
		username, _ := cmd.Flags().GetString("username")
		password := getPassword()

		// Opening JSON file
		jsonText := util.ReadJson("creds.json")

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
		util.OverwriteJson("creds.json", result)
	},
}

func getPassword() string {
	wordPromptContent := promptContent{
		"Please provide a password.",
		"What is the password for this application?",
	}
	password := promptGetInput(wordPromptContent)
	return password
}

func init() {
	rootCmd.AddCommand(sniffCmd)
	sniffCmd.Flags().String("desc", "", "any description for the credential info")
	sniffCmd.Flags().String("app", "", "application/website for the credential")
	sniffCmd.Flags().String("username", "", "username for the app")
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}
