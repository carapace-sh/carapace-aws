package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeGameSessionQueuesCmd = &cobra.Command{
	Use:   "describe-game-session-queues",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves the properties for one or more game session queues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeGameSessionQueuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeGameSessionQueuesCmd).Standalone()

		gamelift_describeGameSessionQueuesCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_describeGameSessionQueuesCmd.Flags().String("names", "", "A list of queue names to retrieve information for.")
		gamelift_describeGameSessionQueuesCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	})
	gameliftCmd.AddCommand(gamelift_describeGameSessionQueuesCmd)
}
