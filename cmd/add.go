// Copyright © 2018 henzai ry0chord@gmail.com
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
	"fmt"
	"log"
	"os"

	"github.com/henzai/liffc/liff"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	NO_LINE_ACCESS_TOKEN = "Error: set environment variable following command. \n$ liffctl init {LINE_ACCESS_TOKEN}"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add LIFF app",
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
		c := liff.NewClient(lineAccessToken)

		ble, err := cmd.PersistentFlags().GetBool("ble")
		if err != nil {
			log.Fatal(err)
		}

		description, err := cmd.PersistentFlags().GetString("description")
		if err != nil {
			log.Fatal(err)
		}

		liffType, err := cmd.PersistentFlags().GetString("type")
		if err != nil {
			log.Fatal(err)
		}

		appOption, err := liff.NewAppOption(description, liffType, args[0], ble)
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		liffID, err := c.Add(appOption)
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}
		fmt.Println(liffID.LiffID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringP("description", "d", "", "you can descript about its LIFF app")
	addCmd.PersistentFlags().StringP("type", "t", "full", "size of LIFF app. you can select full|tall|compact")
	addCmd.PersistentFlags().BoolP("ble", "b", false, "enable LINE Things")
}
