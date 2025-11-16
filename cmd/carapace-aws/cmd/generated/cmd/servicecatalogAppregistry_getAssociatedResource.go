package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_getAssociatedResourceCmd = &cobra.Command{
	Use:   "get-associated-resource",
	Short: "Gets the resource associated with the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_getAssociatedResourceCmd).Standalone()

	servicecatalogAppregistry_getAssociatedResourceCmd.Flags().String("application", "", "The name, ID, or ARN of the application.")
	servicecatalogAppregistry_getAssociatedResourceCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	servicecatalogAppregistry_getAssociatedResourceCmd.Flags().String("next-token", "", "A unique pagination token for each page of results.")
	servicecatalogAppregistry_getAssociatedResourceCmd.Flags().String("resource", "", "The name or ID of the resource associated with the application.")
	servicecatalogAppregistry_getAssociatedResourceCmd.Flags().String("resource-tag-status", "", "States whether an application tag is applied, not applied, in the process of being applied, or skipped.")
	servicecatalogAppregistry_getAssociatedResourceCmd.Flags().String("resource-type", "", "The type of resource associated with the application.")
	servicecatalogAppregistry_getAssociatedResourceCmd.MarkFlagRequired("application")
	servicecatalogAppregistry_getAssociatedResourceCmd.MarkFlagRequired("resource")
	servicecatalogAppregistry_getAssociatedResourceCmd.MarkFlagRequired("resource-type")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_getAssociatedResourceCmd)
}
