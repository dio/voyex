package generate

import (
	"github.com/dio/voyex/generators/filter"
	"github.com/gobuffalo/makr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// FilterCmd ...
var FilterCmd = &cobra.Command{
	Use: "filter [name]",
	RunE: func(cmd *cobra.Command, args []string) error {
		f := newFilter()
		if err := f.Run(makr.Data{}); err != nil {
			return errors.WithStack(err)
		}
		return nil
	},
}

func newFilter() filter.Generator {
	return filter.Generator{}
}
