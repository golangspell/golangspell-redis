package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell-redis/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["redisinit"] = runredisinit
}

func runredisinit(cmd *cobra.Command, args []string) {
	// Example on how to deal when the expected arguments were not provided
	// if len(args) != 2 {
	// 	fmt.Println(`The command redisinit requires n arguments
	// 	Args:
	// 	[PUT HERE THE NEW COMMAND ARGS DESCRIPTION]
		
	// 	Syntax: 
	// 	golangspell [PUT HERE THE NEW COMMAND SYNTAX]
		
	// 	Examples:
	// 	[PUT HERE THE NEW COMMAND EXAMPLES IF NEEDED]`)
	// 	return
	// }

	//Here your template, hosted on the folder "templates" is rendered 
	err := usecase.RenderredisinitTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return
	}
	//TODO: Create your additional logic here
	fmt.Println("redisinit executed!")
}
