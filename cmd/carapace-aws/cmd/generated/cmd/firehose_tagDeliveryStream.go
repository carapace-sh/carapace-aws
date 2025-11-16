package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_tagDeliveryStreamCmd = &cobra.Command{
	Use:   "tag-delivery-stream",
	Short: "Adds or updates tags for the specified Firehose stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_tagDeliveryStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(firehose_tagDeliveryStreamCmd).Standalone()

		firehose_tagDeliveryStreamCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream to which you want to add the tags.")
		firehose_tagDeliveryStreamCmd.Flags().String("tags", "", "A set of key-value pairs to use to create the tags.")
		firehose_tagDeliveryStreamCmd.MarkFlagRequired("delivery-stream-name")
		firehose_tagDeliveryStreamCmd.MarkFlagRequired("tags")
	})
	firehoseCmd.AddCommand(firehose_tagDeliveryStreamCmd)
}
