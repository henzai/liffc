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

package liff

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init liffc. generate dotenv file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Bad argumentes. i.e. >liff init $LINE_ACCESS_TOKEN")
		}

		lineAccessToken := args[0]
		file, err := os.Create(`.env`)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		data := []byte(fmt.Sprintf("LINE_ACCESS_TOKEN=%v", lineAccessToken))
		_, err = file.Write(data)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
}

func NewInitCommand() *cobra.Command {
	return initCmd
}
