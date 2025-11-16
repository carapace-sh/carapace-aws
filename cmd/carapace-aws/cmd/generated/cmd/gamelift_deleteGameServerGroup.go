package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteGameServerGroupCmd = &cobra.Command{
	Use:   "delete-game-server-group",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nTerminates a game server group and permanently deletes the game server group record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteGameServerGroupCmd).Standalone()

	gamelift_deleteGameServerGroupCmd.Flags().String("delete-option", "", "The type of delete to perform.")
	gamelift_deleteGameServerGroupCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group.")
	gamelift_deleteGameServerGroupCmd.MarkFlagRequired("game-server-group-name")
	gameliftCmd.AddCommand(gamelift_deleteGameServerGroupCmd)
}
