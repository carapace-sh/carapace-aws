package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_disableOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "disable-organization-admin-account",
	Short: "Removes the existing GuardDuty delegated administrator of the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_disableOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_disableOrganizationAdminAccountCmd).Standalone()

		guardduty_disableOrganizationAdminAccountCmd.Flags().String("admin-account-id", "", "The Amazon Web Services Account ID for the organizations account to be disabled as a GuardDuty delegated administrator.")
		guardduty_disableOrganizationAdminAccountCmd.MarkFlagRequired("admin-account-id")
	})
	guarddutyCmd.AddCommand(guardduty_disableOrganizationAdminAccountCmd)
}
