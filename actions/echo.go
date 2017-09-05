package actions

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// EchoAction echoes back the request body
type EchoAction struct {
	Status int
	Type   string
	Body   []byte
}

func (ea *EchoAction) String() string {
	contentType := ea.Type
	body := string(ea.Body)
	if contentType == "" {
		contentType = "<request-content-type>"
	}

	if body == "" {
		body = "<request-body>"
	}
	return fmt.Sprintf("echo[Status=%d, Type=%s, Body=%s]", ea.Status, contentType, body)
}

// Exec executes a echo action
func (ea *EchoAction) Exec(w http.ResponseWriter, r *http.Request) error {

	if len(ea.Type) == 0 {
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	} else {
		w.Header().Set("Content-Type", ea.Type)
	}

	w.WriteHeader(ea.Status)

	if len(ea.Body) == 0 {
		io.Copy(w, r.Body)
	} else {
		switch ea.Body[0] {
		case '@':
			writeFile(string(ea.Body[1:]), w)
		default:
			w.Write(ea.Body)
		}

	}

	return nil
}

// SetPattern sets the matched pattern in the action
func (ea *EchoAction) SetPattern(p string) {}

// NewEchoActionFromExpr creates a new EchoAction from an expression
func NewEchoActionFromExpr(exp string) (*EchoAction, error) {
	ac, err := parseExpr(exp, "echo", func(args []string) (Action, error) {

		switch len(args) {
		case 0:
			return NewDefaultEchoAction()
		case 1:
			return NewEchoActionWithStatus(args[0])
		case 2:
			return NewEchoActionWithStatusAndType(args[0], args[1])
		case 3:
			return NewEchoAction(args[0], args[1], args[2])
		default:
			return nil, errors.New("invalid expression")
		}
	})

	return ac.(*EchoAction), err
}

// NewDefaultEchoAction constructs a new echo action with status 200, contentType of the request and body of the request
func NewDefaultEchoAction() (*EchoAction, error) {
	return NewEchoAction("200", "", "")
}

// NewEchoActionWithStatus constructs a new echo action with status
func NewEchoActionWithStatus(status string) (*EchoAction, error) {
	return NewEchoAction(status, "", "")
}

// NewEchoActionWithStatusAndType constructs a new echo action with status, contentType
func NewEchoActionWithStatusAndType(status string, contentType string) (*EchoAction, error) {
	return NewEchoAction(status, contentType, "")
}

// NewEchoAction constructs a new echo action with status, contentType and body
func NewEchoAction(status string, contentType string, body string) (*EchoAction, error) {
	nStatus, err := strconv.Atoi(status)

	if err != nil {
		return nil, err
	}

	return &EchoAction{
		Status: nStatus,
		Type:   contentType,
		Body:   []byte(body),
	}, nil
}
