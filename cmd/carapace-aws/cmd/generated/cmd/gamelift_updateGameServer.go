package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateGameServerCmd = &cobra.Command{
	Use:   "update-game-server",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nUpdates information about a registered game server to help Amazon GameLift Servers FleetIQ track game server availability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateGameServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_updateGameServerCmd).Standalone()

		gamelift_updateGameServerCmd.Flags().String("game-server-data", "", "A set of custom game server properties, formatted as a single string value.")
		gamelift_updateGameServerCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group where the game server is running.")
		gamelift_updateGameServerCmd.Flags().String("game-server-id", "", "A custom string that uniquely identifies the game server to update.")
		gamelift_updateGameServerCmd.Flags().String("health-check", "", "Indicates health status of the game server.")
		gamelift_updateGameServerCmd.Flags().String("utilization-status", "", "Indicates if the game server is available or is currently hosting gameplay.")
		gamelift_updateGameServerCmd.MarkFlagRequired("game-server-group-name")
		gamelift_updateGameServerCmd.MarkFlagRequired("game-server-id")
	})
	gameliftCmd.AddCommand(gamelift_updateGameServerCmd)
}
