package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createMatchmakingRuleSetCmd = &cobra.Command{
	Use:   "create-matchmaking-rule-set",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nCreates a new rule set for FlexMatch matchmaking.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createMatchmakingRuleSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createMatchmakingRuleSetCmd).Standalone()

		gamelift_createMatchmakingRuleSetCmd.Flags().String("name", "", "A unique identifier for the matchmaking rule set.")
		gamelift_createMatchmakingRuleSetCmd.Flags().String("rule-set-body", "", "A collection of matchmaking rules, formatted as a JSON string.")
		gamelift_createMatchmakingRuleSetCmd.Flags().String("tags", "", "A list of labels to assign to the new matchmaking rule set resource.")
		gamelift_createMatchmakingRuleSetCmd.MarkFlagRequired("name")
		gamelift_createMatchmakingRuleSetCmd.MarkFlagRequired("rule-set-body")
	})
	gameliftCmd.AddCommand(gamelift_createMatchmakingRuleSetCmd)
}
