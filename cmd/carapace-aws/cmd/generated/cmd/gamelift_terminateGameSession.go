package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_terminateGameSessionCmd = &cobra.Command{
	Use:   "terminate-game-session",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nEnds a game session that's currently in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_terminateGameSessionCmd).Standalone()

	gamelift_terminateGameSessionCmd.Flags().String("game-session-id", "", "A unique identifier for the game session to be terminated.")
	gamelift_terminateGameSessionCmd.Flags().String("termination-mode", "", "The method to use to terminate the game session.")
	gamelift_terminateGameSessionCmd.MarkFlagRequired("game-session-id")
	gamelift_terminateGameSessionCmd.MarkFlagRequired("termination-mode")
	gameliftCmd.AddCommand(gamelift_terminateGameSessionCmd)
}
