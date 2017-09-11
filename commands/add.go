package commands

import (
	"fmt"
	"github.com/ethereal-go/ethereal/utils"
	"github.com/spf13/cobra"
	"path"
	"runtime"
)

var addLocale = &cobra.Command{
	Use:     "add",
	Short:   "Add or create files(or anything else)",
	Long:    ``,
	Example: "add config $GOPATH/src/...",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var nameConfFile = "app.json"

		typeAction := args[0]

		switch typeAction {
		case "config":
			if utils.DirExist(args[1]) {
				_, filename, _, _ := runtime.Caller(0)
				utils.Copy(args[1]+nameConfFile, path.Dir(filename)+"/../stubs/config/"+nameConfFile)

				fmt.Fprintf(color, "%sCreate default configuration file : %s %s\n", "\x1b[32m", "\x1b[0m", args[1]+nameConfFile)
			} else {
				fmt.Fprintf(color, "%sDirectory %s is not exist.%s\n", "\x1b[31m", args[1], "\x1b[0m")
			}
		default:
			fmt.Fprintf(color, "%sArgument %s is not defined.%s\n", "\x1b[31m", typeAction, "\x1b[0m")
		}
	},
}
