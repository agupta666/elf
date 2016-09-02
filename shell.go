package main

import (
	"strings"

	readline "gopkg.in/readline.v1"
)

func processCmd(line string) {
	args := strings.Split(line, " ")
	handler := lookupHandler(args[0])
	handler(args[1:])
}

func startShell() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      "hash> ",
		HistoryFile: ".hash.hist",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		processCmd(line)
	}
}
