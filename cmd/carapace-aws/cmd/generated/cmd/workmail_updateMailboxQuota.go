package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updateMailboxQuotaCmd = &cobra.Command{
	Use:   "update-mailbox-quota",
	Short: "Updates a user's current mailbox quota for a specified organization and user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updateMailboxQuotaCmd).Standalone()

	workmail_updateMailboxQuotaCmd.Flags().String("mailbox-quota", "", "The updated mailbox quota, in MB, for the specified user.")
	workmail_updateMailboxQuotaCmd.Flags().String("organization-id", "", "The identifier for the organization that contains the user for whom to update the mailbox quota.")
	workmail_updateMailboxQuotaCmd.Flags().String("user-id", "", "The identifer for the user for whom to update the mailbox quota.")
	workmail_updateMailboxQuotaCmd.MarkFlagRequired("mailbox-quota")
	workmail_updateMailboxQuotaCmd.MarkFlagRequired("organization-id")
	workmail_updateMailboxQuotaCmd.MarkFlagRequired("user-id")
	workmailCmd.AddCommand(workmail_updateMailboxQuotaCmd)
}
