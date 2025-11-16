package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_startFleetActionsCmd = &cobra.Command{
	Use:   "start-fleet-actions",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nResumes certain types of activity on fleet instances that were suspended with [StopFleetActions](https://docs.aws.amazon.com/gamelift/latest/apireference/API_StopFleetActions.html). For multi-location fleets, fleet actions are managed separately for each location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_startFleetActionsCmd).Standalone()

	gamelift_startFleetActionsCmd.Flags().String("actions", "", "List of actions to restart on the fleet.")
	gamelift_startFleetActionsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to restart actions on.")
	gamelift_startFleetActionsCmd.Flags().String("location", "", "The fleet location to restart fleet actions for.")
	gamelift_startFleetActionsCmd.MarkFlagRequired("actions")
	gamelift_startFleetActionsCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_startFleetActionsCmd)
}
