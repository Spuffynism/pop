package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var directoriesConfigKey = "directories"
var directoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "Where to look for projects",
}

func GetDirectories() []string {
	return viper.GetStringSlice(directoriesConfigKey)
}

func FindInDirectories(file string, directories []string) (string, error) {
	for _, directory := range directories {
		projectInDirectory := directory + "/" + file
		if _, err := os.Stat(projectInDirectory); !os.IsNotExist(err) {
			return projectInDirectory, nil
		}
	}

	return "", errors.New("'" + file + "' not found after looking in:\n" + formatDirectories(directories))
}

func formatDirectories(directories []string) string {
	return "- " + strings.Join(directories, "\n- ")
}

func init() {
	configCmd.AddCommand(directoryCmd)
}
