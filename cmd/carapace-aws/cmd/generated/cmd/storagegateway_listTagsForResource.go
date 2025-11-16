package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags that have been added to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_listTagsForResourceCmd).Standalone()

		storagegateway_listTagsForResourceCmd.Flags().String("limit", "", "Specifies that the list of tags returned be limited to the specified number of items.")
		storagegateway_listTagsForResourceCmd.Flags().String("marker", "", "An opaque string that indicates the position at which to begin returning the list of tags.")
		storagegateway_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to list tags.")
		storagegateway_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_listTagsForResourceCmd)
}
