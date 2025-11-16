package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_removeTagsFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_removeTagsFromResourceCmd).Standalone()

		storagegateway_removeTagsFromResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource you want to remove the tags from.")
		storagegateway_removeTagsFromResourceCmd.Flags().String("tag-keys", "", "The keys of the tags you want to remove from the specified resource.")
		storagegateway_removeTagsFromResourceCmd.MarkFlagRequired("resource-arn")
		storagegateway_removeTagsFromResourceCmd.MarkFlagRequired("tag-keys")
	})
	storagegatewayCmd.AddCommand(storagegateway_removeTagsFromResourceCmd)
}
