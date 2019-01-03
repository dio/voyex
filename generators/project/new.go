package project

import (
	"github.com/gobuffalo/buffalo/generators"
	"github.com/gobuffalo/makr"
	"github.com/pkg/errors"
)

// Run ...
func (a Generator) Run(data makr.Data) error {
	g := makr.New()
	data["opts"] = a
	files, err := generators.FindByBox(Templates)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, f := range files {
		g.Add(makr.NewFile(f.WritePath, f.Body))
	}
	return g.Run("/tmp/ok", data)
}
