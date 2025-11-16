package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_syncResourceCmd = &cobra.Command{
	Use:   "sync-resource",
	Short: "Syncs the resource with current AppRegistry records.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_syncResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_syncResourceCmd).Standalone()

		servicecatalogAppregistry_syncResourceCmd.Flags().String("resource", "", "An entity you can work with and specify with a name or ID.")
		servicecatalogAppregistry_syncResourceCmd.Flags().String("resource-type", "", "The type of resource of which the application will be associated.")
		servicecatalogAppregistry_syncResourceCmd.MarkFlagRequired("resource")
		servicecatalogAppregistry_syncResourceCmd.MarkFlagRequired("resource-type")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_syncResourceCmd)
}
