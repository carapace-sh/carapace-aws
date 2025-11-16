package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd = &cobra.Command{
	Use:   "update-configured-table-association-analysis-rule",
	Short: "Updates the analysis rule for a configured table association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd).Standalone()

		cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd.Flags().String("analysis-rule-policy", "", "The updated analysis rule policy for the conﬁgured table association.")
		cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd.Flags().String("analysis-rule-type", "", "The analysis rule type that you want to update.")
		cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd.Flags().String("configured-table-association-identifier", "", "The identifier for the configured table association to update.")
		cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership that the configured table association belongs to.")
		cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("analysis-rule-policy")
		cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("analysis-rule-type")
		cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("configured-table-association-identifier")
		cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_updateConfiguredTableAssociationAnalysisRuleCmd)
}
