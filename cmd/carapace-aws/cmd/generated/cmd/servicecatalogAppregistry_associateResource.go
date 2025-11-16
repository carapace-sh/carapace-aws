package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_associateResourceCmd = &cobra.Command{
	Use:   "associate-resource",
	Short: "Associates a resource with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_associateResourceCmd).Standalone()

	servicecatalogAppregistry_associateResourceCmd.Flags().String("application", "", "The name, ID, or ARN of the application.")
	servicecatalogAppregistry_associateResourceCmd.Flags().String("options", "", "Determines whether an application tag is applied or skipped.")
	servicecatalogAppregistry_associateResourceCmd.Flags().String("resource", "", "The name or ID of the resource of which the application will be associated.")
	servicecatalogAppregistry_associateResourceCmd.Flags().String("resource-type", "", "The type of resource of which the application will be associated.")
	servicecatalogAppregistry_associateResourceCmd.MarkFlagRequired("application")
	servicecatalogAppregistry_associateResourceCmd.MarkFlagRequired("resource")
	servicecatalogAppregistry_associateResourceCmd.MarkFlagRequired("resource-type")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_associateResourceCmd)
}
