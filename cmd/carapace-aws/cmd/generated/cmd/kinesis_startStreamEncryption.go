package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_startStreamEncryptionCmd = &cobra.Command{
	Use:   "start-stream-encryption",
	Short: "Enables or updates server-side encryption using an Amazon Web Services KMS key for a specified stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_startStreamEncryptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_startStreamEncryptionCmd).Standalone()

		kinesis_startStreamEncryptionCmd.Flags().String("encryption-type", "", "The encryption type to use.")
		kinesis_startStreamEncryptionCmd.Flags().String("key-id", "", "The GUID for the customer-managed Amazon Web Services KMS key to use for encryption.")
		kinesis_startStreamEncryptionCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
		kinesis_startStreamEncryptionCmd.Flags().String("stream-name", "", "The name of the stream for which to start encrypting records.")
		kinesis_startStreamEncryptionCmd.MarkFlagRequired("encryption-type")
		kinesis_startStreamEncryptionCmd.MarkFlagRequired("key-id")
	})
	kinesisCmd.AddCommand(kinesis_startStreamEncryptionCmd)
}
