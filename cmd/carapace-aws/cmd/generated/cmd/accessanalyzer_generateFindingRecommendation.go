package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_generateFindingRecommendationCmd = &cobra.Command{
	Use:   "generate-finding-recommendation",
	Short: "Creates a recommendation for an unused permissions finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_generateFindingRecommendationCmd).Standalone()

	accessanalyzer_generateFindingRecommendationCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) used to generate the finding recommendation.")
	accessanalyzer_generateFindingRecommendationCmd.Flags().String("id", "", "The unique ID for the finding recommendation.")
	accessanalyzer_generateFindingRecommendationCmd.MarkFlagRequired("analyzer-arn")
	accessanalyzer_generateFindingRecommendationCmd.MarkFlagRequired("id")
	accessanalyzerCmd.AddCommand(accessanalyzer_generateFindingRecommendationCmd)
}
