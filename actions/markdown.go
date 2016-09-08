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
func NewMarkdownAction(s string) *MarkdownAction {
	return &MarkdownAction{Path: s}
}

// Exec executes a string action
func (ma *MarkdownAction) Exec(w http.ResponseWriter) error {

	r, err := os.Open(ma.Path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(r)

	if err != nil {
		return err
	}

	w.Write(blackfriday.MarkdownBasic(data))

	return nil
}
