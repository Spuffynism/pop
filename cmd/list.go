package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List directories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(formatDirectories(GetDirectories()))
	},
}

func init() {
	directoryCmd.AddCommand(listCmd)
}
