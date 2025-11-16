package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Retrieves a list of all of your applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_listApplicationsCmd).Standalone()

		servicecatalogAppregistry_listApplicationsCmd.Flags().String("max-results", "", "The upper bound of the number of results to return (cannot exceed 25).")
		servicecatalogAppregistry_listApplicationsCmd.Flags().String("next-token", "", "The token to use to get the next page of results after a previous API call.")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_listApplicationsCmd)
}
