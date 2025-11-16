package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeMatchmakingConfigurationsCmd = &cobra.Command{
	Use:   "describe-matchmaking-configurations",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves the details of FlexMatch matchmaking configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeMatchmakingConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeMatchmakingConfigurationsCmd).Standalone()

		gamelift_describeMatchmakingConfigurationsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_describeMatchmakingConfigurationsCmd.Flags().String("names", "", "A unique identifier for the matchmaking configuration(s) to retrieve.")
		gamelift_describeMatchmakingConfigurationsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_describeMatchmakingConfigurationsCmd.Flags().String("rule-set-name", "", "A unique identifier for the matchmaking rule set.")
	})
	gameliftCmd.AddCommand(gamelift_describeMatchmakingConfigurationsCmd)
}
