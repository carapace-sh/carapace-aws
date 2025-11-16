package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearchCmd = &cobra.Command{
	Use:   "cloudsearch",
	Short: "Amazon CloudSearch Configuration Service\n\nYou use the Amazon CloudSearch configuration service to create, configure, and manage search domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearchCmd).Standalone()

	rootCmd.AddCommand(cloudsearchCmd)
}
