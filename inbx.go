package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jackvhall/inbx/config"
)

const (
	Version = "0.0.1"
)

func main() {
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)
	configCmd := flag.NewFlagSet("config", flag.ExitOnError)

	newCmd.String("note", "", "note to save")
	configCmd.String("file", "", "specify configuration (yaml) file to use instead of default")

	if len(os.Args) < 2 {
		fmt.Println("expected 'new' or 'config' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		newCmd.Parse(os.Args[2:])
		note := newCmd.Lookup("note").Value.String()
		fmt.Println("note: ", note)
	case "config":
		configCmd.Parse(os.Args[2:])
		configArgs := configCmd.Lookup("file").Value.String()
		config.Configure(configArgs)
	default:
		fmt.Println("expected 'new' or 'config' subcommand")
		os.Exit(1)
	}
}
