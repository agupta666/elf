package actions

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/agupta666/hash/utils"
)

// UploadAction represents actions which uploads files to a given folder
type UploadAction struct {
	PatternHolder
	Key  string
	Path string
}

func (ua *UploadAction) String() string {
	return fmt.Sprintf("upload[Path=%s, Key=%s]", ua.Path, ua.Key)
}

// NewUploadActionFromExpr creates a new Upload action from an expression
func NewUploadActionFromExpr(p string) (*UploadAction, error) {
	return parseUploadExpr(p)
}

// NewUploadAction creates a new UploadAction
func NewUploadAction(uploadPath string, key string) (*UploadAction, error) {
	uploadPath = strings.TrimSpace(uploadPath)
	key = strings.TrimSpace(key)

	ok, _ := utils.FileExists(uploadPath)
	if !ok {
		return nil, errors.New("specified folder does not exist")
	}

	return &UploadAction{Key: key, Path: uploadPath}, nil
}

// DefaultUploadAction createa an upload action with form file key as upload and path
// as current directory
func DefaultUploadAction() (*UploadAction, error) {
	workDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	return NewUploadAction(workDir, "upload")
}

// NewUploadActionWithPath creates an upload action with form file key as upload and the supplied path
func NewUploadActionWithPath(uploadPath string) (*UploadAction, error) {
	return NewUploadAction(uploadPath, "upload")
}

// Exec executes an upload action
func (ua *UploadAction) Exec(w http.ResponseWriter, r *http.Request) error {
	file, header, err := r.FormFile(ua.Key)

	if err != nil {
		return err
	}

	defer file.Close()
	fPath := path.Join(ua.Path, header.Filename)
	out, err := os.Create(fPath)

	if err != nil {
		return errors.New("Unable to create the file for writing. Check your write access privilege")
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "[%s] upload complete\r\n", header.Filename)
	return nil
}

func parseUploadExpr(s string) (*UploadAction, error) {
	s = strings.TrimPrefix(s, "upload[")
	s = strings.TrimSuffix(s, "]")
	args := strings.Split(s, ",")

	switch len(args) {
	case 0:
		return DefaultUploadAction()
	case 1:
		return NewUploadActionWithPath(args[0])
	case 2:
		return NewUploadAction(args[0], args[1])
	default:
		return nil, errors.New("invalid expression")

	}
}

// SetPattern sets the matched pattern in the action
func (ua *UploadAction) SetPattern(p string) {}
