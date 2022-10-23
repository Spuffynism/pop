package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "Add a directory",
	Run: func(cmd *cobra.Command, args []string) {
		newDirectory := args[0]
		newDirectories := append(GetDirectories(), newDirectory)

		viper.Set(directoriesConfigKey, newDirectories)
		cobra.CheckErr(viper.WriteConfig())

		fmt.Printf("Set %s to:\n%s\n", directoriesConfigKey, formatDirectories(GetDirectories()))
	},
}

func init() {
	directoryCmd.AddCommand(addCmd)
}
