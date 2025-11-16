package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deleteInvitationsCmd = &cobra.Command{
	Use:   "delete-invitations",
	Short: "Deletes invitations sent to the current member account by Amazon Web Services accounts specified by their account IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deleteInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_deleteInvitationsCmd).Standalone()

		guardduty_deleteInvitationsCmd.Flags().String("account-ids", "", "A list of account IDs of the Amazon Web Services accounts that sent invitations to the current member account that you want to delete invitations from.")
		guardduty_deleteInvitationsCmd.MarkFlagRequired("account-ids")
	})
	guarddutyCmd.AddCommand(guardduty_deleteInvitationsCmd)
}
