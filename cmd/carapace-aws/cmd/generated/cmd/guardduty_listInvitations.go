package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listInvitationsCmd = &cobra.Command{
	Use:   "list-invitations",
	Short: "Lists all GuardDuty membership invitations that were sent to the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listInvitationsCmd).Standalone()

	guardduty_listInvitationsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items that you want in the response.")
	guardduty_listInvitationsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	guarddutyCmd.AddCommand(guardduty_listInvitationsCmd)
}
