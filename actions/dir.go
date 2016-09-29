package actions

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/agupta666/wish/utils"
)

// DirAction represents actions which serve static files out of a given folder
type DirAction struct {
	PatternHolder
	Root    string
	pattern string
}

func (da *DirAction) String() string {
	return fmt.Sprintf("dir[Root=%s]", da.Root)
}

func (da *DirAction) resolvePath(p string) string {
	trimmedPath := strings.TrimPrefix(p, da.pattern)
	return path.Join(da.Root, trimmedPath)
}

// Exec executes a dir action
func (da *DirAction) Exec(w http.ResponseWriter, r *http.Request) error {
	filePath := da.resolvePath(r.URL.Path)
	http.ServeFile(w, r, filePath)
	return nil
}

// NewDirActionFromExpr creates a new Dir action from an expression
func NewDirActionFromExpr(p string) (*DirAction, error) {
	return parseDirExpr(p)
}

// DefaultDirAction creates a dir action which serves files from current working directory
func DefaultDirAction() (*DirAction, error) {
	rootDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	return NewDirAction(rootDir)
}

// NewDirAction creates a new dir action which serves files from the supplied folder
func NewDirAction(root string) (*DirAction, error) {
	ok, _ := utils.FileExists(root)
	if !ok {
		return nil, errors.New("specified folder does not exist")
	}

	return &DirAction{Root: root}, nil
}

func parseDirExpr(s string) (*DirAction, error) {
	s = strings.TrimPrefix(s, "dir[")
	s = strings.TrimSuffix(s, "]")
	args := utils.SplitAndTrim(s, ",")

	switch len(args) {
	case 0:
		return DefaultDirAction()
	case 1:
		return NewDirAction(args[0])
	default:
		return nil, errors.New("invalid expression")

	}

}

// SetPattern sets the matched pattern in the action
func (da *DirAction) SetPattern(p string) {
	da.pattern = p
}
