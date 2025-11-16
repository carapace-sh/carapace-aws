package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateGameSessionCmd = &cobra.Command{
	Use:   "update-game-session",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nUpdates the mutable properties of a game session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateGameSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_updateGameSessionCmd).Standalone()

		gamelift_updateGameSessionCmd.Flags().String("game-properties", "", "A set of key-value pairs that can store custom data in a game session.")
		gamelift_updateGameSessionCmd.Flags().String("game-session-id", "", "A unique identifier for the game session to update.")
		gamelift_updateGameSessionCmd.Flags().String("maximum-player-session-count", "", "The maximum number of players that can be connected simultaneously to the game session.")
		gamelift_updateGameSessionCmd.Flags().String("name", "", "A descriptive label that is associated with a game session.")
		gamelift_updateGameSessionCmd.Flags().String("player-session-creation-policy", "", "A policy that determines whether the game session is accepting new players.")
		gamelift_updateGameSessionCmd.Flags().String("protection-policy", "", "Game session protection policy to apply to this game session only.")
		gamelift_updateGameSessionCmd.MarkFlagRequired("game-session-id")
	})
	gameliftCmd.AddCommand(gamelift_updateGameSessionCmd)
}
