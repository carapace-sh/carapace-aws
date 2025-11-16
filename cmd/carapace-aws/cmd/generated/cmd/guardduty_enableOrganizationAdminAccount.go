package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_enableOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "enable-organization-admin-account",
	Short: "Designates an Amazon Web Services account within the organization as your GuardDuty delegated administrator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_enableOrganizationAdminAccountCmd).Standalone()

	guardduty_enableOrganizationAdminAccountCmd.Flags().String("admin-account-id", "", "The Amazon Web Services account ID for the organization account to be enabled as a GuardDuty delegated administrator.")
	guardduty_enableOrganizationAdminAccountCmd.MarkFlagRequired("admin-account-id")
	guarddutyCmd.AddCommand(guardduty_enableOrganizationAdminAccountCmd)
}
