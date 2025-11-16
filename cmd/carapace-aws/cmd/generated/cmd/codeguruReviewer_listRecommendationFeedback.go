package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_listRecommendationFeedbackCmd = &cobra.Command{
	Use:   "list-recommendation-feedback",
	Short: "Returns a list of [RecommendationFeedbackSummary](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RecommendationFeedbackSummary.html) objects that contain customer recommendation feedback for all CodeGuru Reviewer users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_listRecommendationFeedbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_listRecommendationFeedbackCmd).Standalone()

		codeguruReviewer_listRecommendationFeedbackCmd.Flags().String("code-review-arn", "", "The Amazon Resource Name (ARN) of the [CodeReview](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html) object.")
		codeguruReviewer_listRecommendationFeedbackCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		codeguruReviewer_listRecommendationFeedbackCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
		codeguruReviewer_listRecommendationFeedbackCmd.Flags().String("recommendation-ids", "", "Used to query the recommendation feedback for a given recommendation.")
		codeguruReviewer_listRecommendationFeedbackCmd.Flags().String("user-ids", "", "An Amazon Web Services user's account ID or Amazon Resource Name (ARN).")
		codeguruReviewer_listRecommendationFeedbackCmd.MarkFlagRequired("code-review-arn")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_listRecommendationFeedbackCmd)
}
