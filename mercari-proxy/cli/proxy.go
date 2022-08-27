package main

import (
	"os"

	"sagungw/mercari/cli/cmd"

	"github.com/sagungw/gotrunks/log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.From("main", "main").Error(err.Error())
		os.Exit(-1)
	}
}
