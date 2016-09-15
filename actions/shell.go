package actions

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

// ShellAction action represents actions which responds with output of a shell command
type ShellAction struct {
	PatternHolder
	Cmd string
}

func (sa *ShellAction) String() string {
	return fmt.Sprintf("![Command=%s]", sa.Cmd)
}

// NewShellAction creates a new ShellAction
func NewShellAction(s string) (*ShellAction, error) {
	return &ShellAction{Cmd: s}, nil
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

// SetPattern sets the matched pattern in the action
func (sa *ShellAction) SetPattern(p string) {}
