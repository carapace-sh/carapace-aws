package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogCmd = &cobra.Command{
	Use:   "servicecatalog",
	Short: "Service Catalog\n\n[Service Catalog](http://aws.amazon.com/servicecatalog) enables organizations to create and manage catalogs of IT services that are approved for Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogCmd).Standalone()

	})
	rootCmd.AddCommand(servicecatalogCmd)
}
