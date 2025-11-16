package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getFindingRecommendationCmd = &cobra.Command{
	Use:   "get-finding-recommendation",
	Short: "Retrieves information about a finding recommendation for the specified analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getFindingRecommendationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_getFindingRecommendationCmd).Standalone()

		accessanalyzer_getFindingRecommendationCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) used to generate the finding recommendation.")
		accessanalyzer_getFindingRecommendationCmd.Flags().String("id", "", "The unique ID for the finding recommendation.")
		accessanalyzer_getFindingRecommendationCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		accessanalyzer_getFindingRecommendationCmd.Flags().String("next-token", "", "A token used for pagination of results returned.")
		accessanalyzer_getFindingRecommendationCmd.MarkFlagRequired("analyzer-arn")
		accessanalyzer_getFindingRecommendationCmd.MarkFlagRequired("id")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_getFindingRecommendationCmd)
}
