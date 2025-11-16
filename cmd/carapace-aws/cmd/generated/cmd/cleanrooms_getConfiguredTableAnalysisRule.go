package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getConfiguredTableAnalysisRuleCmd = &cobra.Command{
	Use:   "get-configured-table-analysis-rule",
	Short: "Retrieves a configured table analysis rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getConfiguredTableAnalysisRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getConfiguredTableAnalysisRuleCmd).Standalone()

		cleanrooms_getConfiguredTableAnalysisRuleCmd.Flags().String("analysis-rule-type", "", "The analysis rule to be retrieved.")
		cleanrooms_getConfiguredTableAnalysisRuleCmd.Flags().String("configured-table-identifier", "", "The unique identifier for the configured table to retrieve.")
		cleanrooms_getConfiguredTableAnalysisRuleCmd.MarkFlagRequired("analysis-rule-type")
		cleanrooms_getConfiguredTableAnalysisRuleCmd.MarkFlagRequired("configured-table-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getConfiguredTableAnalysisRuleCmd)
}
