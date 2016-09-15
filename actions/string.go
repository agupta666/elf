package actions

import (
	"fmt"
	"net/http"
)

// StringAction action represents actions which responds with a string
type StringAction struct {
	PatternHolder
	Value string
}

func (sa *StringAction) String() string {
	return fmt.Sprintf("[Value=%s]", sa.Value)
}

// NewStringAction creates a new StringAction
func NewStringAction(s string) (*StringAction, error) {
	return &StringAction{Value: s}, nil
}

// Exec executes a string action
func (sa *StringAction) Exec(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte(sa.Value))
	return nil
}

// SetPattern sets the matched pattern in the action
func (sa *StringAction) SetPattern(p string) {}
