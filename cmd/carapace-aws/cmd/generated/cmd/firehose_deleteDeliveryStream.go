package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_deleteDeliveryStreamCmd = &cobra.Command{
	Use:   "delete-delivery-stream",
	Short: "Deletes a Firehose stream and its data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_deleteDeliveryStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(firehose_deleteDeliveryStreamCmd).Standalone()

		firehose_deleteDeliveryStreamCmd.Flags().String("allow-force-delete", "", "Set this to true if you want to delete the Firehose stream even if Firehose is unable to retire the grant for the CMK.")
		firehose_deleteDeliveryStreamCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream.")
		firehose_deleteDeliveryStreamCmd.MarkFlagRequired("delivery-stream-name")
	})
	firehoseCmd.AddCommand(firehose_deleteDeliveryStreamCmd)
}
