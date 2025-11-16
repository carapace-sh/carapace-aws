package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteMailboxPermissionsCmd = &cobra.Command{
	Use:   "delete-mailbox-permissions",
	Short: "Deletes permissions granted to a member (user or group).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteMailboxPermissionsCmd).Standalone()

	workmail_deleteMailboxPermissionsCmd.Flags().String("entity-id", "", "The identifier of the entity that owns the mailbox.")
	workmail_deleteMailboxPermissionsCmd.Flags().String("grantee-id", "", "The identifier of the entity for which to delete granted permissions.")
	workmail_deleteMailboxPermissionsCmd.Flags().String("organization-id", "", "The identifier of the organization under which the member (user or group) exists.")
	workmail_deleteMailboxPermissionsCmd.MarkFlagRequired("entity-id")
	workmail_deleteMailboxPermissionsCmd.MarkFlagRequired("grantee-id")
	workmail_deleteMailboxPermissionsCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deleteMailboxPermissionsCmd)
}
