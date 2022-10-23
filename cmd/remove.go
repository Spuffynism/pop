package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Args:  cobra.ExactArgs(1),
	Short: "Remove a directory",
	Run: func(cmd *cobra.Command, args []string) {
		directoryToRemove := args[0]

		if directoryToRemove == "." {
			fullPathToCurrentDirectory, err := os.Getwd()
			cobra.CheckErr(err)

			directoryToRemove = fullPathToCurrentDirectory
		}

		newDirectories := []string{}
		for _, directory := range GetDirectories() {
			if directory == directoryToRemove {
				continue
			}

			newDirectories = append(newDirectories, directory)
		}

		viper.Set(directoriesConfigKey, newDirectories)
		cobra.CheckErr(viper.WriteConfig())

		fmt.Printf("Set %s to:\n%s\n", directoriesConfigKey, formatDirectories(GetDirectories()))
	},
}

func init() {
	directoryCmd.AddCommand(removeCmd)
}
