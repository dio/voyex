package project

import (
	"path/filepath"

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

	root, err := filepath.Abs(a.Project.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	return g.Run(root, data)
}
