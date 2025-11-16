package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistryCmd = &cobra.Command{
	Use:   "servicecatalog-appregistry",
	Short: "Amazon Web Services Service Catalog AppRegistry enables organizations to understand the application context of their Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistryCmd).Standalone()

	})
	rootCmd.AddCommand(servicecatalogAppregistryCmd)
}
