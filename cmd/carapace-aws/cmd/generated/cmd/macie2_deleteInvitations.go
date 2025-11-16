package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_deleteInvitationsCmd = &cobra.Command{
	Use:   "delete-invitations",
	Short: "Deletes Amazon Macie membership invitations that were received from specific accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_deleteInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_deleteInvitationsCmd).Standalone()

		macie2_deleteInvitationsCmd.Flags().String("account-ids", "", "An array that lists Amazon Web Services account IDs, one for each account that sent an invitation to delete.")
		macie2_deleteInvitationsCmd.MarkFlagRequired("account-ids")
	})
	macie2Cmd.AddCommand(macie2_deleteInvitationsCmd)
}
