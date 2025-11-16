package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns the list of tags associated with an associated repository resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_listTagsForResourceCmd).Standalone()

		codeguruReviewer_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the [RepositoryAssociation](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html) object.")
		codeguruReviewer_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_listTagsForResourceCmd)
}
