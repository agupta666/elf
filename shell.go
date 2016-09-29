package main

import (
	"fmt"
	"os"

	"github.com/agupta666/elf/commands"
	shellwords "github.com/mattn/go-shellwords"
	readline "gopkg.in/readline.v1"
)

func processCmd(line string) {
	args, err := shellwords.Parse(line)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: syntax error")
		return
	}

	if len(args) > 0 {
		handler := commands.LookupHandler(args[0])
		handler(args[1:])
	}
}

func startShell() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "elf> ",
		HistoryFile:  ".elf.hist",
		AutoComplete: completer,
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

var completer = readline.NewPrefixCompleter(
	readline.PcItem("route",
		readline.PcItem("/some/path",
			readline.PcItem("!\"\""),
			readline.PcItem("@\"\""),
			readline.PcItem("#\"\""),
			readline.PcItem("^\"\""),
			readline.PcItem("%\"\""),
			readline.PcItem("\"data[]\""),
			readline.PcItem("\"image[]\""),
			readline.PcItem("\"json[]\""),
			readline.PcItem("\"upload[]\""),
			readline.PcItem("\"dir[]\""),
			readline.PcItem("\"dump[]\""),
			readline.PcItem("\"echo[]\""),
		),
	),
	readline.PcItem("lsrt"),
	readline.PcItem("delrt"),
	readline.PcItem("kvset"),
	readline.PcItem("exit"),
)
