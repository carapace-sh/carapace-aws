package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listOrganizationAdminAccountsCmd = &cobra.Command{
	Use:   "list-organization-admin-accounts",
	Short: "Retrieves information about the delegated Amazon Macie administrator account for an organization in Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listOrganizationAdminAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_listOrganizationAdminAccountsCmd).Standalone()

		macie2_listOrganizationAdminAccountsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of a paginated response.")
		macie2_listOrganizationAdminAccountsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	})
	macie2Cmd.AddCommand(macie2_listOrganizationAdminAccountsCmd)
}
