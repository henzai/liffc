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
	"os"

	"github.com/henzai/liffc/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send message to userID",
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
		liffID := args[0]
		userID := args[1]
		message := c.LIFF.NewPushMessage(liffID, userID)
		err := c.LIFF.Send(message)
		if err != nil {
			cmd.Println(err)
		}
		return
	},
}

func init() {
	liffCmd.AddCommand(sendCmd)
}
