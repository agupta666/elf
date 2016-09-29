package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/agupta666/wish/router"
)

func startDefaultEp(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	http.Handle("/", new(router.Router))

	fmt.Fprintln(os.Stdout, "starting default http endpoint", addr)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", "failed to start default http endpoint", err)
		os.Exit(1)
	}
}
