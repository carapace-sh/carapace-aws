package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_stopStreamEncryptionCmd = &cobra.Command{
	Use:   "stop-stream-encryption",
	Short: "Disables server-side encryption for a specified stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_stopStreamEncryptionCmd).Standalone()

	kinesis_stopStreamEncryptionCmd.Flags().String("encryption-type", "", "The encryption type.")
	kinesis_stopStreamEncryptionCmd.Flags().String("key-id", "", "The GUID for the customer-managed Amazon Web Services KMS key to use for encryption.")
	kinesis_stopStreamEncryptionCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_stopStreamEncryptionCmd.Flags().String("stream-name", "", "The name of the stream on which to stop encrypting records.")
	kinesis_stopStreamEncryptionCmd.MarkFlagRequired("encryption-type")
	kinesis_stopStreamEncryptionCmd.MarkFlagRequired("key-id")
	kinesisCmd.AddCommand(kinesis_stopStreamEncryptionCmd)
}
