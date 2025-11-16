package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateReviewTemplateLensReviewCmd = &cobra.Command{
	Use:   "update-review-template-lens-review",
	Short: "Update a lens review associated with a review template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateReviewTemplateLensReviewCmd).Standalone()

	wellarchitected_updateReviewTemplateLensReviewCmd.Flags().String("lens-alias", "", "")
	wellarchitected_updateReviewTemplateLensReviewCmd.Flags().String("lens-notes", "", "")
	wellarchitected_updateReviewTemplateLensReviewCmd.Flags().String("pillar-notes", "", "")
	wellarchitected_updateReviewTemplateLensReviewCmd.Flags().String("template-arn", "", "The review template ARN.")
	wellarchitected_updateReviewTemplateLensReviewCmd.MarkFlagRequired("lens-alias")
	wellarchitected_updateReviewTemplateLensReviewCmd.MarkFlagRequired("template-arn")
	wellarchitectedCmd.AddCommand(wellarchitected_updateReviewTemplateLensReviewCmd)
}
