package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var editorCmd = &cobra.Command{
	Use:   "editor",
	Args:  cobra.ExactArgs(1),
	Short: "Editor used to pop open projects",
	Run: func(cmd *cobra.Command, args []string) {
		editor := args[0]
		viper.Set("editor", editor)

		if err := viper.WriteConfig(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Set editor to '" + editor + "'")
		}
	},
}

func init() {
	configCmd.AddCommand(editorCmd)
}
