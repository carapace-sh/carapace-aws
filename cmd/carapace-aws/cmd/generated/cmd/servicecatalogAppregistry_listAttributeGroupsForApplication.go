package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_listAttributeGroupsForApplicationCmd = &cobra.Command{
	Use:   "list-attribute-groups-for-application",
	Short: "Lists the details of all attribute groups associated with a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_listAttributeGroupsForApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_listAttributeGroupsForApplicationCmd).Standalone()

		servicecatalogAppregistry_listAttributeGroupsForApplicationCmd.Flags().String("application", "", "The name or ID of the application.")
		servicecatalogAppregistry_listAttributeGroupsForApplicationCmd.Flags().String("max-results", "", "The upper bound of the number of results to return.")
		servicecatalogAppregistry_listAttributeGroupsForApplicationCmd.Flags().String("next-token", "", "This token retrieves the next page of results after a previous API call.")
		servicecatalogAppregistry_listAttributeGroupsForApplicationCmd.MarkFlagRequired("application")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_listAttributeGroupsForApplicationCmd)
}
