package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controlcatalogCmd = &cobra.Command{
	Use:   "controlcatalog",
	Short: "Welcome to the Control Catalog API reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controlcatalogCmd).Standalone()

	rootCmd.AddCommand(controlcatalogCmd)
}
