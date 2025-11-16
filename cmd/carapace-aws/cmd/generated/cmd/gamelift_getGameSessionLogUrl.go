package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_getGameSessionLogUrlCmd = &cobra.Command{
	Use:   "get-game-session-log-url",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves the location of stored game session logs for a specified game session on Amazon GameLift Servers managed fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_getGameSessionLogUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_getGameSessionLogUrlCmd).Standalone()

		gamelift_getGameSessionLogUrlCmd.Flags().String("game-session-id", "", "A unique identifier for the game session to get logs for.")
		gamelift_getGameSessionLogUrlCmd.MarkFlagRequired("game-session-id")
	})
	gameliftCmd.AddCommand(gamelift_getGameSessionLogUrlCmd)
}
