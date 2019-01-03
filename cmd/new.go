package cmd

import (
	"github.com/dio/voyex/generators/project"
	"github.com/dio/voyex/meta"
	"github.com/gobuffalo/makr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use: "new [name]",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := newProject()
		if err := p.Run(makr.Data{}); err != nil {
			return errors.WithStack(err)
		}
		return nil
	},
}

func newProject() project.Generator {
	return project.Generator{
		Project: meta.New("ok"),
	}
}

func init() {
	RootCmd.AddCommand(newCmd)
}
