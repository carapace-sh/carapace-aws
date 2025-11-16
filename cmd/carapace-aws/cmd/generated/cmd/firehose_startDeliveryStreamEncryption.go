package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_startDeliveryStreamEncryptionCmd = &cobra.Command{
	Use:   "start-delivery-stream-encryption",
	Short: "Enables server-side encryption (SSE) for the Firehose stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_startDeliveryStreamEncryptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(firehose_startDeliveryStreamEncryptionCmd).Standalone()

		firehose_startDeliveryStreamEncryptionCmd.Flags().String("delivery-stream-encryption-configuration-input", "", "Used to specify the type and Amazon Resource Name (ARN) of the KMS key needed for Server-Side Encryption (SSE).")
		firehose_startDeliveryStreamEncryptionCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream for which you want to enable server-side encryption (SSE).")
		firehose_startDeliveryStreamEncryptionCmd.MarkFlagRequired("delivery-stream-name")
	})
	firehoseCmd.AddCommand(firehose_startDeliveryStreamEncryptionCmd)
}
