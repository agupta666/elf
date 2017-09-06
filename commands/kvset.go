package commands

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/agupta666/elf/store"
)

func kvsetCmd(args []string) {
	if len(args) < 2 {
		fmt.Println("ERROR:", "wrong number of arguments for 'kvset' command")
		return
	}
	keyName := args[0]
	kvs := make(store.KVSet)

	for _, arg := range args[1:] {
		xs := strings.Split(arg, ":")
		if len(xs) != 2 {
			fmt.Println("ERROR:", "syntax error in 'kvset' command")
			return
		}

		k := strings.TrimSpace(xs[0])
		v := strings.TrimSpace(xs[1])
		kvs[k] = v
	}
	store.SaveKVSet(keyName, kvs)
}

func lskvCmd(args []string) {
	kvss := store.All()
	for k, v := range kvss {
		s, _ := json.Marshal(v)
		fmt.Println(k, string(s))
	}
}
