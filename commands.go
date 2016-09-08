package main

import (
	"fmt"
	"os"

	"github.com/agupta666/hash/actions"
)

// CmdHandler is a type for all command handler functions
type CmdHandler func(args []string)

func routeCmd(args []string) {
	addRoute(args[0], actions.GetAction(args[1]))
}

func exitCmd(args []string) {
	os.Exit(1)
}

func lsroutesCmd(args []string) {
	for k, v := range routes {
		fmt.Println(k, ":", v)
	}
}

func notFoundHandler(args []string) {
	fmt.Fprintln(os.Stderr, "ERROR:", "command not found")
}

var commandsMap = map[string]CmdHandler{
	"route":    routeCmd,
	"lsroutes": lsroutesCmd,
	"exit":     exitCmd,
}

func lookupHandler(cmd string) CmdHandler {
	handler, ok := commandsMap[cmd]

	if ok {
		return handler
	}
	return notFoundHandler
}
