package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to an associated repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_tagResourceCmd).Standalone()

	codeguruReviewer_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the [RepositoryAssociation](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html) object.")
	codeguruReviewer_tagResourceCmd.Flags().String("tags", "", "An array of key-value pairs used to tag an associated repository.")
	codeguruReviewer_tagResourceCmd.MarkFlagRequired("resource-arn")
	codeguruReviewer_tagResourceCmd.MarkFlagRequired("tags")
	codeguruReviewerCmd.AddCommand(codeguruReviewer_tagResourceCmd)
}
