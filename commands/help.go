package commands

import "fmt"

const helpMain = `
Type: "help <command>" for help on <command>
      "help <tab>" to get a list of possible help topics
      "quit" to exit
`

var helpMap = map[string]string{
	"route": `
  route path action
  summary: add a route with given path and attach the given action
  `,

	"lsrt": `
  lsrt -
  summary: Display the list of all predefined routes
  `,

	"delrt": `
  delrt route
  summary: Remove the specified route
  `,

	"kvset": `
  kvset name [key:value] ...
  summary: Define a new key-value set
  `,

	"lskv": `
  lskv -
  summary: Display the list of all predefined key-value sets
  `,

	"help": `
  help command
  summary: Display help for the given command
  `,

	"quit": `
  quit -
  summary: Stop the server and exit the shell
  `,
}

func helpCmd(args []string) {
	if len(args) == 0 {
		fmt.Print(helpMain)
		return
	}

	helpText, ok := helpMap[args[0]]

	if ok {
		fmt.Print(helpText)
	}

}
