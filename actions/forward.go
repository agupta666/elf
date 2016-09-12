package actions

import (
	"io"
	"net/http"
)

// ForwardAction action represents actions which forwards the request to a given url
type ForwardAction struct {
	URL string
}

// NewForwardAction creates a new ForwardAction
func NewForwardAction(s string) (*ForwardAction, error) {
	return &ForwardAction{URL: s}, nil
}

// Exec executes a foreward action
func (fa *ForwardAction) Exec(w http.ResponseWriter, r *http.Request) error {

	req, err := http.NewRequest(r.Method, fa.URL, r.Body)

	if err != nil {
		return err
	}

	req.Header = copyHeader(r.Header)

	ua := r.Header.Get("User-Agent")

	req.Header.Set("X-Forwarded-User-Agent", ua)
	req.Header.Set("User-Agent", "Hash")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)

	if err != nil {
		return err
	}
	return nil
}

func copyHeader(h http.Header) http.Header {
	nh := make(http.Header)
	for key, value := range h {
		nh[key] = value
	}
	return nh
}
