package commands

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"github.com/shiena/ansicolor"
	"io"
)

var cfgFile string
var color io.Writer

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI your application",
	Long: `
╔═══╗ ╔════╗ ╔╗╔╗ ╔═══╗ ╔═══╗ ╔═══╗ ╔══╗ ╔╗
║╔══╝ ╚═╗╔═╝ ║║║║ ║╔══╝ ║╔═╗║ ║╔══╝ ║╔╗║ ║║
║╚══╗   ║║   ║╚╝║ ║╚══╗ ║╚═╝║ ║╚══╗ ║╚╝║ ║║
║╔══╝   ║║   ║╔╗║ ║╔══╝ ║╔╗╔╝ ║╔══╝ ║╔╗║ ║║
║╚══╗   ║║   ║║║║ ║╚══╗ ║║║║  ║╚══╗ ║║║║ ║╚═╗
╚═══╝   ╚╝   ╚╝╚╝ ╚═══╝ ╚╝╚╝  ╚═══╝ ╚╝╚╝ ╚══╝
	`,
}

func CliExecute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	color  = ansicolor.NewAnsiColorWriter(os.Stdout)
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $GOPATH/.app.json)")


	RootCmd.AddCommand(cmdLocale)
	RootCmd.AddCommand(addLocale)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".seeCobraTest" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".app")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintf(color, "%sUsing config file: " + cfgFile + "%s\n", "\x1b[32m", "\x1b[0m")
	} else{
		fmt.Fprintf(color, "%sNot found configuration file : " + cfgFile + "%s\n", "\x1b[31m", "\x1b[0m")
	}
}
