package commands

import (
	"fmt"

	"github.com/agupta666/wish/router"
)

func lsroutesCmd(args []string) {
	for k, v := range router.Routes() {
		fmt.Println(k, ":", v)
	}
}
