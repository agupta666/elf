package main

import "net/http"

func epHandler(w http.ResponseWriter, r *http.Request) {
	action, ok := routes[r.URL.Path]

	if !ok {
		http.Error(w, "No action defined for this route", http.StatusInternalServerError)
		return
	}

	err := action.Exec(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func startDefaultEp(addr string) {
	http.HandleFunc("/", epHandler)
	http.ListenAndServe(addr, nil)
}
