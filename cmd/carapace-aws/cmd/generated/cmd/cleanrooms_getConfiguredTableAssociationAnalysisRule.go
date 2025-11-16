package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd = &cobra.Command{
	Use:   "get-configured-table-association-analysis-rule",
	Short: "Retrieves the analysis rule for a configured table association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd).Standalone()

		cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd.Flags().String("analysis-rule-type", "", "The type of analysis rule that you want to retrieve.")
		cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd.Flags().String("configured-table-association-identifier", "", "The identiﬁer for the conﬁgured table association that's related to the analysis rule.")
		cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership that the configured table association belongs to.")
		cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("analysis-rule-type")
		cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("configured-table-association-identifier")
		cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getConfiguredTableAssociationAnalysisRuleCmd)
}
