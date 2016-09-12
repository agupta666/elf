package actions

import "net/http"

// RedirectAction action represents actions which responds with a string
type RedirectAction struct {
	URL string
}

// NewRedirectAction creates a new RedirectAction
func NewRedirectAction(s string) (*RedirectAction, error) {
	return &RedirectAction{URL: s}, nil
}

// Exec executes a redirect action
func (ra *RedirectAction) Exec(w http.ResponseWriter, r *http.Request) error {

	http.Redirect(w, r, ra.URL, http.StatusTemporaryRedirect)
	return nil
}
