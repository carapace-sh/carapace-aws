package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_acceptInvitationCmd = &cobra.Command{
	Use:   "accept-invitation",
	Short: "Accepts the invitation to be monitored by a GuardDuty administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_acceptInvitationCmd).Standalone()

	guardduty_acceptInvitationCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty member account.")
	guardduty_acceptInvitationCmd.Flags().String("invitation-id", "", "The value that is used to validate the administrator account to the member account.")
	guardduty_acceptInvitationCmd.Flags().String("master-id", "", "The account ID of the GuardDuty administrator account whose invitation you're accepting.")
	guardduty_acceptInvitationCmd.MarkFlagRequired("detector-id")
	guardduty_acceptInvitationCmd.MarkFlagRequired("invitation-id")
	guardduty_acceptInvitationCmd.MarkFlagRequired("master-id")
	guarddutyCmd.AddCommand(guardduty_acceptInvitationCmd)
}
