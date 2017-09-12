package commands

import (
	"encoding/json"
	"fmt"
	"github.com/ethereal-go/ethereal/root/i18n"
	"github.com/ethereal-go/ethereal/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/shiena/ansicolor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"os"
)

var cfgFile string
var pathFileLanguage string
var color io.Writer
var locale i18n.StorageLocale

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
	color = ansicolor.NewAnsiColorWriter(os.Stdout)
	//cobra.OnInitialize(initConfig, initPathFileLanguage)

	cmdLocale.Flags().StringVar(&cfgFile, "config", "", "config file (default is $GOPATH/.app.json)")
	cmdLocale.Flags().StringVar(&pathFileLanguage, "source", "", "path to file with text language")

	RootCmd.AddCommand(cmdLocale)
	RootCmd.AddCommand(cmdAdd)
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

		// Search config in home directory with name ".app" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".app")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintf(color, "%sUsing config file: %s %s\n", "\x1b[32m", cfgFile, "\x1b[0m")
	} else {
		fmt.Fprintf(color, "%sNot found configuration file : %s %s\n", "\x1b[31m", cfgFile, "\x1b[0m")
	}

}

// initConfig reads in file with language
func initPathFileLanguage() {
	if utils.FileExists(pathFileLanguage) {
		file, err := ioutil.ReadFile(pathFileLanguage)
		if err != nil {
			fmt.Fprintf(color, "%s %s %s\n", "\x1b[31m", err, "\x1b[0m")
		}

		err = json.Unmarshal(file, &locale.Structure)
		if err != nil {
			fmt.Fprintf(color, "%s %s %s\n", "\x1b[31m", err, "\x1b[0m")
		}

		fmt.Fprintf(color, "%sUsing file i18n: %s %s\n", "\x1b[32m", pathFileLanguage, "\x1b[0m")
	} else {
		fmt.Fprintf(color, "%sPath to file with language text not found : %s %s\n", "\x1b[31m", pathFileLanguage, "\x1b[0m")
	}

}
