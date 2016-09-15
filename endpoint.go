package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/agupta666/hash/router"
)

func startDefaultEp(addr string) {
	http.Handle("/", new(router.Router))
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", "failed to start default http endpoint", err)
		os.Exit(1)
	}
}
