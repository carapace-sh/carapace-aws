package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeGameSessionDetailsCmd = &cobra.Command{
	Use:   "describe-game-session-details",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves additional game session properties, including the game session protection policy in force, a set of one or more game sessions in a specific fleet location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeGameSessionDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeGameSessionDetailsCmd).Standalone()

		gamelift_describeGameSessionDetailsCmd.Flags().String("alias-id", "", "A unique identifier for the alias associated with the fleet to retrieve all game sessions for.")
		gamelift_describeGameSessionDetailsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to retrieve all game sessions active on the fleet.")
		gamelift_describeGameSessionDetailsCmd.Flags().String("game-session-id", "", "A unique identifier for the game session to retrieve.")
		gamelift_describeGameSessionDetailsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_describeGameSessionDetailsCmd.Flags().String("location", "", "A fleet location to get game session details for.")
		gamelift_describeGameSessionDetailsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_describeGameSessionDetailsCmd.Flags().String("status-filter", "", "Game session status to filter results on.")
	})
	gameliftCmd.AddCommand(gamelift_describeGameSessionDetailsCmd)
}
