package actions

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/agupta666/hash/store"
)

// HMACAction action represents actions which does HMAC authentication on the attached route
type HMACAction struct {
	PatternHolder
	Name string
}

func (ha *HMACAction) String() string {
	return fmt.Sprintf("hmac[Name=%s]", ha.Name)
}

// NewHMACActionFromExpr creates a new hmac action from an expression
func NewHMACActionFromExpr(p string) (*HMACAction, error) {
	return parseHMACExpr(p)
}

func parseHMACExpr(exp string) (*HMACAction, error) {
	ac, err := parseExpr(exp, "hmac", func(args []string) (Action, error) {

		switch len(args) {
		case 0:
			return nil, errors.New("invalid expression")
		case 1:
			return &HMACAction{Name: args[0]}, nil
		default:
			return nil, errors.New("invalid expression")

		}
	})
	return ac.(*HMACAction), err
}

func getSecretKey(kvs store.KVSet, clientID string) (string, error) {
	secret, ok := kvs[clientID]

	if !ok {
		return "", errors.New("secret not found")
	}

	return secret, nil
}

func computeSignature(r *http.Request, secret string) (string, error) {
	fmt.Println(r)
	return "", nil
}

// Exec executes a HMAC action
func (ha *HMACAction) Exec(w http.ResponseWriter, r *http.Request) error {

	kvs := store.GetKVSet(ha.Name)
	clientID, signature, err := parseAuthHeader(r)

	if err != nil {
		return err
	}

	clientSecret, err := getSecretKey(kvs, clientID)

	if err != nil {
		return err
	}

	computedSignature, err := computeSignature(r, clientSecret)

	if err != nil {
		return err
	}

	if signature != computedSignature {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return nil
	}

	w.Write([]byte("authorized"))
	return nil
}

// SetPattern sets the matched pattern in the action
func (ha *HMACAction) SetPattern(p string) {}
