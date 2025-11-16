package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_listRecommendationsCmd = &cobra.Command{
	Use:   "list-recommendations",
	Short: "Returns the list of all recommendations for a completed code review.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_listRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_listRecommendationsCmd).Standalone()

		codeguruReviewer_listRecommendationsCmd.Flags().String("code-review-arn", "", "The Amazon Resource Name (ARN) of the [CodeReview](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html) object.")
		codeguruReviewer_listRecommendationsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		codeguruReviewer_listRecommendationsCmd.Flags().String("next-token", "", "Pagination token.")
		codeguruReviewer_listRecommendationsCmd.MarkFlagRequired("code-review-arn")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_listRecommendationsCmd)
}
