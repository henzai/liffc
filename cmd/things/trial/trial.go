package trial

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	NO_LINE_ACCESS_TOKEN = "Error: set environment variable following command. \n$ liffctl init {LINE_ACCESS_TOKEN}"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add trial product",
	Long:  `bbb`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		liffID := args[0]
		fmt.Println(liffID)
		productName, _ := cmd.PersistentFlags().GetString("name")
		fmt.Println(productName)
		lineAccessToken := viper.GetString("line_access_token")
		if lineAccessToken == "" {
			cmd.Println(NO_LINE_ACCESS_TOKEN)
			os.Exit(1)
		}
	},
}

func init() {
	addCmd.PersistentFlags().StringP("name", "u", "", "product name")
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list trial product",
	Long:    `bbb`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		lineAccessToken := viper.GetString("line_access_token")
		if lineAccessToken == "" {
			cmd.Println(NO_LINE_ACCESS_TOKEN)
			os.Exit(1)
		}
	},
}

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "remove trial product",
	Long:    `bbb`,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lineAccessToken := viper.GetString("line_access_token")
		if lineAccessToken == "" {
			cmd.Println(NO_LINE_ACCESS_TOKEN)
			os.Exit(1)
		}
		productID := args[0]
		fmt.Println(productID)
	},
}

func NewSubCommand() []*cobra.Command {
	return []*cobra.Command{removeCmd, addCmd, listCmd}
}
