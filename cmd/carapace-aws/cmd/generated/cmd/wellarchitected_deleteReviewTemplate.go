package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_deleteReviewTemplateCmd = &cobra.Command{
	Use:   "delete-review-template",
	Short: "Delete a review template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_deleteReviewTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_deleteReviewTemplateCmd).Standalone()

		wellarchitected_deleteReviewTemplateCmd.Flags().String("client-request-token", "", "")
		wellarchitected_deleteReviewTemplateCmd.Flags().String("template-arn", "", "The review template ARN.")
		wellarchitected_deleteReviewTemplateCmd.MarkFlagRequired("client-request-token")
		wellarchitected_deleteReviewTemplateCmd.MarkFlagRequired("template-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_deleteReviewTemplateCmd)
}
