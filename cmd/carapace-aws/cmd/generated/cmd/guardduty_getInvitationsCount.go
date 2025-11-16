package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getInvitationsCountCmd = &cobra.Command{
	Use:   "get-invitations-count",
	Short: "Returns the count of all GuardDuty membership invitations that were sent to the current member account except the currently accepted invitation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getInvitationsCountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getInvitationsCountCmd).Standalone()

	})
	guarddutyCmd.AddCommand(guardduty_getInvitationsCountCmd)
}
