package commands

import (
	"fmt"
	"os"
)

var cmdHelp = &Command{
	Run:   runHelp,
	Usage: "help [command]",
	Short: "Show help",
	Long:  "Shows usage for a command.",
}

func init() {
	CmdRunner.Use(cmdHelp)
}

func runHelp(cmd *Command, args *Args) {
	if args.IsParamsEmpty() {
		printUsage()

		os.Exit(0)
	}

	for _, cmd := range CmdRunner.All() {
		if cmd.Name() == args.FirstParam() {
			cmd.PrintUsage()

			os.Exit(0)
		}
	}
}

var helpText = `usage: neocities <command> [<args>]

Commands:
   upload    Upload a file to Neocities
   version   Show Neocities API client version

Help for a specific command:
   help [command]
`

func printUsage() {
	fmt.Print(helpText)
}
