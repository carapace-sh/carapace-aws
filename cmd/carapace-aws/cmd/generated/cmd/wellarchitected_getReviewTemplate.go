package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getReviewTemplateCmd = &cobra.Command{
	Use:   "get-review-template",
	Short: "Get review template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getReviewTemplateCmd).Standalone()

	wellarchitected_getReviewTemplateCmd.Flags().String("template-arn", "", "The review template ARN.")
	wellarchitected_getReviewTemplateCmd.MarkFlagRequired("template-arn")
	wellarchitectedCmd.AddCommand(wellarchitected_getReviewTemplateCmd)
}
