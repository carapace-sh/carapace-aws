package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "Adds one or more tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_addTagsToResourceCmd).Standalone()

	storagegateway_addTagsToResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource you want to add tags to.")
	storagegateway_addTagsToResourceCmd.Flags().String("tags", "", "The key-value pair that represents the tag you want to add to the resource.")
	storagegateway_addTagsToResourceCmd.MarkFlagRequired("resource-arn")
	storagegateway_addTagsToResourceCmd.MarkFlagRequired("tags")
	storagegatewayCmd.AddCommand(storagegateway_addTagsToResourceCmd)
}
