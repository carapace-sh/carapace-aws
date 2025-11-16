package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd = &cobra.Command{
	Use:   "delete-configured-table-association-analysis-rule",
	Short: "Deletes an analysis rule for a configured table association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd).Standalone()

		cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd.Flags().String("analysis-rule-type", "", "The type of the analysis rule that you want to delete.")
		cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd.Flags().String("configured-table-association-identifier", "", "The identiﬁer for the conﬁgured table association that's related to the analysis rule that you want to delete.")
		cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership that the configured table association belongs to.")
		cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("analysis-rule-type")
		cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("configured-table-association-identifier")
		cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_deleteConfiguredTableAssociationAnalysisRuleCmd)
}
