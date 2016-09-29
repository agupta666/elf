package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/agupta666/wish/store"
)

func kvsetCmd(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "ERROR:", "syntax error")
		fmt.Fprintln(os.Stderr, "USAGE:", "kvset name [key=value]...")
		return
	}
	keyName := args[0]
	kvs := make(store.KVSet)

	for _, arg := range args[1:] {
		xs := strings.Split(arg, ":")
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
