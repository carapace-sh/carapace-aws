package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_sendMessageBatchCmd = &cobra.Command{
	Use:   "send-message-batch",
	Short: "You can use `SendMessageBatch` to send up to 10 messages to the specified queue by assigning either identical or different values to each message (or by not assigning values at all).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_sendMessageBatchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_sendMessageBatchCmd).Standalone()

		sqs_sendMessageBatchCmd.Flags().String("entries", "", "A list of `SendMessageBatchRequestEntry` items.")
		sqs_sendMessageBatchCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue to which batched messages are sent.")
		sqs_sendMessageBatchCmd.MarkFlagRequired("entries")
		sqs_sendMessageBatchCmd.MarkFlagRequired("queue-url")
	})
	sqsCmd.AddCommand(sqs_sendMessageBatchCmd)
}
