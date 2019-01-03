package cmd

import (
	"github.com/dio/voyex/cmd/generate"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
}

func init() {
	generateCmd.AddCommand(generate.FilterCmd)
	RootCmd.AddCommand(generateCmd)
}
