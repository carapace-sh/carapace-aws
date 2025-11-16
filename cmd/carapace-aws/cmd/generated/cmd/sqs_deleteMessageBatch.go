package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_deleteMessageBatchCmd = &cobra.Command{
	Use:   "delete-message-batch",
	Short: "Deletes up to ten messages from the specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_deleteMessageBatchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_deleteMessageBatchCmd).Standalone()

		sqs_deleteMessageBatchCmd.Flags().String("entries", "", "Lists the receipt handles for the messages to be deleted.")
		sqs_deleteMessageBatchCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue from which messages are deleted.")
		sqs_deleteMessageBatchCmd.MarkFlagRequired("entries")
		sqs_deleteMessageBatchCmd.MarkFlagRequired("queue-url")
	})
	sqsCmd.AddCommand(sqs_deleteMessageBatchCmd)
}
