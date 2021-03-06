package commands

import (
	"fmt"
	"github.com/ethereal-go/ethereal/utils"
	"github.com/spf13/cobra"
	"path"
	"runtime"
)

var cmdAdd = &cobra.Command{
	Aliases: []string{"make"},
	Use:     "add",
	Short:   "Add or create files(or anything else)",
	Long:    ``,
	Example: `
	 - add config $GOPATH/src/... (where is your project)
	 - add env $GOPATH/src/... (where is your project)
	`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			nameConfFile = "app.json"
			envName      = ".env"
		)
		typeAction, NewPath := args[0], args[1]
		_, filename, _, _ := runtime.Caller(0)

		switch typeAction {
		case "config":
			if utils.DirExist(NewPath) {
				utils.Copy(NewPath+nameConfFile, path.Dir(filename)+"/../stubs/config/"+nameConfFile)
				fmt.Fprintf(color, "%sCreate default configuration file : %s %s\n", "\x1b[32m", "\x1b[0m", NewPath+nameConfFile)
			} else {
				fmt.Fprintf(color, "%sDirectory %s is not exist.%s\n", "\x1b[31m", NewPath, "\x1b[0m")
			}
		case "env":
			if utils.DirExist(NewPath) {
				utils.Copy(NewPath+envName, path.Dir(filename)+"/../stubs/config/"+envName)
				fmt.Fprintf(color, "%sCreate default env file : %s %s\n", "\x1b[32m", "\x1b[0m", NewPath+envName)
			} else {
				fmt.Fprintf(color, "%sDirectory %s is not exist.%s\n", "\x1b[31m", NewPath, "\x1b[0m")
			}
		default:
			fmt.Fprintf(color, "%sArgument %s is not defined.%s\n", "\x1b[31m", typeAction, "\x1b[0m")
		}
	},
}
