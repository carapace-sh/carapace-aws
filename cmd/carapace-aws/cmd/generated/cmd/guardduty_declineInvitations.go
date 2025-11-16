package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_declineInvitationsCmd = &cobra.Command{
	Use:   "decline-invitations",
	Short: "Declines invitations sent to the current member account by Amazon Web Services accounts specified by their account IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_declineInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_declineInvitationsCmd).Standalone()

		guardduty_declineInvitationsCmd.Flags().String("account-ids", "", "A list of account IDs of the Amazon Web Services accounts that sent invitations to the current member account that you want to decline invitations from.")
		guardduty_declineInvitationsCmd.MarkFlagRequired("account-ids")
	})
	guarddutyCmd.AddCommand(guardduty_declineInvitationsCmd)
}
