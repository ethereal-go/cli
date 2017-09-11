package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ethereal-go/ethereal/root/i18n/storage/mysql"
	"github.com/spf13/viper"
)

var cmdLocale = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "locale",
	Short: "Localization management",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		arg := args[0]
		switch arg {
		case "fill":
			storage := mysql.LocaleStorageMysql{}.EstablishConnection(map[string]string{
				"login" : viper.GetString("database.login"),
				"password":viper.GetString("database.password"),
				"name":viper.GetString("database.name"),
			})
			storage.Add(locale)
			fmt.Fprintf(color, "%sSuccess fill locale in database! Good job!  %s\n", "\x1b[32m", "\x1b[0m")
		default:
			fmt.Fprintf(color, "%sArgument %s is not defined.%s\n", "\x1b[31m", arg, "\x1b[0m")
		}
	},
}
