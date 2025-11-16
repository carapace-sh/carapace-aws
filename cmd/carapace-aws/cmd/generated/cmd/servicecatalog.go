package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogCmd = &cobra.Command{
	Use:   "servicecatalog",
	Short: "Service Catalog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogCmd).Standalone()

	})
	rootCmd.AddCommand(servicecatalogCmd)
}
