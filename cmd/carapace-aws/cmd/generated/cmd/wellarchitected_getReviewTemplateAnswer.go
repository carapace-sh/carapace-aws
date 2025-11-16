package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getReviewTemplateAnswerCmd = &cobra.Command{
	Use:   "get-review-template-answer",
	Short: "Get review template answer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getReviewTemplateAnswerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_getReviewTemplateAnswerCmd).Standalone()

		wellarchitected_getReviewTemplateAnswerCmd.Flags().String("lens-alias", "", "")
		wellarchitected_getReviewTemplateAnswerCmd.Flags().String("question-id", "", "")
		wellarchitected_getReviewTemplateAnswerCmd.Flags().String("template-arn", "", "The review template ARN.")
		wellarchitected_getReviewTemplateAnswerCmd.MarkFlagRequired("lens-alias")
		wellarchitected_getReviewTemplateAnswerCmd.MarkFlagRequired("question-id")
		wellarchitected_getReviewTemplateAnswerCmd.MarkFlagRequired("template-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_getReviewTemplateAnswerCmd)
}
