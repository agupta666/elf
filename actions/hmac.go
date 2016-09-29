package actions

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"crypto/hmac"
	"crypto/sha1"

	"github.com/agupta666/elf/store"
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

func getCanonicalString(r *http.Request) string {
	httpMethod := r.Method
	contentType := r.Header.Get("")
	contentMD5 := r.Header.Get("")
	requestURI := r.URL
	timestamp := r.Header.Get("Date")

	return fmt.Sprintf("%s,%s,%s,%s,%s", httpMethod, contentType, contentMD5, requestURI, timestamp)
}

func computeSignature(r *http.Request, secret string) (string, error) {

	canStr := getCanonicalString(r)
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(canStr))

	signature := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature), nil
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
