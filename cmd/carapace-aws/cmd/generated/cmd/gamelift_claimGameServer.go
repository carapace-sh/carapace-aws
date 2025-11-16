package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_claimGameServerCmd = &cobra.Command{
	Use:   "claim-game-server",
	Short: "**This API works with the following fleet types:** EC2 (FleetIQ)\n\nLocates an available game server and temporarily reserves it to host gameplay and players.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_claimGameServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_claimGameServerCmd).Standalone()

		gamelift_claimGameServerCmd.Flags().String("filter-option", "", "Object that restricts how a claimed game server is chosen.")
		gamelift_claimGameServerCmd.Flags().String("game-server-data", "", "A set of custom game server properties, formatted as a single string value.")
		gamelift_claimGameServerCmd.Flags().String("game-server-group-name", "", "A unique identifier for the game server group where the game server is running.")
		gamelift_claimGameServerCmd.Flags().String("game-server-id", "", "A custom string that uniquely identifies the game server to claim.")
		gamelift_claimGameServerCmd.MarkFlagRequired("game-server-group-name")
	})
	gameliftCmd.AddCommand(gamelift_claimGameServerCmd)
}
