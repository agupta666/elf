package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Action represents actions to be taken when a reuest matches a given route
type Action interface {
	Exec() ([]byte, error)
}

// StringAction action represents actions which responds with a string
type StringAction struct {
	Value string
}

// NewStringAction creates a new StringAction
func NewStringAction(s string) *StringAction {
	return &StringAction{Value: s}
}

// Exec executes a string action
func (sa *StringAction) Exec() ([]byte, error) {
	return []byte(sa.Value), nil
}

// ShellAction action represents actions which responds with output of a shell command
type ShellAction struct {
	Cmd string
}

// NewShellAction creates a new ShellAction
func NewShellAction(s string) *ShellAction {
	return &ShellAction{Cmd: s}
}

// Exec executes a string action
func (sa *ShellAction) Exec() ([]byte, error) {
	xs := strings.Split(sa.Cmd, " ")
	cmd := exec.Command(xs[0], xs[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return make([]byte, 0), err
	}

	return out.Bytes(), nil
}

// FileAction action represents actions which responds with the contents of a file
type FileAction struct {
	Path string
}

// NewFileAction creates a new FileAction
func NewFileAction(p string) *FileAction {
	return &FileAction{Path: p}
}

// Exec executes a file action
func (fa *FileAction) Exec() ([]byte, error) {
	r, err := os.Open(fa.Path)
	if err != nil {
		return make([]byte, 0), err
	}
	b, err := ioutil.ReadAll(r)

	if err != nil {
		return make([]byte, 0), err
	}
	return b, nil
}

func getAction(arg string) Action {
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
