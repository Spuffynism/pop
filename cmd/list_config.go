package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

var listConfigCmd = &cobra.Command{
	Use:   "list",
	Short: "List the current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ioutil.ReadFile(viper.ConfigFileUsed())
		cobra.CheckErr(err)

		fmt.Print(string(config))
	},
}

func init() {
	configCmd.AddCommand(listConfigCmd)
}
