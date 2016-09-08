package actions

import "net/http"

// Action represents actions to be taken when a reuest matches a given route
type Action interface {
	Exec(w http.ResponseWriter) error
}

// GetAction creates an action object based on its identifier
func GetAction(arg string) Action {
	actionIdentifier := arg[0]

	switch actionIdentifier {
	case '!':
		return NewShellAction(arg[1:])
	case '@':
		return NewFileAction(arg[1:])
	default:
		return NewStringAction(arg)
	}

}
