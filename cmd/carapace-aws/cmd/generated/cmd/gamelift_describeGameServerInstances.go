package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeGameServerInstancesCmd = &cobra.Command{
	Use:   "describe-game-server-instances",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nRetrieves status information about the Amazon EC2 instances associated with a Amazon GameLift Servers FleetIQ game server group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeGameServerInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeGameServerInstancesCmd).Standalone()

		gamelift_describeGameServerInstancesCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group.")
		gamelift_describeGameServerInstancesCmd.Flags().String("instance-ids", "", "The Amazon EC2 instance IDs that you want to retrieve status on.")
		gamelift_describeGameServerInstancesCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_describeGameServerInstancesCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_describeGameServerInstancesCmd.MarkFlagRequired("game-server-group-name")
	})
	gameliftCmd.AddCommand(gamelift_describeGameServerInstancesCmd)
}
