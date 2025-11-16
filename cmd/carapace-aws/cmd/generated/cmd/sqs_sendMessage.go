package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_sendMessageCmd = &cobra.Command{
	Use:   "send-message",
	Short: "Delivers a message to the specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_sendMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_sendMessageCmd).Standalone()

		sqs_sendMessageCmd.Flags().String("delay-seconds", "", "The length of time, in seconds, for which to delay a specific message.")
		sqs_sendMessageCmd.Flags().String("message-attributes", "", "Each message attribute consists of a `Name`, `Type`, and `Value`.")
		sqs_sendMessageCmd.Flags().String("message-body", "", "The message to send.")
		sqs_sendMessageCmd.Flags().String("message-deduplication-id", "", "This parameter applies only to FIFO (first-in-first-out) queues.")
		sqs_sendMessageCmd.Flags().String("message-group-id", "", "`MessageGroupId` is an attribute used in Amazon SQS FIFO (First-In-First-Out) and standard queues.")
		sqs_sendMessageCmd.Flags().String("message-system-attributes", "", "The message system attribute to send.")
		sqs_sendMessageCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue to which a message is sent.")
		sqs_sendMessageCmd.MarkFlagRequired("message-body")
		sqs_sendMessageCmd.MarkFlagRequired("queue-url")
	})
	sqsCmd.AddCommand(sqs_sendMessageCmd)
}
