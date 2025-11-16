package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getReviewTemplateLensReviewCmd = &cobra.Command{
	Use:   "get-review-template-lens-review",
	Short: "Get a lens review associated with a review template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getReviewTemplateLensReviewCmd).Standalone()

	wellarchitected_getReviewTemplateLensReviewCmd.Flags().String("lens-alias", "", "")
	wellarchitected_getReviewTemplateLensReviewCmd.Flags().String("template-arn", "", "The review template ARN.")
	wellarchitected_getReviewTemplateLensReviewCmd.MarkFlagRequired("lens-alias")
	wellarchitected_getReviewTemplateLensReviewCmd.MarkFlagRequired("template-arn")
	wellarchitectedCmd.AddCommand(wellarchitected_getReviewTemplateLensReviewCmd)
}
