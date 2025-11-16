package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateReviewTemplateAnswerCmd = &cobra.Command{
	Use:   "update-review-template-answer",
	Short: "Update a review template answer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateReviewTemplateAnswerCmd).Standalone()

	wellarchitected_updateReviewTemplateAnswerCmd.Flags().String("choice-updates", "", "A list of choices to be updated.")
	wellarchitected_updateReviewTemplateAnswerCmd.Flags().String("is-applicable", "", "")
	wellarchitected_updateReviewTemplateAnswerCmd.Flags().String("lens-alias", "", "")
	wellarchitected_updateReviewTemplateAnswerCmd.Flags().String("notes", "", "")
	wellarchitected_updateReviewTemplateAnswerCmd.Flags().String("question-id", "", "")
	wellarchitected_updateReviewTemplateAnswerCmd.Flags().String("reason", "", "The update reason.")
	wellarchitected_updateReviewTemplateAnswerCmd.Flags().String("selected-choices", "", "")
	wellarchitected_updateReviewTemplateAnswerCmd.Flags().String("template-arn", "", "The review template ARN.")
	wellarchitected_updateReviewTemplateAnswerCmd.MarkFlagRequired("lens-alias")
	wellarchitected_updateReviewTemplateAnswerCmd.MarkFlagRequired("question-id")
	wellarchitected_updateReviewTemplateAnswerCmd.MarkFlagRequired("template-arn")
	wellarchitectedCmd.AddCommand(wellarchitected_updateReviewTemplateAnswerCmd)
}
