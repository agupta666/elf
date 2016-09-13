package actions

import (
	"fmt"
	"net/http"
)

// RedirectAction action represents actions which redirect to a given url
type RedirectAction struct {
	URL string
}

func (ra *RedirectAction) String() string {
	return fmt.Sprintf("^[URL=%s]", ra.URL)
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
