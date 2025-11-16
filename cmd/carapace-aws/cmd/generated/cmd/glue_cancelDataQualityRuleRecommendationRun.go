package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_cancelDataQualityRuleRecommendationRunCmd = &cobra.Command{
	Use:   "cancel-data-quality-rule-recommendation-run",
	Short: "Cancels the specified recommendation run that was being used to generate rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_cancelDataQualityRuleRecommendationRunCmd).Standalone()

	glue_cancelDataQualityRuleRecommendationRunCmd.Flags().String("run-id", "", "The unique run identifier associated with this run.")
	glue_cancelDataQualityRuleRecommendationRunCmd.MarkFlagRequired("run-id")
	glueCmd.AddCommand(glue_cancelDataQualityRuleRecommendationRunCmd)
}
