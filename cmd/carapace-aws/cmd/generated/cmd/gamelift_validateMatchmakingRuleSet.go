package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_validateMatchmakingRuleSetCmd = &cobra.Command{
	Use:   "validate-matchmaking-rule-set",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nValidates the syntax of a matchmaking rule or rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_validateMatchmakingRuleSetCmd).Standalone()

	gamelift_validateMatchmakingRuleSetCmd.Flags().String("rule-set-body", "", "A collection of matchmaking rules to validate, formatted as a JSON string.")
	gamelift_validateMatchmakingRuleSetCmd.MarkFlagRequired("rule-set-body")
	gameliftCmd.AddCommand(gamelift_validateMatchmakingRuleSetCmd)
}
