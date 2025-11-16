package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_putMailboxPermissionsCmd = &cobra.Command{
	Use:   "put-mailbox-permissions",
	Short: "Sets permissions for a user, group, or resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_putMailboxPermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_putMailboxPermissionsCmd).Standalone()

		workmail_putMailboxPermissionsCmd.Flags().String("entity-id", "", "The identifier of the user or resource for which to update mailbox permissions.")
		workmail_putMailboxPermissionsCmd.Flags().String("grantee-id", "", "The identifier of the user, group, or resource to which to grant the permissions.")
		workmail_putMailboxPermissionsCmd.Flags().String("organization-id", "", "The identifier of the organization under which the user, group, or resource exists.")
		workmail_putMailboxPermissionsCmd.Flags().String("permission-values", "", "The permissions granted to the grantee.")
		workmail_putMailboxPermissionsCmd.MarkFlagRequired("entity-id")
		workmail_putMailboxPermissionsCmd.MarkFlagRequired("grantee-id")
		workmail_putMailboxPermissionsCmd.MarkFlagRequired("organization-id")
		workmail_putMailboxPermissionsCmd.MarkFlagRequired("permission-values")
	})
	workmailCmd.AddCommand(workmail_putMailboxPermissionsCmd)
}
