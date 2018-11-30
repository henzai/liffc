// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"os"

	"github.com/henzai/liffc/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	NO_LINE_ACCESS_TOKEN = "Error: set environment variable following command. \n$ liffctl init {LINE_ACCESS_TOKEN}"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lineAccessToken := viper.GetString("line_access_token")
		if lineAccessToken == "" {
			cmd.Println(NO_LINE_ACCESS_TOKEN)
			os.Exit(1)
		}
		if len(args) == 0 {
			cmd.Println("Bad argumentes. i.e. >liffctl add URL")
			os.Exit(1)
		}
		c := api.NewClient(lineAccessToken)

		ble, err := cmd.Flags().GetBool("ble")
		if err != nil {
			log.Fatal(err)
		}

		description, err := cmd.Flags().GetString("description")
		if err != nil {
			log.Fatal(err)
		}

		liffType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatal(err)
		}

		addOption := api.NewAddOption(description, liffType, args[0], ble)
		err = c.Add(addOption)
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addCmd.Flags().StringP("description", "d", "", "A help for foo")
	addCmd.Flags().StringP("type", "t", "full", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolP("ble", "b", false, "Help message for toggle")
}
