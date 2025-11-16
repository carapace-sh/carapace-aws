package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateReviewTemplateCmd = &cobra.Command{
	Use:   "update-review-template",
	Short: "Update a review template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateReviewTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_updateReviewTemplateCmd).Standalone()

		wellarchitected_updateReviewTemplateCmd.Flags().String("description", "", "The review template description.")
		wellarchitected_updateReviewTemplateCmd.Flags().String("lenses-to-associate", "", "A list of lens aliases or ARNs to apply to the review template.")
		wellarchitected_updateReviewTemplateCmd.Flags().String("lenses-to-disassociate", "", "A list of lens aliases or ARNs to unapply to the review template.")
		wellarchitected_updateReviewTemplateCmd.Flags().String("notes", "", "")
		wellarchitected_updateReviewTemplateCmd.Flags().String("template-arn", "", "The review template ARN.")
		wellarchitected_updateReviewTemplateCmd.Flags().String("template-name", "", "The review template name.")
		wellarchitected_updateReviewTemplateCmd.MarkFlagRequired("template-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_updateReviewTemplateCmd)
}
