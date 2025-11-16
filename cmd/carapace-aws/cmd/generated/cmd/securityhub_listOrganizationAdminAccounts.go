package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listOrganizationAdminAccountsCmd = &cobra.Command{
	Use:   "list-organization-admin-accounts",
	Short: "Lists the Security Hub administrator accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listOrganizationAdminAccountsCmd).Standalone()

	securityhub_listOrganizationAdminAccountsCmd.Flags().String("feature", "", "The feature where the delegated administrator account is listed.")
	securityhub_listOrganizationAdminAccountsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
	securityhub_listOrganizationAdminAccountsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	securityhubCmd.AddCommand(securityhub_listOrganizationAdminAccountsCmd)
}
