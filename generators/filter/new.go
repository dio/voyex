package filter

import (
	"path/filepath"
	"strings"

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
		// LOL!
		writePath := strings.Replace(f.WritePath, "~impl", a.Name, 1)
		writePath = strings.Replace(writePath, "~filter", a.Name+"_filter", 1)
		writePath = strings.Replace(writePath, "http/", "", 1)
		g.Add(makr.NewFile(writePath, f.Body))
	}
	path, _ := filepath.Abs(a.Root)
	return g.Run(path, data)
}
