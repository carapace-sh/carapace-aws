package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_listAssociatedResourcesCmd = &cobra.Command{
	Use:   "list-associated-resources",
	Short: "Lists all of the resources that are associated with the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_listAssociatedResourcesCmd).Standalone()

	servicecatalogAppregistry_listAssociatedResourcesCmd.Flags().String("application", "", "The name, ID, or ARN of the application.")
	servicecatalogAppregistry_listAssociatedResourcesCmd.Flags().String("max-results", "", "The upper bound of the number of results to return (cannot exceed 25).")
	servicecatalogAppregistry_listAssociatedResourcesCmd.Flags().String("next-token", "", "The token to use to get the next page of results after a previous API call.")
	servicecatalogAppregistry_listAssociatedResourcesCmd.MarkFlagRequired("application")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_listAssociatedResourcesCmd)
}
