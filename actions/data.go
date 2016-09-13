package actions

import (
	"crypto/rand"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// DataAction action represents actions which responds with random bytes of data of given size and format
type DataAction struct {
	Size int
	Type string
	Name string
}

// HasName returns true if the data action has a name
func (da *DataAction) HasName() bool {
	return len(da.Name) != 0
}

// NewDataActionFromExpr creates a new DataAction from an expression
func NewDataActionFromExpr(p string) (*DataAction, error) {
	return parseDataExpr(p)
}

// NewDataAction creates a data action with size and type
func NewDataAction(size string, ext string) (*DataAction, error) {
	sz, err := strconv.Atoi(size)

	if err != nil {
		return nil, errors.New("size must be integer")
	}

	extn := strings.TrimSpace(ext)

	if !strings.HasPrefix(ext, ".") {
		extn = "." + ext
	}

	return &DataAction{Size: sz, Type: extn}, nil
}

// NewDataActionWithName adds a file name to the data action
func NewDataActionWithName(size string, ext string, name string) (*DataAction, error) {
	dAction, err := NewDataAction(size, ext)

	if err != nil {
		return nil, err
	}

	dAction.Name = strings.TrimSpace(name)
	return dAction, nil
}

func parseDataExpr(s string) (*DataAction, error) {
	s = strings.TrimPrefix(s, "data[")
	s = strings.TrimSuffix(s, "]")
	args := strings.Split(s, ",")

	switch len(args) {
	case 0, 1:
		return nil, errors.New("invalid expression")
	case 2:
		return NewDataAction(args[0], args[1])
	case 3:
		return NewDataActionWithName(args[0], args[1], args[2])
	default:
		return nil, errors.New("invalid expression")

	}
}

// Exec executes a data action
func (da *DataAction) Exec(w http.ResponseWriter, r *http.Request) error {

	data, err := makeData(da.Size)

	if err != nil {
		return err
	}
	if da.HasName() {
		writeFileName(da.Name, da.Type, w)
	}

	writeMimeType(da.Type, w)

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func makeData(sz int) ([]byte, error) {
	data := make([]byte, sz)

	_, err := rand.Read(data)
	if err != nil {
		return make([]byte, 0), err
	}

	return data, nil
}
