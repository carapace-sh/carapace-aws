package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_untagDeliveryStreamCmd = &cobra.Command{
	Use:   "untag-delivery-stream",
	Short: "Removes tags from the specified Firehose stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_untagDeliveryStreamCmd).Standalone()

	firehose_untagDeliveryStreamCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream.")
	firehose_untagDeliveryStreamCmd.Flags().String("tag-keys", "", "A list of tag keys.")
	firehose_untagDeliveryStreamCmd.MarkFlagRequired("delivery-stream-name")
	firehose_untagDeliveryStreamCmd.MarkFlagRequired("tag-keys")
	firehoseCmd.AddCommand(firehose_untagDeliveryStreamCmd)
}
