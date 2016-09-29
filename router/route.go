package router

import "github.com/agupta666/elf/actions"

var routes = make(map[string]actions.Action)

// AddRoute adds a route to the routing table
func AddRoute(path string, action actions.Action) {
	routes[path] = action
}

// Routes returns the routing table
func Routes() map[string]actions.Action {
	return routes
}

// DeleteRoute deletes a route from the routing table
func DeleteRoute(path string) {
	delete(routes, path)
}
