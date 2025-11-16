package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_inviteUsersCmd = &cobra.Command{
	Use:   "invite-users",
	Short: "Sends email to a maximum of 50 users, inviting them to the specified Amazon Chime `Team` account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_inviteUsersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_inviteUsersCmd).Standalone()

		chime_inviteUsersCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_inviteUsersCmd.Flags().String("user-email-list", "", "The user email addresses to which to send the email invitation.")
		chime_inviteUsersCmd.Flags().String("user-type", "", "The user type.")
		chime_inviteUsersCmd.MarkFlagRequired("account-id")
		chime_inviteUsersCmd.MarkFlagRequired("user-email-list")
	})
	chimeCmd.AddCommand(chime_inviteUsersCmd)
}
