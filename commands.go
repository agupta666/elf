package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/agupta666/hash/actions"
	"github.com/agupta666/hash/store"
)

// CmdHandler is a type for all command handler functions
type CmdHandler func(args []string)

func routeCmd(args []string) {
	act, err := actions.GetAction(args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	addRoute(args[0], act)
}

func exitCmd(args []string) {
	os.Exit(1)
}

func lsroutesCmd(args []string) {
	for k, v := range routes {
		fmt.Println(k, ":", v)
	}
}

func kvsetCmd(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "ERROR:", "syntax error")
	}
	keyName := args[0]
	kvs := make(store.KVSet)

	for _, arg := range args[1:] {
		xs := strings.Split(arg, "=")
		k := strings.TrimSpace(xs[0])
		v := strings.TrimSpace(xs[1])
		kvs[k] = v
	}
	store.SaveKVSet(keyName, kvs)
}

func notFoundHandler(args []string) {
	fmt.Fprintln(os.Stderr, "ERROR:", "command not found")
}

var commandsMap = map[string]CmdHandler{
	"route":    routeCmd,
	"lsroutes": lsroutesCmd,
	"kvset":    kvsetCmd,
	"exit":     exitCmd,
}

func lookupHandler(cmd string) CmdHandler {
	handler, ok := commandsMap[cmd]

	if ok {
		return handler
	}
	return notFoundHandler
}
