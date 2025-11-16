package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Lists tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_listTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_listTagsCmd).Standalone()

		mq_listTagsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource tag.")
		mq_listTagsCmd.MarkFlagRequired("resource-arn")
	})
	mqCmd.AddCommand(mq_listTagsCmd)
}
