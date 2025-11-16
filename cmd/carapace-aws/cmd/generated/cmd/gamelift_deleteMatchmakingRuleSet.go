package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteMatchmakingRuleSetCmd = &cobra.Command{
	Use:   "delete-matchmaking-rule-set",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nDeletes an existing matchmaking rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteMatchmakingRuleSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_deleteMatchmakingRuleSetCmd).Standalone()

		gamelift_deleteMatchmakingRuleSetCmd.Flags().String("name", "", "A unique identifier for the matchmaking rule set to be deleted.")
		gamelift_deleteMatchmakingRuleSetCmd.MarkFlagRequired("name")
	})
	gameliftCmd.AddCommand(gamelift_deleteMatchmakingRuleSetCmd)
}
