package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from an index, FAQ, data source, or other resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_untagResourceCmd).Standalone()

	kendra_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the index, FAQ, data source, or other resource to remove a tag.")
	kendra_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the index, FAQ, data source, or other resource.")
	kendra_untagResourceCmd.MarkFlagRequired("resource-arn")
	kendra_untagResourceCmd.MarkFlagRequired("tag-keys")
	kendraCmd.AddCommand(kendra_untagResourceCmd)
}
