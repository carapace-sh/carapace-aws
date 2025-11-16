package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describePlayerSessionsCmd = &cobra.Command{
	Use:   "describe-player-sessions",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves properties for one or more player sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describePlayerSessionsCmd).Standalone()

	gamelift_describePlayerSessionsCmd.Flags().String("game-session-id", "", "A unique identifier for the game session to retrieve player sessions for.")
	gamelift_describePlayerSessionsCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_describePlayerSessionsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gamelift_describePlayerSessionsCmd.Flags().String("player-id", "", "A unique identifier for a player to retrieve player sessions for.")
	gamelift_describePlayerSessionsCmd.Flags().String("player-session-id", "", "A unique identifier for a player session to retrieve.")
	gamelift_describePlayerSessionsCmd.Flags().String("player-session-status-filter", "", "Player session status to filter results on.")
	gameliftCmd.AddCommand(gamelift_describePlayerSessionsCmd)
}
