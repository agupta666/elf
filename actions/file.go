package actions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/agupta666/hash/utils"
)

// FileAction action represents actions which responds with the contents of a file
type FileAction struct {
	PatternHolder
	Path string
}

func (fa *FileAction) String() string {
	return fmt.Sprintf("@[Path=%s]", fa.Path)
}

// SetPattern sets the matched pattern in the action
func (fa *FileAction) SetPattern(p string) {}

// NewFileAction creates a new FileAction
func NewFileAction(p string) (*FileAction, error) {
	return &FileAction{Path: p}, nil
}

// Exec executes a file action
func (fa *FileAction) Exec(w http.ResponseWriter, r *http.Request) error {
	reader, err := os.Open(fa.Path)
	if err != nil {
		return err
	}

	emitMimeType(fa.Path, w)
	_, err = io.Copy(w, reader)
	if err != nil {
		return err
	}

	return nil
}

func emitMimeType(fname string, w http.ResponseWriter) {
	ext := path.Ext(fname)
	mimeType := utils.TypeByExtension(ext)
	w.Header().Set("Content-Type", mimeType)
}
