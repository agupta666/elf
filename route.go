package main

var routes = make(map[string]Action)

func addRoute(path string, action Action) {
	routes[path] = action
}
