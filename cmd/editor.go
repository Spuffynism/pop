package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var editorConfigKey = "editor"

var editorCmd = &cobra.Command{
	Use:   "editor [editor]",
	Args:  cobra.ExactArgs(1),
	Short: "Editor used to pop open projects",
	Run: func(cmd *cobra.Command, args []string) {
		editor := args[0]

		viper.Set(editorConfigKey, editor)
		cobra.CheckErr(viper.WriteConfig())

		fmt.Printf("Set %s to '%s'\n", editorConfigKey, editor)
	},
}

func GetEditor() string {
	return viper.GetString(editorConfigKey)
}

func init() {
	configCmd.AddCommand(editorCmd)
}
