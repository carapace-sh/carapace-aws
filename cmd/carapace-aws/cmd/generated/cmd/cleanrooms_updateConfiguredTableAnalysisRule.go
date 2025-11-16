package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateConfiguredTableAnalysisRuleCmd = &cobra.Command{
	Use:   "update-configured-table-analysis-rule",
	Short: "Updates a configured table analysis rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateConfiguredTableAnalysisRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_updateConfiguredTableAnalysisRuleCmd).Standalone()

		cleanrooms_updateConfiguredTableAnalysisRuleCmd.Flags().String("analysis-rule-policy", "", "The new analysis rule policy for the configured table analysis rule.")
		cleanrooms_updateConfiguredTableAnalysisRuleCmd.Flags().String("analysis-rule-type", "", "The analysis rule type to be updated.")
		cleanrooms_updateConfiguredTableAnalysisRuleCmd.Flags().String("configured-table-identifier", "", "The unique identifier for the configured table that the analysis rule applies to.")
		cleanrooms_updateConfiguredTableAnalysisRuleCmd.MarkFlagRequired("analysis-rule-policy")
		cleanrooms_updateConfiguredTableAnalysisRuleCmd.MarkFlagRequired("analysis-rule-type")
		cleanrooms_updateConfiguredTableAnalysisRuleCmd.MarkFlagRequired("configured-table-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_updateConfiguredTableAnalysisRuleCmd)
}
