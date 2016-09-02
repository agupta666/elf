package main

import "net/http"

func epHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func startDefaultEp(addr string) {
	http.HandleFunc("/", epHandler)
	http.ListenAndServe(addr, nil)
}
