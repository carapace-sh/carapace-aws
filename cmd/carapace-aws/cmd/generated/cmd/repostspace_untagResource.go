package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the association of the tag with the AWS re:Post Private resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_untagResourceCmd).Standalone()

	repostspace_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	repostspace_untagResourceCmd.Flags().String("tag-keys", "", "The key values of the tag.")
	repostspace_untagResourceCmd.MarkFlagRequired("resource-arn")
	repostspace_untagResourceCmd.MarkFlagRequired("tag-keys")
	repostspaceCmd.AddCommand(repostspace_untagResourceCmd)
}
