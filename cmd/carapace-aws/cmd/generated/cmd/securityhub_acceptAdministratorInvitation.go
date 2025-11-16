package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_acceptAdministratorInvitationCmd = &cobra.Command{
	Use:   "accept-administrator-invitation",
	Short: "We recommend using Organizations instead of Security Hub invitations to manage your member accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_acceptAdministratorInvitationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_acceptAdministratorInvitationCmd).Standalone()

		securityhub_acceptAdministratorInvitationCmd.Flags().String("administrator-id", "", "The account ID of the Security Hub administrator account that sent the invitation.")
		securityhub_acceptAdministratorInvitationCmd.Flags().String("invitation-id", "", "The identifier of the invitation sent from the Security Hub administrator account.")
		securityhub_acceptAdministratorInvitationCmd.MarkFlagRequired("administrator-id")
		securityhub_acceptAdministratorInvitationCmd.MarkFlagRequired("invitation-id")
	})
	securityhubCmd.AddCommand(securityhub_acceptAdministratorInvitationCmd)
}
