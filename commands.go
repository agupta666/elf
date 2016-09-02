package main

import (
	"fmt"
	"os"
)

type CmdHandler func(args []string)

func routeCmd(args []string) {

}

func exitCmd(args []string) {
	os.Exit(1)
}

func notFoundHandler(args []string) {
	fmt.Fprintln(os.Stderr, "ERROR:", "command not found")
}

var commands = map[string]CmdHandler{
	"route": routeCmd,
	"exit":  exitCmd,
}

func lookupHandler(cmd string) CmdHandler {
	handler, ok := commands[cmd]

	if ok {
		return handler
	}
	return notFoundHandler
}
