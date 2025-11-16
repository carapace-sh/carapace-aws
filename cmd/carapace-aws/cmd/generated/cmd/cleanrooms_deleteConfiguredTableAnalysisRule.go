package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteConfiguredTableAnalysisRuleCmd = &cobra.Command{
	Use:   "delete-configured-table-analysis-rule",
	Short: "Deletes a configured table analysis rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteConfiguredTableAnalysisRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_deleteConfiguredTableAnalysisRuleCmd).Standalone()

		cleanrooms_deleteConfiguredTableAnalysisRuleCmd.Flags().String("analysis-rule-type", "", "The analysis rule type to be deleted.")
		cleanrooms_deleteConfiguredTableAnalysisRuleCmd.Flags().String("configured-table-identifier", "", "The unique identifier for the configured table that the analysis rule applies to.")
		cleanrooms_deleteConfiguredTableAnalysisRuleCmd.MarkFlagRequired("analysis-rule-type")
		cleanrooms_deleteConfiguredTableAnalysisRuleCmd.MarkFlagRequired("configured-table-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_deleteConfiguredTableAnalysisRuleCmd)
}
