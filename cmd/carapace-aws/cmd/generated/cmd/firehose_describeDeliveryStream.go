package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_describeDeliveryStreamCmd = &cobra.Command{
	Use:   "describe-delivery-stream",
	Short: "Describes the specified Firehose stream and its status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_describeDeliveryStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(firehose_describeDeliveryStreamCmd).Standalone()

		firehose_describeDeliveryStreamCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream.")
		firehose_describeDeliveryStreamCmd.Flags().String("exclusive-start-destination-id", "", "The ID of the destination to start returning the destination information.")
		firehose_describeDeliveryStreamCmd.Flags().String("limit", "", "The limit on the number of destinations to return.")
		firehose_describeDeliveryStreamCmd.MarkFlagRequired("delivery-stream-name")
	})
	firehoseCmd.AddCommand(firehose_describeDeliveryStreamCmd)
}
