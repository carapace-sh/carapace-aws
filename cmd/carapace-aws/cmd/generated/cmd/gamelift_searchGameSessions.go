package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_searchGameSessionsCmd = &cobra.Command{
	Use:   "search-game-sessions",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves all active game sessions that match a set of search criteria and sorts them into a specified order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_searchGameSessionsCmd).Standalone()

	gamelift_searchGameSessionsCmd.Flags().String("alias-id", "", "A unique identifier for the alias associated with the fleet to search for active game sessions.")
	gamelift_searchGameSessionsCmd.Flags().String("filter-expression", "", "String containing the search criteria for the session search.")
	gamelift_searchGameSessionsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to search for active game sessions.")
	gamelift_searchGameSessionsCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_searchGameSessionsCmd.Flags().String("location", "", "A fleet location to search for game sessions.")
	gamelift_searchGameSessionsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gamelift_searchGameSessionsCmd.Flags().String("sort-expression", "", "Instructions on how to sort the search results.")
	gameliftCmd.AddCommand(gamelift_searchGameSessionsCmd)
}
