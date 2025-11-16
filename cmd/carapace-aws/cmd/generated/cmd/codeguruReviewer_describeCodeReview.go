package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_describeCodeReviewCmd = &cobra.Command{
	Use:   "describe-code-review",
	Short: "Returns the metadata associated with the code review along with its status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_describeCodeReviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_describeCodeReviewCmd).Standalone()

		codeguruReviewer_describeCodeReviewCmd.Flags().String("code-review-arn", "", "The Amazon Resource Name (ARN) of the [CodeReview](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html) object.")
		codeguruReviewer_describeCodeReviewCmd.MarkFlagRequired("code-review-arn")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_describeCodeReviewCmd)
}
