package main

import (
	"fmt"

	"github.com/golangspell/golangspell-redis/cmd"
	_ "github.com/golangspell/golangspell-redis/config"
	_ "github.com/golangspell/golangspell-redis/gateway/customlog"
	_ "github.com/golangspell/golangspell-redis/gateway/template"
	_ "github.com/golangspell/golangspell/gateway/filesystem"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("An error occurred while executing the command. Message: %s\n", err.Error())
	}
}
