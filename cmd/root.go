package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "et",
	Short: "A brief description of your application",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName(".et")  // name of config file (without extension)
	viper.SetConfigType("yaml") // use yaml file for variables
	// uses Mitch H.'s homedir to allow for cross compilation
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	viper.Set("host", "https://learn.launchacademy.com")

	viper.AddConfigPath(home) // adding home directory as first search path
	viper.AddConfigPath(".")  // add current directory
	viper.AutomaticEnv()      // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
