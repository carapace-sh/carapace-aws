package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_deleteQueueCmd = &cobra.Command{
	Use:   "delete-queue",
	Short: "Deletes the queue specified by the `QueueUrl`, regardless of the queue's contents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_deleteQueueCmd).Standalone()

	sqs_deleteQueueCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue to delete.")
	sqs_deleteQueueCmd.MarkFlagRequired("queue-url")
	sqsCmd.AddCommand(sqs_deleteQueueCmd)
}
