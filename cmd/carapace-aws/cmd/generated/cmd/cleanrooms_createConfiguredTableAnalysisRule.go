package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createConfiguredTableAnalysisRuleCmd = &cobra.Command{
	Use:   "create-configured-table-analysis-rule",
	Short: "Creates a new analysis rule for a configured table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createConfiguredTableAnalysisRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_createConfiguredTableAnalysisRuleCmd).Standalone()

		cleanrooms_createConfiguredTableAnalysisRuleCmd.Flags().String("analysis-rule-policy", "", "The analysis rule policy that was created for the configured table.")
		cleanrooms_createConfiguredTableAnalysisRuleCmd.Flags().String("analysis-rule-type", "", "The type of analysis rule.")
		cleanrooms_createConfiguredTableAnalysisRuleCmd.Flags().String("configured-table-identifier", "", "The identifier for the configured table to create the analysis rule for.")
		cleanrooms_createConfiguredTableAnalysisRuleCmd.MarkFlagRequired("analysis-rule-policy")
		cleanrooms_createConfiguredTableAnalysisRuleCmd.MarkFlagRequired("analysis-rule-type")
		cleanrooms_createConfiguredTableAnalysisRuleCmd.MarkFlagRequired("configured-table-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_createConfiguredTableAnalysisRuleCmd)
}
