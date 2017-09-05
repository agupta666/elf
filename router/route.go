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

// RouteNames returns the names of configured routes
func RouteNames(args string) []string {
	names := make([]string, 0)
	for k := range routes {
		names = append(names, k)
	}

	return names
}
