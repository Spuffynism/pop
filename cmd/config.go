package cmd

import (
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var configFileName = ".pop"

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure pop",
}

func InitConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetDefault("editor", "idea")

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFileName)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			cobra.CheckErr(err)
		}

		createConfigFile()
	}
}

func createConfigFile() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	cobra.CheckErr(viper.SafeWriteConfigAs(home + "/" + configFileName))
}

func init() {
	rootCmd.AddCommand(configCmd)
}
