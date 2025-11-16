package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_disableProactiveEngagementCmd = &cobra.Command{
	Use:   "disable-proactive-engagement",
	Short: "Removes authorization from the Shield Response Team (SRT) to notify contacts about escalations to the SRT and to initiate proactive customer support.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_disableProactiveEngagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_disableProactiveEngagementCmd).Standalone()

	})
	shieldCmd.AddCommand(shield_disableProactiveEngagementCmd)
}
