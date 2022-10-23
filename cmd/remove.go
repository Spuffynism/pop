package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Args:  cobra.ExactArgs(1),
	Short: "Remove a directory",
	Run: func(cmd *cobra.Command, args []string) {
		directoryToRemove := args[0]
		newDirectories := []string{}
		for _, directory := range GetDirectories() {
			if directory == directoryToRemove {
				continue
			}

			newDirectories = append(newDirectories, directory)
		}

		viper.Set(directoriesConfigKey, newDirectories)

		if err := viper.WriteConfig(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Set %s to:\n%s\n", directoriesConfigKey, formatDirectories(GetDirectories()))
		}
	},
}

func init() {
	directoryCmd.AddCommand(removeCmd)
}
