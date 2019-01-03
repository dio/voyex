package project

import (
	"github.com/dio/voyex/meta"
	"github.com/gobuffalo/packr"
)

// Templates ...
var Templates = packr.NewBox("../project/templates")

// Generator ...
type Generator struct {
	meta.Project
}

// New ...
func New(name string) (Generator, error) {
	g := Generator{
		Project: meta.New(name),
	}

	return g, g.Validate()
}

// Validate ...
func (g Generator) Validate() error {
	return nil
}
