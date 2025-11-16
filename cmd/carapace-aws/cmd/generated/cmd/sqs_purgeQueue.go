package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_purgeQueueCmd = &cobra.Command{
	Use:   "purge-queue",
	Short: "Deletes available messages in a queue (including in-flight messages) specified by the `QueueURL` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_purgeQueueCmd).Standalone()

	sqs_purgeQueueCmd.Flags().String("queue-url", "", "The URL of the queue from which the `PurgeQueue` action deletes messages.")
	sqs_purgeQueueCmd.MarkFlagRequired("queue-url")
	sqsCmd.AddCommand(sqs_purgeQueueCmd)
}
