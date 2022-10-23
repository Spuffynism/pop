package cmd

import (
	"fmt"
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
		project := args[0]

		projectInDirectory, err := FindInDirectories(project, GetDirectories())
		cobra.CheckErr(err)

		openProject := exec.Command(GetEditor(), projectInDirectory)

		if openProject.Run() != nil {
			fmt.Println(err)
		}
	},
}
