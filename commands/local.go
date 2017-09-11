package commands

import (
	"fmt"
	"github.com/ethereal-go/ethereal"
	"github.com/spf13/cobra"
)

var cmdLocale = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "locale",
	Short: "Localization management",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		arg := args[0]
		switch arg {
		case "fill":
			ethereal.I18nGraphQL().Fill()
			fmt.Fprintf(color, "%sSuccess fill locale in database! Good job %s\n!", "\x1b[32m", "\x1b[0m")
		default:
			fmt.Fprintf(color, "%sArgument %s is not defined.%s\n", "\x1b[31m", arg, "\x1b[0m")
		}
	},
}
