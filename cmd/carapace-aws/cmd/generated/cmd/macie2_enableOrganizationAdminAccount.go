package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_enableOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "enable-organization-admin-account",
	Short: "Designates an account as the delegated Amazon Macie administrator account for an organization in Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_enableOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_enableOrganizationAdminAccountCmd).Standalone()

		macie2_enableOrganizationAdminAccountCmd.Flags().String("admin-account-id", "", "The Amazon Web Services account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.")
		macie2_enableOrganizationAdminAccountCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
		macie2_enableOrganizationAdminAccountCmd.MarkFlagRequired("admin-account-id")
	})
	macie2Cmd.AddCommand(macie2_enableOrganizationAdminAccountCmd)
}
