package filter

import (
	"github.com/gobuffalo/packr"
)

// Templates ...
var Templates = packr.NewBox("../filter/templates")

// Generator ...
type Generator struct {
	Name string
	Type string
}

// New ...
func New(name string) (Generator, error) {
	g := Generator{
		Name: name,
		Type: "http-decoder",
	}

	return g, g.Validate()
}

// Validate ...
func (g Generator) Validate() error {
	return nil
}
