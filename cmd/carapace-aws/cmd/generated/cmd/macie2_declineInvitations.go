package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_declineInvitationsCmd = &cobra.Command{
	Use:   "decline-invitations",
	Short: "Declines Amazon Macie membership invitations that were received from specific accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_declineInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_declineInvitationsCmd).Standalone()

		macie2_declineInvitationsCmd.Flags().String("account-ids", "", "An array that lists Amazon Web Services account IDs, one for each account that sent an invitation to decline.")
		macie2_declineInvitationsCmd.MarkFlagRequired("account-ids")
	})
	macie2Cmd.AddCommand(macie2_declineInvitationsCmd)
}
