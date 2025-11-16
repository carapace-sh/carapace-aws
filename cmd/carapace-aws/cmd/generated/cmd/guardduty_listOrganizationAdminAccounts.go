package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listOrganizationAdminAccountsCmd = &cobra.Command{
	Use:   "list-organization-admin-accounts",
	Short: "Lists the accounts designated as GuardDuty delegated administrators.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listOrganizationAdminAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_listOrganizationAdminAccountsCmd).Standalone()

		guardduty_listOrganizationAdminAccountsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		guardduty_listOrganizationAdminAccountsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	})
	guarddutyCmd.AddCommand(guardduty_listOrganizationAdminAccountsCmd)
}
