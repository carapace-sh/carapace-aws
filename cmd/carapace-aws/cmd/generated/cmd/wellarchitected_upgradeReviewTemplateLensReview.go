package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_upgradeReviewTemplateLensReviewCmd = &cobra.Command{
	Use:   "upgrade-review-template-lens-review",
	Short: "Upgrade the lens review of a review template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_upgradeReviewTemplateLensReviewCmd).Standalone()

	wellarchitected_upgradeReviewTemplateLensReviewCmd.Flags().String("client-request-token", "", "")
	wellarchitected_upgradeReviewTemplateLensReviewCmd.Flags().String("lens-alias", "", "")
	wellarchitected_upgradeReviewTemplateLensReviewCmd.Flags().String("template-arn", "", "The ARN of the review template.")
	wellarchitected_upgradeReviewTemplateLensReviewCmd.MarkFlagRequired("lens-alias")
	wellarchitected_upgradeReviewTemplateLensReviewCmd.MarkFlagRequired("template-arn")
	wellarchitectedCmd.AddCommand(wellarchitected_upgradeReviewTemplateLensReviewCmd)
}
