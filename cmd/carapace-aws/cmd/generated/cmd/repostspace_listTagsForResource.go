package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns the tags that are associated with the AWS re:Post Private resource specified by the resourceArn.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_listTagsForResourceCmd).Standalone()

	repostspace_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that the tags are associated with.")
	repostspace_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	repostspaceCmd.AddCommand(repostspace_listTagsForResourceCmd)
}
