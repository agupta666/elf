package actions

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/agupta666/hash/store"
)

// JSONAction action represents actions which responds with JSON
type JSONAction struct {
	Name string
}

// NewJSONActionFromExpr creates a new json action from an expression
func NewJSONActionFromExpr(p string) (*JSONAction, error) {
	return parseJSONExpr(p)
}

func parseJSONExpr(s string) (*JSONAction, error) {
	s = strings.TrimPrefix(s, "json[")
	s = strings.TrimSuffix(s, "]")
	args := strings.Split(s, ",")

	switch len(args) {
	case 0:
		return nil, errors.New("invalid expression")
	case 1:
		return &JSONAction{Name: args[0]}, nil
	default:
		return nil, errors.New("invalid expression")

	}
}

// Exec executes a json action
func (ja *JSONAction) Exec(w http.ResponseWriter, r *http.Request) error {

	kvs := store.GetKVSet(ja.Name)

	writeMimeType(".json", w)
	b, err := json.Marshal(kvs)

	if err != nil {
		return err
	}

	w.Write(b)
	return nil
}
