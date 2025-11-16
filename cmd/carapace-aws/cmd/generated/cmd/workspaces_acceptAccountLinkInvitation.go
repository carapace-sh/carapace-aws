package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_acceptAccountLinkInvitationCmd = &cobra.Command{
	Use:   "accept-account-link-invitation",
	Short: "Accepts the account link invitation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_acceptAccountLinkInvitationCmd).Standalone()

	workspaces_acceptAccountLinkInvitationCmd.Flags().String("client-token", "", "A string of up to 64 ASCII characters that Amazon WorkSpaces uses to ensure idempotent creation.")
	workspaces_acceptAccountLinkInvitationCmd.Flags().String("link-id", "", "The identifier of the account link.")
	workspaces_acceptAccountLinkInvitationCmd.MarkFlagRequired("link-id")
	workspacesCmd.AddCommand(workspaces_acceptAccountLinkInvitationCmd)
}
