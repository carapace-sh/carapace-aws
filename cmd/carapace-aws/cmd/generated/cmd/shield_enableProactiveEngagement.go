package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_enableProactiveEngagementCmd = &cobra.Command{
	Use:   "enable-proactive-engagement",
	Short: "Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_enableProactiveEngagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_enableProactiveEngagementCmd).Standalone()

	})
	shieldCmd.AddCommand(shield_enableProactiveEngagementCmd)
}
