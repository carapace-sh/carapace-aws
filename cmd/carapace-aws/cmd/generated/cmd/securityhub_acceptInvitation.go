package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_acceptInvitationCmd = &cobra.Command{
	Use:   "accept-invitation",
	Short: "This method is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_acceptInvitationCmd).Standalone()

	securityhub_acceptInvitationCmd.Flags().String("invitation-id", "", "The identifier of the invitation sent from the Security Hub administrator account.")
	securityhub_acceptInvitationCmd.Flags().String("master-id", "", "The account ID of the Security Hub administrator account that sent the invitation.")
	securityhub_acceptInvitationCmd.MarkFlagRequired("invitation-id")
	securityhub_acceptInvitationCmd.MarkFlagRequired("master-id")
	securityhubCmd.AddCommand(securityhub_acceptInvitationCmd)
}
