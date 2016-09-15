package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/agupta666/hash/actions"
)

// Router is a basic http router for routes added from console
type Router struct {
}

func (rt *Router) match(path string) (actions.Action, string, bool) {
	// ac, ok := routes[path]

	for k, v := range routes {
		if strings.HasPrefix(path, k) {
			return v, k, true
		}
	}
	return nil, "", false
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	action, pattern, ok := rt.match(r.URL.Path)

	if !ok {
		http.Error(w, "No action defined for this route", http.StatusInternalServerError)
		return
	}

	action.SetPattern(pattern)
	err := action.Exec(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func startDefaultEp(addr string) {
	http.Handle("/", new(Router))
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", "failed to start default http endpoint", err)
		os.Exit(1)
	}
}
