package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_stopDeliveryStreamEncryptionCmd = &cobra.Command{
	Use:   "stop-delivery-stream-encryption",
	Short: "Disables server-side encryption (SSE) for the Firehose stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_stopDeliveryStreamEncryptionCmd).Standalone()

	firehose_stopDeliveryStreamEncryptionCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream for which you want to disable server-side encryption (SSE).")
	firehose_stopDeliveryStreamEncryptionCmd.MarkFlagRequired("delivery-stream-name")
	firehoseCmd.AddCommand(firehose_stopDeliveryStreamEncryptionCmd)
}
