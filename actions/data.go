package actions

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/agupta666/wish/utils"
)

// DataAction action represents actions which responds with random bytes of data of given size and format
type DataAction struct {
	PatternHolder
	Size int
	Type string
	Name string
}

func (da *DataAction) String() string {
	return fmt.Sprintf("data[Size=%d, Type=%s, Name=%s]", da.Size, da.Type, da.Name)
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

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	return &DataAction{Size: sz, Type: ext}, nil
}

// NewDataActionWithName adds a file name to the data action
func NewDataActionWithName(size string, ext string, name string) (*DataAction, error) {
	dAction, err := NewDataAction(size, ext)

	if err != nil {
		return nil, err
	}

	dAction.Name = name
	return dAction, nil
}

func parseDataExpr(s string) (*DataAction, error) {
	s = strings.TrimPrefix(s, "data[")
	s = strings.TrimSuffix(s, "]")
	args := utils.SplitAndTrim(s, ",")

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

// SetPattern sets the matched pattern in the action
func (da *DataAction) SetPattern(p string) {}
