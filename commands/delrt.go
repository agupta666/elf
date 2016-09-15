package commands

import (
	"fmt"
	"os"

	"github.com/agupta666/hash/router"
)

func deleteRouteCmd(args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "syntax error")
		return
	}

	router.DeleteRoute(args[0])
}
