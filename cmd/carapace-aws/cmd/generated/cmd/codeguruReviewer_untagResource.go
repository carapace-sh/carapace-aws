package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from an associated repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_untagResourceCmd).Standalone()

	codeguruReviewer_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the [RepositoryAssociation](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html) object.")
	codeguruReviewer_untagResourceCmd.Flags().String("tag-keys", "", "A list of the keys for each tag you want to remove from an associated repository.")
	codeguruReviewer_untagResourceCmd.MarkFlagRequired("resource-arn")
	codeguruReviewer_untagResourceCmd.MarkFlagRequired("tag-keys")
	codeguruReviewerCmd.AddCommand(codeguruReviewer_untagResourceCmd)
}
