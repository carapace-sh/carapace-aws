package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createReviewTemplateCmd = &cobra.Command{
	Use:   "create-review-template",
	Short: "Create a review template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createReviewTemplateCmd).Standalone()

	wellarchitected_createReviewTemplateCmd.Flags().String("client-request-token", "", "")
	wellarchitected_createReviewTemplateCmd.Flags().String("description", "", "The review template description.")
	wellarchitected_createReviewTemplateCmd.Flags().String("lenses", "", "Lenses applied to the review template.")
	wellarchitected_createReviewTemplateCmd.Flags().String("notes", "", "")
	wellarchitected_createReviewTemplateCmd.Flags().String("tags", "", "The tags assigned to the review template.")
	wellarchitected_createReviewTemplateCmd.Flags().String("template-name", "", "Name of the review template.")
	wellarchitected_createReviewTemplateCmd.MarkFlagRequired("client-request-token")
	wellarchitected_createReviewTemplateCmd.MarkFlagRequired("description")
	wellarchitected_createReviewTemplateCmd.MarkFlagRequired("lenses")
	wellarchitected_createReviewTemplateCmd.MarkFlagRequired("template-name")
	wellarchitectedCmd.AddCommand(wellarchitected_createReviewTemplateCmd)
}
