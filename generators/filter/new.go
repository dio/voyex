package filter

import (
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
		writePath := strings.Replace(f.WritePath, "~impl", "basic_auth", 1)
		writePath = strings.Replace(writePath, "~filter", "basic_auth_filter", 1)
		writePath = strings.Replace(writePath, "http/", "", 1)
		g.Add(makr.NewFile(writePath, f.Body))
	}
	return g.Run("/tmp/ok", data)
}
