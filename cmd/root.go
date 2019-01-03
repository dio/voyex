package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is ...
var RootCmd = &cobra.Command{
	Use:   "voyex",
	Short: "Help you to extend envoy",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// Execute ...
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
