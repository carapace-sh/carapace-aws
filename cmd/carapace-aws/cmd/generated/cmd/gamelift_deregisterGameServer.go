package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deregisterGameServerCmd = &cobra.Command{
	Use:   "deregister-game-server",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nRemoves the game server from a game server group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deregisterGameServerCmd).Standalone()

	gamelift_deregisterGameServerCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group where the game server is running.")
	gamelift_deregisterGameServerCmd.Flags().String("game-server-id", "", "A custom string that uniquely identifies the game server to deregister.")
	gamelift_deregisterGameServerCmd.MarkFlagRequired("game-server-group-name")
	gamelift_deregisterGameServerCmd.MarkFlagRequired("game-server-id")
	gameliftCmd.AddCommand(gamelift_deregisterGameServerCmd)
}
