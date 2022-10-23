package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(InitConfig)
}

var rootCmd = &cobra.Command{
	Use:   "pop [project]",
	Args:  cobra.ExactArgs(1),
	Short: "Find & open project in directories",
	Run: func(cmd *cobra.Command, args []string) {
		directories := []string{
			"/Users/nico/projects/coveo",
			"/Users/nico/projects/qubit",
			"/Users/nico/projects/nico",
		}

		project := args[0]

		projectInDirectory, err := findProjectInDirectories(project, directories)

		if err != nil {
			fmt.Print(err)
			os.Exit(0)
		}

		openProject := exec.Command(viper.GetString("editor"), projectInDirectory)

		if openProject.Run() != nil {
			log.Fatal(err)
		}
	},
}

func findProjectInDirectories(project string, directories []string) (string, error) {
	for _, directory := range directories {
		projectInDirectory := directory + "/" + project
		if _, err := os.Stat(projectInDirectory); !os.IsNotExist(err) {
			return projectInDirectory, nil
		}
	}

	formattedDirectories := ""
	for _, directory := range directories {
		formattedDirectories += "- " + directory + "\n"
	}

	return "", errors.New("Project '" + project + "' not found after looking in:\n" + formattedDirectories)
}
