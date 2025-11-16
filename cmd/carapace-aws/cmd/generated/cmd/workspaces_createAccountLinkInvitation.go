package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createAccountLinkInvitationCmd = &cobra.Command{
	Use:   "create-account-link-invitation",
	Short: "Creates the account link invitation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createAccountLinkInvitationCmd).Standalone()

	workspaces_createAccountLinkInvitationCmd.Flags().String("client-token", "", "A string of up to 64 ASCII characters that Amazon WorkSpaces uses to ensure idempotent creation.")
	workspaces_createAccountLinkInvitationCmd.Flags().String("target-account-id", "", "The identifier of the target account.")
	workspaces_createAccountLinkInvitationCmd.MarkFlagRequired("target-account-id")
	workspacesCmd.AddCommand(workspaces_createAccountLinkInvitationCmd)
}
