package generate

import (
	"github.com/dio/voyex/generators/filter"
	"github.com/gobuffalo/makr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// FilterCmd ...
var FilterCmd = &cobra.Command{
	Use: "filter name",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, err := cmd.Flags().GetString("root")
		if err != nil {
			return errors.WithStack(errors.New("root is required"))
		}
		f, err := newFilter(args[0], root)
		if err != nil {
			return errors.WithStack(err)
		}
		if err := f.Run(makr.Data{}); err != nil {
			return errors.WithStack(err)
		}
		return nil
	},
}

func init() {
	FilterCmd.Flags().StringP("root", "r", "", "root")
}

func newFilter(name, root string) (filter.Generator, error) {
	return filter.New(name, root)
}


