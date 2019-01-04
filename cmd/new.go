package cmd

import (
	"os"
	"github.com/dio/voyex/generators/project"
	"github.com/dio/voyex/meta"
	"github.com/gobuffalo/makr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use: "new name",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.WithStack(errors.New("`name` is required"))
		}
		name := args[0]
		if _, err := os.Stat(name); !os.IsNotExist(err) {
			return errors.WithStack(errors.New("cannot create " + name))
		}
		p := newProject(name)
		if err := p.Run(makr.Data{}); err != nil {
			return errors.WithStack(err)
		}
		return nil
	},
}

func newProject(name string) project.Generator {
	return project.Generator{
		Project: meta.New(name),
	}
}

func init() {
	RootCmd.AddCommand(newCmd)
}
