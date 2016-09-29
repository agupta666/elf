package commands

import (
	"fmt"

	"github.com/agupta666/elf/router"
)

func lsroutesCmd(args []string) {
	for k, v := range router.Routes() {
		fmt.Println(k, ":", v)
	}
}
