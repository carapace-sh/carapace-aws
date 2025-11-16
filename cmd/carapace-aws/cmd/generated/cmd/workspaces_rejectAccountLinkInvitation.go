package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_rejectAccountLinkInvitationCmd = &cobra.Command{
	Use:   "reject-account-link-invitation",
	Short: "Rejects the account link invitation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_rejectAccountLinkInvitationCmd).Standalone()

	workspaces_rejectAccountLinkInvitationCmd.Flags().String("client-token", "", "The client token of the account link invitation to reject.")
	workspaces_rejectAccountLinkInvitationCmd.Flags().String("link-id", "", "The identifier of the account link")
	workspaces_rejectAccountLinkInvitationCmd.MarkFlagRequired("link-id")
	workspacesCmd.AddCommand(workspaces_rejectAccountLinkInvitationCmd)
}
