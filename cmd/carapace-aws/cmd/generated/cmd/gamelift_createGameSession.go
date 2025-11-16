package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createGameSessionCmd = &cobra.Command{
	Use:   "create-game-session",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nCreates a multiplayer game session for players in a specific fleet location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createGameSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createGameSessionCmd).Standalone()

		gamelift_createGameSessionCmd.Flags().String("alias-id", "", "A unique identifier for the alias associated with the fleet to create a game session in.")
		gamelift_createGameSessionCmd.Flags().String("creator-id", "", "A unique identifier for a player or entity creating the game session.")
		gamelift_createGameSessionCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to create a game session in.")
		gamelift_createGameSessionCmd.Flags().String("game-properties", "", "A set of key-value pairs that can store custom data in a game session.")
		gamelift_createGameSessionCmd.Flags().String("game-session-data", "", "A set of custom game session properties, formatted as a single string value.")
		gamelift_createGameSessionCmd.Flags().String("game-session-id", "", "*This parameter is deprecated.")
		gamelift_createGameSessionCmd.Flags().String("idempotency-token", "", "Custom string that uniquely identifies the new game session request.")
		gamelift_createGameSessionCmd.Flags().String("location", "", "A fleet's remote location to place the new game session in.")
		gamelift_createGameSessionCmd.Flags().String("maximum-player-session-count", "", "The maximum number of players that can be connected simultaneously to the game session.")
		gamelift_createGameSessionCmd.Flags().String("name", "", "A descriptive label that is associated with a game session.")
		gamelift_createGameSessionCmd.MarkFlagRequired("maximum-player-session-count")
	})
	gameliftCmd.AddCommand(gamelift_createGameSessionCmd)
}
