package commands

import (
	"fmt"

	"github.com/agupta666/elf/actions"
	"github.com/agupta666/elf/router"
)

func routeCmd(args []string) {
	act, err := actions.GetAction(args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	router.AddRoute(args[0], act)
}
