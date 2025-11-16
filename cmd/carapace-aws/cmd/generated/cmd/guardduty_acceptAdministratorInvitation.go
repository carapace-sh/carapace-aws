package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_acceptAdministratorInvitationCmd = &cobra.Command{
	Use:   "accept-administrator-invitation",
	Short: "Accepts the invitation to be a member account and get monitored by a GuardDuty administrator account that sent the invitation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_acceptAdministratorInvitationCmd).Standalone()

	guardduty_acceptAdministratorInvitationCmd.Flags().String("administrator-id", "", "The account ID of the GuardDuty administrator account whose invitation you're accepting.")
	guardduty_acceptAdministratorInvitationCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty member account.")
	guardduty_acceptAdministratorInvitationCmd.Flags().String("invitation-id", "", "The value that is used to validate the administrator account to the member account.")
	guardduty_acceptAdministratorInvitationCmd.MarkFlagRequired("administrator-id")
	guardduty_acceptAdministratorInvitationCmd.MarkFlagRequired("detector-id")
	guardduty_acceptAdministratorInvitationCmd.MarkFlagRequired("invitation-id")
	guarddutyCmd.AddCommand(guardduty_acceptAdministratorInvitationCmd)
}
