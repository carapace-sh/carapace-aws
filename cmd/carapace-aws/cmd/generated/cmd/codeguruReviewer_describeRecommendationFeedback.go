package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_describeRecommendationFeedbackCmd = &cobra.Command{
	Use:   "describe-recommendation-feedback",
	Short: "Describes the customer feedback for a CodeGuru Reviewer recommendation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_describeRecommendationFeedbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_describeRecommendationFeedbackCmd).Standalone()

		codeguruReviewer_describeRecommendationFeedbackCmd.Flags().String("code-review-arn", "", "The Amazon Resource Name (ARN) of the [CodeReview](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html) object.")
		codeguruReviewer_describeRecommendationFeedbackCmd.Flags().String("recommendation-id", "", "The recommendation ID that can be used to track the provided recommendations and then to collect the feedback.")
		codeguruReviewer_describeRecommendationFeedbackCmd.Flags().String("user-id", "", "Optional parameter to describe the feedback for a given user.")
		codeguruReviewer_describeRecommendationFeedbackCmd.MarkFlagRequired("code-review-arn")
		codeguruReviewer_describeRecommendationFeedbackCmd.MarkFlagRequired("recommendation-id")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_describeRecommendationFeedbackCmd)
}
