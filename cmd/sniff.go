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
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Credential struct {
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	App         string    `json:"app"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	LastUsed    time.Time `json:"lastUsed"`
}

var (
	application string
	username    string
	password    string
	email       string
	description string
)

// sniffCmd represents the sniff command
var sniffCmd = &cobra.Command{
	Use:   "sniff",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sniff called", application, username, password, email, description)
		jsonFile, err := os.Open("creds.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		// read our opened xmlFile as a byte array.
		jsonText, _ := ioutil.ReadAll(jsonFile)

		var credentials []Credential

		err = json.Unmarshal([]byte(jsonText), &credentials)

		newCredential := Credential{
			Email:       "rupeshharode@gmail.com",
			Username:    "techytoes",
			Password:    "abc123",
			App:         "anything",
			Description: description,
			Created:     time.Now(),
			LastUsed:    time.Now(),
		}

		credentials = append(credentials, newCredential)
		// now Marshal it
		result, _ := json.Marshal(credentials)

		f, err := os.OpenFile("creds.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Println(err)
		}

		n, err := io.WriteString(f, string(result))
		if err != nil {
			fmt.Println(n, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sniffCmd)
	sniffCmd.Flags().StringVarP(&application, "app", "", "", "Help message for toggle")
	sniffCmd.Flags().StringVarP(&username, "username", "", "", "Help message for toggle")
	sniffCmd.Flags().StringVarP(&password, "password", "", "", "Help message for toggle")
	sniffCmd.Flags().StringVarP(&email, "email", "", "", "Help message for toggle")
	sniffCmd.Flags().StringVarP(&description, "desc", "", "", "Help message for toggle")
}
