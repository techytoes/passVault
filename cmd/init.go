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
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/techytoes/passVault/util"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize JSON file to store creds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		// Create file to store creds
		if err = ioutil.WriteFile(fmt.Sprintf("%s/creds.json", dirname), []byte("[]"), 0755); err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}

		encKey := getEncKey()
		config := fmt.Sprintf("{\"enc_key\":\"%s\"}", string(encKey))
		// Create file to store configs
		if err = ioutil.WriteFile(fmt.Sprintf("%s/app.json", dirname), []byte(config), 0755); err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func getEncKey() string {
	encKeyPromptContent := util.PromptContent{
		ErrorMsg: "Please provide a valid Encryption Key",
		Label:    "What is the Encryption Key for this application?",
	}
	encKey := util.PromptGetInput(encKeyPromptContent)
	return encKey
}
