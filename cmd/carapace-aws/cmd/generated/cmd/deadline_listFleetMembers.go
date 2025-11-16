package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listFleetMembersCmd = &cobra.Command{
	Use:   "list-fleet-members",
	Short: "Lists fleet members.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listFleetMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listFleetMembersCmd).Standalone()

		deadline_listFleetMembersCmd.Flags().String("farm-id", "", "The farm ID of the fleet.")
		deadline_listFleetMembersCmd.Flags().String("fleet-id", "", "The fleet ID to include on the list.")
		deadline_listFleetMembersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listFleetMembersCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listFleetMembersCmd.MarkFlagRequired("farm-id")
		deadline_listFleetMembersCmd.MarkFlagRequired("fleet-id")
	})
	deadlineCmd.AddCommand(deadline_listFleetMembersCmd)
}
