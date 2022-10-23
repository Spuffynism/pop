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
	Use:   "pop [project] [branch]",
	Args:  cobra.RangeArgs(1, 2),
	Short: "Find & open project in directories",
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]

		projectInDirectory, err := FindInDirectories(project, GetDirectories())
		cobra.CheckErr(err)

		if len(args) == 2 {
			branch := args[1]

			gitFetch := exec.Command("git", "-C", projectInDirectory, "fetch")
			cobra.CheckErr(gitFetch.Run())

			gitCheckout := exec.Command("git", "-C", projectInDirectory, "checkout", branch)
			cobra.CheckErr(gitCheckout.Run())
		}

		openProject := exec.Command(GetEditor(), projectInDirectory)

		if openProject.Run() != nil {
			fmt.Println(err)
		}
	},
}
