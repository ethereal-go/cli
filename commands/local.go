package commands

import (
	"github.com/spf13/cobra"
	//"github.com/ethereal-go/ethereal"
	"fmt"
)

var cmdLocale = &cobra.Command{
	Use:   "locale",
	Short: "Localization management",
	Long: ``,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		//arg := args[0]
		//switch arg {
		//case "fill" :
		//	//ethereal.I18nGraphQL().Fill()
		//	fmt.Println("Success fill locale in database! Good job!")
		//default:
		//	fmt.Println("Argument '" + arg + "' is not defined. ")
		//}
	},
}
