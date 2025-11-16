package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_enableOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "enable-organization-admin-account",
	Short: "Designates the Security Hub administrator account for an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_enableOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_enableOrganizationAdminAccountCmd).Standalone()

		securityhub_enableOrganizationAdminAccountCmd.Flags().String("admin-account-id", "", "The Amazon Web Services account identifier of the account to designate as the Security Hub administrator account.")
		securityhub_enableOrganizationAdminAccountCmd.Flags().String("feature", "", "The feature for which the delegated admin account is enabled.")
		securityhub_enableOrganizationAdminAccountCmd.MarkFlagRequired("admin-account-id")
	})
	securityhubCmd.AddCommand(securityhub_enableOrganizationAdminAccountCmd)
}
