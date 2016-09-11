package actions

import "net/http"

// StringAction action represents actions which responds with a string
type StringAction struct {
	Value string
}

// NewStringAction creates a new StringAction
func NewStringAction(s string) *StringAction {
	return &StringAction{Value: s}
}

// Exec executes a string action
func (sa *StringAction) Exec(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte(sa.Value))
	return nil
}
