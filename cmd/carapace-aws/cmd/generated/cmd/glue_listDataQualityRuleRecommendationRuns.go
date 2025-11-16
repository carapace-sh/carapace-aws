package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listDataQualityRuleRecommendationRunsCmd = &cobra.Command{
	Use:   "list-data-quality-rule-recommendation-runs",
	Short: "Lists the recommendation runs meeting the filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listDataQualityRuleRecommendationRunsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listDataQualityRuleRecommendationRunsCmd).Standalone()

		glue_listDataQualityRuleRecommendationRunsCmd.Flags().String("filter", "", "The filter criteria.")
		glue_listDataQualityRuleRecommendationRunsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		glue_listDataQualityRuleRecommendationRunsCmd.Flags().String("next-token", "", "A paginated token to offset the results.")
	})
	glueCmd.AddCommand(glue_listDataQualityRuleRecommendationRunsCmd)
}
