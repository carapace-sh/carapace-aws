package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_putRecommendationFeedbackCmd = &cobra.Command{
	Use:   "put-recommendation-feedback",
	Short: "Stores customer feedback for a CodeGuru Reviewer recommendation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_putRecommendationFeedbackCmd).Standalone()

	codeguruReviewer_putRecommendationFeedbackCmd.Flags().String("code-review-arn", "", "The Amazon Resource Name (ARN) of the [CodeReview](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html) object.")
	codeguruReviewer_putRecommendationFeedbackCmd.Flags().String("reactions", "", "List for storing reactions.")
	codeguruReviewer_putRecommendationFeedbackCmd.Flags().String("recommendation-id", "", "The recommendation ID that can be used to track the provided recommendations and then to collect the feedback.")
	codeguruReviewer_putRecommendationFeedbackCmd.MarkFlagRequired("code-review-arn")
	codeguruReviewer_putRecommendationFeedbackCmd.MarkFlagRequired("reactions")
	codeguruReviewer_putRecommendationFeedbackCmd.MarkFlagRequired("recommendation-id")
	codeguruReviewerCmd.AddCommand(codeguruReviewer_putRecommendationFeedbackCmd)
}
