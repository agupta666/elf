package actions

import (
	"net/http"
	"strings"
)

// PatternHolder holds the matched pattern for later use
type PatternHolder interface {
	SetPattern(p string)
}

// Action represents actions to be taken when a reuest matches a given route
type Action interface {
	PatternHolder
	Exec(w http.ResponseWriter, r *http.Request) error
}

// GetAction creates an action object based on its identifier
func GetAction(arg string) (Action, error) {
	actionIdentifier := arg[0]

	switch actionIdentifier {
	case '!':
		return NewShellAction(arg[1:])
	case '@':
		return NewFileAction(arg[1:])
	case '#':
		return NewMarkdownAction(arg[1:])
	case '^':
		return NewRedirectAction(arg[1:])
	case '%':
		return NewForwardAction(arg[1:])
	default:
		return getBuiltinAction(arg)
	}

}

func getBuiltinAction(arg string) (Action, error) {
	switch {
	case strings.HasPrefix(arg, "data["):
		return NewDataActionFromExpr(arg)
	case strings.HasPrefix(arg, "image["):
		return NewImageActionFromExpr(arg)
	case strings.HasPrefix(arg, "json["):
		return NewJSONActionFromExpr(arg)
	case strings.HasPrefix(arg, "upload["):
		return NewUploadActionFromExpr(arg)
	case strings.HasPrefix(arg, "dir["):
		return NewDirActionFromExpr(arg)
	case strings.HasPrefix(arg, "dump["):
		return NewDumpActionFromExpr(arg)
	default:
		return NewStringAction(arg)
	}
}
