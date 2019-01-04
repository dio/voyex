package filter

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
)

// Templates ...
var Templates = packr.NewBox("../filter/templates")

// Generator ...
type Generator struct {
	Name string
	Type string
	Root string
}

// New ...
func New(name string, root string) (Generator, error) {
	g := Generator{
		Name: name,
		Type: "http-decoder",
		Root: root,
	}

	return g, g.Validate()
}

// Validate ...
func (g Generator) Validate() error {
	path, err := filepath.Abs(g.Root)
	if err != nil {
		return errors.New("root is invalid")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("root is invalid")
	}
	if _, err := os.Stat(filepath.Join(path, "WORKSPACE")); os.IsNotExist(err) {
		return errors.New("root is invalid")
	}
	return nil
}
