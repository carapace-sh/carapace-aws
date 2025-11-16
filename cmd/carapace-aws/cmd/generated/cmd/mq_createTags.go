package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_createTagsCmd = &cobra.Command{
	Use:   "create-tags",
	Short: "Add a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_createTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_createTagsCmd).Standalone()

		mq_createTagsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource tag.")
		mq_createTagsCmd.Flags().String("tags", "", "The key-value pair for the resource tag.")
		mq_createTagsCmd.MarkFlagRequired("resource-arn")
	})
	mqCmd.AddCommand(mq_createTagsCmd)
}
