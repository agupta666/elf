package main

import "github.com/agupta666/hash/actions"

var routes = make(map[string]actions.Action)

func addRoute(path string, action actions.Action) {
	routes[path] = action
}
