package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createPlayerSessionsCmd = &cobra.Command{
	Use:   "create-player-sessions",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nReserves open slots in a game session for a group of players.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createPlayerSessionsCmd).Standalone()

	gamelift_createPlayerSessionsCmd.Flags().String("game-session-id", "", "A unique identifier for the game session to add players to.")
	gamelift_createPlayerSessionsCmd.Flags().String("player-data-map", "", "Map of string pairs, each specifying a player ID and a set of developer-defined information related to the player.")
	gamelift_createPlayerSessionsCmd.Flags().String("player-ids", "", "List of unique identifiers for the players to be added.")
	gamelift_createPlayerSessionsCmd.MarkFlagRequired("game-session-id")
	gamelift_createPlayerSessionsCmd.MarkFlagRequired("player-ids")
	gameliftCmd.AddCommand(gamelift_createPlayerSessionsCmd)
}
