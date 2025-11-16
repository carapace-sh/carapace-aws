package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_untagQueueCmd = &cobra.Command{
	Use:   "untag-queue",
	Short: "Remove cost allocation tags from the specified Amazon SQS queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_untagQueueCmd).Standalone()

	sqs_untagQueueCmd.Flags().String("queue-url", "", "The URL of the queue.")
	sqs_untagQueueCmd.Flags().String("tag-keys", "", "The list of tags to be removed from the specified queue.")
	sqs_untagQueueCmd.MarkFlagRequired("queue-url")
	sqs_untagQueueCmd.MarkFlagRequired("tag-keys")
	sqsCmd.AddCommand(sqs_untagQueueCmd)
}
