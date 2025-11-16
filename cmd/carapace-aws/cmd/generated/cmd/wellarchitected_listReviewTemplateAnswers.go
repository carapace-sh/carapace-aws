package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listReviewTemplateAnswersCmd = &cobra.Command{
	Use:   "list-review-template-answers",
	Short: "List the answers of a review template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listReviewTemplateAnswersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_listReviewTemplateAnswersCmd).Standalone()

		wellarchitected_listReviewTemplateAnswersCmd.Flags().String("lens-alias", "", "")
		wellarchitected_listReviewTemplateAnswersCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
		wellarchitected_listReviewTemplateAnswersCmd.Flags().String("next-token", "", "")
		wellarchitected_listReviewTemplateAnswersCmd.Flags().String("pillar-id", "", "")
		wellarchitected_listReviewTemplateAnswersCmd.Flags().String("template-arn", "", "The ARN of the review template.")
		wellarchitected_listReviewTemplateAnswersCmd.MarkFlagRequired("lens-alias")
		wellarchitected_listReviewTemplateAnswersCmd.MarkFlagRequired("template-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_listReviewTemplateAnswersCmd)
}
