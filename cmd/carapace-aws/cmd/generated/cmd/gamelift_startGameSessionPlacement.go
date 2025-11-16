package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_startGameSessionPlacementCmd = &cobra.Command{
	Use:   "start-game-session-placement",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nMakes a request to start a new game session using a game session queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_startGameSessionPlacementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_startGameSessionPlacementCmd).Standalone()

		gamelift_startGameSessionPlacementCmd.Flags().String("desired-player-sessions", "", "Set of information on each player to create a player session for.")
		gamelift_startGameSessionPlacementCmd.Flags().String("game-properties", "", "A set of key-value pairs that can store custom data in a game session.")
		gamelift_startGameSessionPlacementCmd.Flags().String("game-session-data", "", "A set of custom game session properties, formatted as a single string value.")
		gamelift_startGameSessionPlacementCmd.Flags().String("game-session-name", "", "A descriptive label that is associated with a game session.")
		gamelift_startGameSessionPlacementCmd.Flags().String("game-session-queue-name", "", "Name of the queue to use to place the new game session.")
		gamelift_startGameSessionPlacementCmd.Flags().String("maximum-player-session-count", "", "The maximum number of players that can be connected simultaneously to the game session.")
		gamelift_startGameSessionPlacementCmd.Flags().String("placement-id", "", "A unique identifier to assign to the new game session placement.")
		gamelift_startGameSessionPlacementCmd.Flags().String("player-latencies", "", "A set of values, expressed in milliseconds, that indicates the amount of latency that a player experiences when connected to Amazon Web Services Regions.")
		gamelift_startGameSessionPlacementCmd.Flags().String("priority-configuration-override", "", "A prioritized list of locations to use for the game session placement and instructions on how to use it.")
		gamelift_startGameSessionPlacementCmd.MarkFlagRequired("game-session-queue-name")
		gamelift_startGameSessionPlacementCmd.MarkFlagRequired("maximum-player-session-count")
		gamelift_startGameSessionPlacementCmd.MarkFlagRequired("placement-id")
	})
	gameliftCmd.AddCommand(gamelift_startGameSessionPlacementCmd)
}
