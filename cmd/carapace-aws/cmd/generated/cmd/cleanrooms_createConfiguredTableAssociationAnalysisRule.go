package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd = &cobra.Command{
	Use:   "create-configured-table-association-analysis-rule",
	Short: "Creates a new analysis rule for an associated configured table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd).Standalone()

	cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd.Flags().String("analysis-rule-policy", "", "The analysis rule policy that was created for the configured table association.")
	cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd.Flags().String("analysis-rule-type", "", "The type of analysis rule.")
	cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd.Flags().String("configured-table-association-identifier", "", "The unique ID for the configured table association.")
	cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership that the configured table association belongs to.")
	cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("analysis-rule-policy")
	cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("analysis-rule-type")
	cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("configured-table-association-identifier")
	cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_createConfiguredTableAssociationAnalysisRuleCmd)
}
