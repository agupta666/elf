package actions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path"

	"github.com/agupta666/elf/utils"
	uuid "github.com/satori/go.uuid"
)

// DumpAction dumps all incoming requests to a file in the supplied folder
type DumpAction struct {
	Root string
}

func (da *DumpAction) filePath() string {
	fName := fmt.Sprintf("request-%s.dump", uuid.NewV4().String())

	return path.Join(da.Root, fName)
}

func (da *DumpAction) String() string {
	return fmt.Sprintf("dump[Path=%s]", da.Root)
}

// Exec executes a dir action
func (da *DumpAction) Exec(w http.ResponseWriter, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)

	if err != nil {
		return err
	}

	fPath := da.filePath()

	err = ioutil.WriteFile(fPath, data, os.ModePerm)

	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Finished dumping request to %s\n", fPath)
	return nil
}

// SetPattern sets the matched pattern in the action
func (da *DumpAction) SetPattern(p string) {}

// NewDumpAction creates a new dump action
func NewDumpAction(root string) (*DumpAction, error) {
	ok, _ := utils.FileExists(root)
	if !ok {
		return nil, errors.New("specified folder does not exist")
	}
	return &DumpAction{Root: root}, nil
}

// NewDefaultDumpAction creates a dump action with current wirk dir as root dir
func NewDefaultDumpAction() (*DumpAction, error) {
	workDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}
	return NewDumpAction(workDir)
}

// NewDumpActionFromExpr creates a new DumpAction from an expression
func NewDumpActionFromExpr(exp string) (*DumpAction, error) {

	ac, err := parseExpr(exp, "dump", func(args []string) (Action, error) {
		switch len(args) {
		case 0:
			return NewDefaultDumpAction()
		case 1:
			return NewDumpAction(args[0])
		default:
			return nil, errors.New("invalid expression")
		}
	})

	return ac.(*DumpAction), err
}
