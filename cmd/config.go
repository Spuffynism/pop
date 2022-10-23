package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

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
	viper.SetConfigName(".pop")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
