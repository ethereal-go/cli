package commands

import (
	"fmt"
	pkgDatabase "github.com/ethereal-go/ethereal/root/database"
	"github.com/ethereal-go/ethereal/root/i18n/storage"
	"github.com/spf13/cobra"
)

var cmdLocale = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "locale",
	Short: "Localization management",
	Long:  `
Flags :
	- fill : Fill your source with the localization from the file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
		initPathFileLanguage()
		arg := args[0]
		switch arg {
		case "fill":
			db, err := pkgDatabase.FactoryDatabase(database)
			if err != nil {
				fmt.Fprintf(color, "%s%s%s\n", "\x1b[31m", err, "\x1b[0m")
			}

			gorm := db.Parse().Connection()

			storage := storage.LocaleStorage{}.EstablishConnection(gorm)
			storage.Add(locale)
			fmt.Fprintf(color, "%sSuccess fill locale in database! Good job!  %s\n", "\x1b[32m", "\x1b[0m")
		default:
			fmt.Fprintf(color, "%sArgument %s is not defined.%s\n", "\x1b[31m", arg, "\x1b[0m")
		}
	},
}
