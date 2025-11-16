package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deleteAccountLinkInvitationCmd = &cobra.Command{
	Use:   "delete-account-link-invitation",
	Short: "Deletes the account link invitation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deleteAccountLinkInvitationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_deleteAccountLinkInvitationCmd).Standalone()

		workspaces_deleteAccountLinkInvitationCmd.Flags().String("client-token", "", "A string of up to 64 ASCII characters that Amazon WorkSpaces uses to ensure idempotent creation.")
		workspaces_deleteAccountLinkInvitationCmd.Flags().String("link-id", "", "The identifier of the account link.")
		workspaces_deleteAccountLinkInvitationCmd.MarkFlagRequired("link-id")
	})
	workspacesCmd.AddCommand(workspaces_deleteAccountLinkInvitationCmd)
}
