package actions

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

// MarkdownAction action represents actions which responds with html output of markdown files
type MarkdownAction struct {
	Path string
}

// NewMarkdownAction creates a new MarkdownAction
func NewMarkdownAction(s string) (*MarkdownAction, error) {
	return &MarkdownAction{Path: s}, nil
}

// Exec executes a string action
func (ma *MarkdownAction) Exec(w http.ResponseWriter, r *http.Request) error {

	reader, err := os.Open(ma.Path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return err
	}

	w.Write(blackfriday.MarkdownBasic(data))

	return nil
}
