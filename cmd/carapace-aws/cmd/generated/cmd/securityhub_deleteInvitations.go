package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteInvitationsCmd = &cobra.Command{
	Use:   "delete-invitations",
	Short: "We recommend using Organizations instead of Security Hub invitations to manage your member accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_deleteInvitationsCmd).Standalone()

		securityhub_deleteInvitationsCmd.Flags().String("account-ids", "", "The list of member account IDs that received the invitations you want to delete.")
		securityhub_deleteInvitationsCmd.MarkFlagRequired("account-ids")
	})
	securityhubCmd.AddCommand(securityhub_deleteInvitationsCmd)
}
