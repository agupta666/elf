package actions

import (
	"net/http"
	"os/exec"
	"strings"
)

// ShellAction action represents actions which responds with output of a shell command
type ShellAction struct {
	Cmd string
}

// NewShellAction creates a new ShellAction
func NewShellAction(s string) *ShellAction {
	return &ShellAction{Cmd: s}
}

// Exec executes a string action
func (sa *ShellAction) Exec(w http.ResponseWriter, r *http.Request) error {
	xs := strings.Split(sa.Cmd, " ")
	cmd := exec.Command(xs[0], xs[1:]...)
	cmd.Stdout = w
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
