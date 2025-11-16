package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_disableOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "disable-organization-admin-account",
	Short: "Disables a Security Hub administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_disableOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_disableOrganizationAdminAccountCmd).Standalone()

		securityhub_disableOrganizationAdminAccountCmd.Flags().String("admin-account-id", "", "The Amazon Web Services account identifier of the Security Hub administrator account.")
		securityhub_disableOrganizationAdminAccountCmd.Flags().String("feature", "", "The feature for which the delegated admin account is disabled.")
		securityhub_disableOrganizationAdminAccountCmd.MarkFlagRequired("admin-account-id")
	})
	securityhubCmd.AddCommand(securityhub_disableOrganizationAdminAccountCmd)
}
