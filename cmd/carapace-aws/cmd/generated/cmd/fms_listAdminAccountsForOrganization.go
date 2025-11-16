package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listAdminAccountsForOrganizationCmd = &cobra.Command{
	Use:   "list-admin-accounts-for-organization",
	Short: "Returns a `AdminAccounts` object that lists the Firewall Manager administrators within the organization that are onboarded to Firewall Manager by [AssociateAdminAccount]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listAdminAccountsForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_listAdminAccountsForOrganizationCmd).Standalone()

		fms_listAdminAccountsForOrganizationCmd.Flags().String("max-results", "", "The maximum number of objects that you want Firewall Manager to return for this request.")
		fms_listAdminAccountsForOrganizationCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Firewall Manager returns a `NextToken` value in the response.")
	})
	fmsCmd.AddCommand(fms_listAdminAccountsForOrganizationCmd)
}
