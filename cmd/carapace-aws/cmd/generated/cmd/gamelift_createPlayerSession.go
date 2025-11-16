package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createPlayerSessionCmd = &cobra.Command{
	Use:   "create-player-session",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nReserves an open player slot in a game session for a player.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createPlayerSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createPlayerSessionCmd).Standalone()

		gamelift_createPlayerSessionCmd.Flags().String("game-session-id", "", "A unique identifier for the game session to add a player to.")
		gamelift_createPlayerSessionCmd.Flags().String("player-data", "", "Developer-defined information related to a player.")
		gamelift_createPlayerSessionCmd.Flags().String("player-id", "", "A unique identifier for a player.")
		gamelift_createPlayerSessionCmd.MarkFlagRequired("game-session-id")
		gamelift_createPlayerSessionCmd.MarkFlagRequired("player-id")
	})
	gameliftCmd.AddCommand(gamelift_createPlayerSessionCmd)
}
