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
	"io/ioutil"
	"os"

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
		fmt.Println("spit called")
		// Open our jsonFile
		jsonFile, err := os.Open("creds.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)

		// we initialize our Users array
		var credentials []Credential

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(byteValue, &credentials)

		// we iterate through every user within our users array and
		// print out the user Type, their name, and their facebook url
		// as just an example
		for i := 0; i < len(credentials); i++ {
			fmt.Println("User Type: " + credentials[i].App)
			fmt.Println("User Age: " + credentials[i].App)
			//fmt.Println("User Name: " + users.Users[i].Name)
			//fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		}
	},
}

func init() {
	rootCmd.AddCommand(spitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
