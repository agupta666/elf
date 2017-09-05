package commands

import (
	"fmt"
	"github.com/agupta666/elf/router"
)

func deleteRouteCmd(args []string) {

	if len(args) != 1 {
		fmt.Println("ERROR:", "wrong number of arguments for 'delrt' command")
		return
	}

	router.DeleteRoute(args[0])
}
