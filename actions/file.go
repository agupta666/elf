package actions

import (
	"io"
	"net/http"
	"os"
)

// FileAction action represents actions which responds with the contents of a file
type FileAction struct {
	Path string
}

// NewFileAction creates a new FileAction
func NewFileAction(p string) *FileAction {
	return &FileAction{Path: p}
}

// Exec executes a file action
func (fa *FileAction) Exec(w http.ResponseWriter) error {
	r, err := os.Open(fa.Path)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	return nil
}
