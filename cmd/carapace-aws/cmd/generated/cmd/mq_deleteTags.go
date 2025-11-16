package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_deleteTagsCmd).Standalone()

	mq_deleteTagsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource tag.")
	mq_deleteTagsCmd.Flags().String("tag-keys", "", "An array of tag keys to delete")
	mq_deleteTagsCmd.MarkFlagRequired("resource-arn")
	mq_deleteTagsCmd.MarkFlagRequired("tag-keys")
	mqCmd.AddCommand(mq_deleteTagsCmd)
}
