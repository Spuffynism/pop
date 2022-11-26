package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ioutil.ReadFile(viper.ConfigFileUsed())
		cobra.CheckErr(err)

		fmt.Print(string(config))
	},
}

func init() {
	configCmd.AddCommand(showConfigCmd)
}
