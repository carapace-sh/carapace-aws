package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_createCodeReviewCmd = &cobra.Command{
	Use:   "create-code-review",
	Short: "Use to create a code review with a [CodeReviewType](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReviewType.html) of `RepositoryAnalysis`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_createCodeReviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_createCodeReviewCmd).Standalone()

		codeguruReviewer_createCodeReviewCmd.Flags().String("client-request-token", "", "Amazon CodeGuru Reviewer uses this value to prevent the accidental creation of duplicate code reviews if there are failures and retries.")
		codeguruReviewer_createCodeReviewCmd.Flags().String("name", "", "The name of the code review.")
		codeguruReviewer_createCodeReviewCmd.Flags().String("repository-association-arn", "", "The Amazon Resource Name (ARN) of the [RepositoryAssociation](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html) object.")
		codeguruReviewer_createCodeReviewCmd.Flags().String("type", "", "The type of code review to create.")
		codeguruReviewer_createCodeReviewCmd.MarkFlagRequired("name")
		codeguruReviewer_createCodeReviewCmd.MarkFlagRequired("repository-association-arn")
		codeguruReviewer_createCodeReviewCmd.MarkFlagRequired("type")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_createCodeReviewCmd)
}
