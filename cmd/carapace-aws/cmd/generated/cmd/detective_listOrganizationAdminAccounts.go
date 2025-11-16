package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_listOrganizationAdminAccountsCmd = &cobra.Command{
	Use:   "list-organization-admin-accounts",
	Short: "Returns information about the Detective administrator account for an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_listOrganizationAdminAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_listOrganizationAdminAccountsCmd).Standalone()

		detective_listOrganizationAdminAccountsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		detective_listOrganizationAdminAccountsCmd.Flags().String("next-token", "", "For requests to get the next page of results, the pagination token that was returned with the previous set of results.")
	})
	detectiveCmd.AddCommand(detective_listOrganizationAdminAccountsCmd)
}
