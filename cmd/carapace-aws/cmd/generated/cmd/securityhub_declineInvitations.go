package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_declineInvitationsCmd = &cobra.Command{
	Use:   "decline-invitations",
	Short: "We recommend using Organizations instead of Security Hub invitations to manage your member accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_declineInvitationsCmd).Standalone()

	securityhub_declineInvitationsCmd.Flags().String("account-ids", "", "The list of prospective member account IDs for which to decline an invitation.")
	securityhub_declineInvitationsCmd.MarkFlagRequired("account-ids")
	securityhubCmd.AddCommand(securityhub_declineInvitationsCmd)
}
