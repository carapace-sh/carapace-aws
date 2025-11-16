package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeMatchmakingRuleSetsCmd = &cobra.Command{
	Use:   "describe-matchmaking-rule-sets",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves the details for FlexMatch matchmaking rule sets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeMatchmakingRuleSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeMatchmakingRuleSetsCmd).Standalone()

		gamelift_describeMatchmakingRuleSetsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_describeMatchmakingRuleSetsCmd.Flags().String("names", "", "A list of one or more matchmaking rule set names to retrieve details for.")
		gamelift_describeMatchmakingRuleSetsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	})
	gameliftCmd.AddCommand(gamelift_describeMatchmakingRuleSetsCmd)
}
