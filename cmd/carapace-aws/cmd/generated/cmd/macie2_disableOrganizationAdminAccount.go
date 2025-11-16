package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_disableOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "disable-organization-admin-account",
	Short: "Disables an account as the delegated Amazon Macie administrator account for an organization in Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_disableOrganizationAdminAccountCmd).Standalone()

	macie2_disableOrganizationAdminAccountCmd.Flags().String("admin-account-id", "", "The Amazon Web Services account ID of the delegated Amazon Macie administrator account.")
	macie2_disableOrganizationAdminAccountCmd.MarkFlagRequired("admin-account-id")
	macie2Cmd.AddCommand(macie2_disableOrganizationAdminAccountCmd)
}
