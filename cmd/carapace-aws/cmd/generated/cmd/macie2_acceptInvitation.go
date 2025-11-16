package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_acceptInvitationCmd = &cobra.Command{
	Use:   "accept-invitation",
	Short: "Accepts an Amazon Macie membership invitation that was received from a specific account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_acceptInvitationCmd).Standalone()

	macie2_acceptInvitationCmd.Flags().String("administrator-account-id", "", "The Amazon Web Services account ID for the account that sent the invitation.")
	macie2_acceptInvitationCmd.Flags().String("invitation-id", "", "The unique identifier for the invitation to accept.")
	macie2_acceptInvitationCmd.Flags().String("master-account", "", "(Deprecated) The Amazon Web Services account ID for the account that sent the invitation.")
	macie2_acceptInvitationCmd.MarkFlagRequired("invitation-id")
	macie2Cmd.AddCommand(macie2_acceptInvitationCmd)
}
