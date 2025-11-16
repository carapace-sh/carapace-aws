package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_listAttributeGroupsCmd = &cobra.Command{
	Use:   "list-attribute-groups",
	Short: "Lists all attribute groups which you have access to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_listAttributeGroupsCmd).Standalone()

	servicecatalogAppregistry_listAttributeGroupsCmd.Flags().String("max-results", "", "The upper bound of the number of results to return (cannot exceed 25).")
	servicecatalogAppregistry_listAttributeGroupsCmd.Flags().String("next-token", "", "The token to use to get the next page of results after a previous API call.")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_listAttributeGroupsCmd)
}
