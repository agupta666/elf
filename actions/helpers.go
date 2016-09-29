package actions

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/agupta666/wish/utils"
)

func writeFileName(name string, ext string, w http.ResponseWriter) {
	fileName := fmt.Sprintf("%s%s", name, ext)
	value := fmt.Sprintf("attachment; filename=%s", fileName)
	w.Header().Set("Content-Disposition", value)
}

func writeMimeType(ext string, w http.ResponseWriter) {
	mimeType := utils.TypeByExtension(ext)
	w.Header().Set("Content-Type", mimeType)
}

// ParseComplete handlers a called after the expression parsing is complete
type ParseComplete func(args []string) (Action, error)

func parseExpr(expr, name string, handler ParseComplete) (Action, error) {
	expr = strings.TrimPrefix(expr, name+"[")
	expr = strings.TrimSuffix(expr, "]")
	args := utils.SplitAndTrim(expr, ",")

	return handler(args)
}

func writeFile(filePath string, w http.ResponseWriter) error {
	reader, err := os.Open(filePath)

	if err != nil {
		return err
	}

	_, err = io.Copy(w, reader)
	if err != nil {
		return err
	}

	return nil
}

// [auth scheme] [access_id]:[base64 encoded signature]
func parseAuthHeader(r *http.Request) (string, string, error) {
	authHeader := r.Header.Get("Authorization")

	if len(authHeader) == 0 {
		return "", "", errors.New("auth header not provided")
	}

	sections := strings.Split(authHeader, " ")

	if len(sections) != 2 {
		return "", "", errors.New("invalid auth header")
	}

	creds := strings.Split(sections[1], ":")

	if len(creds) != 2 {
		return "", "", errors.New("invalid auth header")
	}

	return creds[0], creds[1], nil
}
