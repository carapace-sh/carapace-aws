package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_createInvitationsCmd = &cobra.Command{
	Use:   "create-invitations",
	Short: "Sends an Amazon Macie membership invitation to one or more accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_createInvitationsCmd).Standalone()

	macie2_createInvitationsCmd.Flags().String("account-ids", "", "An array that lists Amazon Web Services account IDs, one for each account to send the invitation to.")
	macie2_createInvitationsCmd.Flags().String("disable-email-notification", "", "Specifies whether to send the invitation as an email message.")
	macie2_createInvitationsCmd.Flags().String("message", "", "Custom text to include in the email message that contains the invitation.")
	macie2_createInvitationsCmd.MarkFlagRequired("account-ids")
	macie2Cmd.AddCommand(macie2_createInvitationsCmd)
}
