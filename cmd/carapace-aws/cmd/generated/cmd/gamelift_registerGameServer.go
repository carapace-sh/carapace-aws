package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_registerGameServerCmd = &cobra.Command{
	Use:   "register-game-server",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nCreates a new game server resource and notifies Amazon GameLift Servers FleetIQ that the game server is ready to host gameplay and players.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_registerGameServerCmd).Standalone()

	gamelift_registerGameServerCmd.Flags().String("connection-info", "", "Information that is needed to make inbound client connections to the game server.")
	gamelift_registerGameServerCmd.Flags().String("game-server-data", "", "A set of custom game server properties, formatted as a single string value.")
	gamelift_registerGameServerCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group where the game server is running.")
	gamelift_registerGameServerCmd.Flags().String("game-server-id", "", "A custom string that uniquely identifies the game server to register.")
	gamelift_registerGameServerCmd.Flags().String("instance-id", "", "The unique identifier for the instance where the game server is running.")
	gamelift_registerGameServerCmd.MarkFlagRequired("game-server-group-name")
	gamelift_registerGameServerCmd.MarkFlagRequired("game-server-id")
	gamelift_registerGameServerCmd.MarkFlagRequired("instance-id")
	gameliftCmd.AddCommand(gamelift_registerGameServerCmd)
}
