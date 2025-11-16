package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_listAssociatedAttributeGroupsCmd = &cobra.Command{
	Use:   "list-associated-attribute-groups",
	Short: "Lists all attribute groups that are associated with specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_listAssociatedAttributeGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_listAssociatedAttributeGroupsCmd).Standalone()

		servicecatalogAppregistry_listAssociatedAttributeGroupsCmd.Flags().String("application", "", "The name or ID of the application.")
		servicecatalogAppregistry_listAssociatedAttributeGroupsCmd.Flags().String("max-results", "", "The upper bound of the number of results to return (cannot exceed 25).")
		servicecatalogAppregistry_listAssociatedAttributeGroupsCmd.Flags().String("next-token", "", "The token to use to get the next page of results after a previous API call.")
		servicecatalogAppregistry_listAssociatedAttributeGroupsCmd.MarkFlagRequired("application")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_listAssociatedAttributeGroupsCmd)
}
