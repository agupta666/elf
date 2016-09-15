package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/agupta666/hash/store"
)

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
