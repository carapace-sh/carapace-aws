package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_inviteMembersCmd = &cobra.Command{
	Use:   "invite-members",
	Short: "We recommend using Organizations instead of Security Hub invitations to manage your member accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_inviteMembersCmd).Standalone()

	securityhub_inviteMembersCmd.Flags().String("account-ids", "", "The list of account IDs of the Amazon Web Services accounts to invite to Security Hub as members.")
	securityhub_inviteMembersCmd.MarkFlagRequired("account-ids")
	securityhubCmd.AddCommand(securityhub_inviteMembersCmd)
}
