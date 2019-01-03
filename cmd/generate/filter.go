package generate

import (
	"github.com/spf13/cobra"
)

// FilterCmd ...
var FilterCmd = &cobra.Command{
	Use: "filter [name]",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
