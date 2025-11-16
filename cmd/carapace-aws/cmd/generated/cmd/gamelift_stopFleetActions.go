package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_stopFleetActionsCmd = &cobra.Command{
	Use:   "stop-fleet-actions",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nSuspends certain types of activity in a fleet location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_stopFleetActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_stopFleetActionsCmd).Standalone()

		gamelift_stopFleetActionsCmd.Flags().String("actions", "", "List of actions to suspend on the fleet.")
		gamelift_stopFleetActionsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to stop actions on.")
		gamelift_stopFleetActionsCmd.Flags().String("location", "", "The fleet location to stop fleet actions for.")
		gamelift_stopFleetActionsCmd.MarkFlagRequired("actions")
		gamelift_stopFleetActionsCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_stopFleetActionsCmd)
}
