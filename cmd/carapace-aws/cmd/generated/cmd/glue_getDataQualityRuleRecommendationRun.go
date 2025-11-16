package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDataQualityRuleRecommendationRunCmd = &cobra.Command{
	Use:   "get-data-quality-rule-recommendation-run",
	Short: "Gets the specified recommendation run that was used to generate rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDataQualityRuleRecommendationRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getDataQualityRuleRecommendationRunCmd).Standalone()

		glue_getDataQualityRuleRecommendationRunCmd.Flags().String("run-id", "", "The unique run identifier associated with this run.")
		glue_getDataQualityRuleRecommendationRunCmd.MarkFlagRequired("run-id")
	})
	glueCmd.AddCommand(glue_getDataQualityRuleRecommendationRunCmd)
}
