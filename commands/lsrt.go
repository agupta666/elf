package commands

import (
	"fmt"

	"github.com/agupta666/hash/router"
)

func lsroutesCmd(args []string) {
	for k, v := range router.Routes() {
		fmt.Println(k, ":", v)
	}
}
