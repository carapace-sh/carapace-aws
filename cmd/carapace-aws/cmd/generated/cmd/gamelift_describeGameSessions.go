package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeGameSessionsCmd = &cobra.Command{
	Use:   "describe-game-sessions",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves a set of one or more game sessions in a specific fleet location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeGameSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeGameSessionsCmd).Standalone()

		gamelift_describeGameSessionsCmd.Flags().String("alias-id", "", "A unique identifier for the alias associated with the fleet to retrieve game sessions for.")
		gamelift_describeGameSessionsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to retrieve game sessions for.")
		gamelift_describeGameSessionsCmd.Flags().String("game-session-id", "", "A unique identifier for the game session to retrieve.")
		gamelift_describeGameSessionsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_describeGameSessionsCmd.Flags().String("location", "", "A fleet location to get game sessions for.")
		gamelift_describeGameSessionsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_describeGameSessionsCmd.Flags().String("status-filter", "", "Game session status to filter results on.")
	})
	gameliftCmd.AddCommand(gamelift_describeGameSessionsCmd)
}
