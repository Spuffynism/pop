package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

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

		for _, directory := range directories {
			projectInDirectory := directory + "/" + project
			if _, err := os.Stat(projectInDirectory); !os.IsNotExist(err) {
				cmd := exec.Command("idea", projectInDirectory)

				err := cmd.Run()

				if err != nil {
					log.Fatal(err)
				}

				os.Exit(0)
			}
		}

		fmt.Printf("Project '%s' not found after looking in:\n", project)
		for _, directory := range directories {
			fmt.Printf("- %s\n", directory)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
