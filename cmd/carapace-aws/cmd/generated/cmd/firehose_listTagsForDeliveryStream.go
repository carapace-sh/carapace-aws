package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_listTagsForDeliveryStreamCmd = &cobra.Command{
	Use:   "list-tags-for-delivery-stream",
	Short: "Lists the tags for the specified Firehose stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_listTagsForDeliveryStreamCmd).Standalone()

	firehose_listTagsForDeliveryStreamCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream whose tags you want to list.")
	firehose_listTagsForDeliveryStreamCmd.Flags().String("exclusive-start-tag-key", "", "The key to use as the starting point for the list of tags.")
	firehose_listTagsForDeliveryStreamCmd.Flags().String("limit", "", "The number of tags to return.")
	firehose_listTagsForDeliveryStreamCmd.MarkFlagRequired("delivery-stream-name")
	firehoseCmd.AddCommand(firehose_listTagsForDeliveryStreamCmd)
}
