package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_receiveMessageCmd = &cobra.Command{
	Use:   "receive-message",
	Short: "Retrieves one or more messages (up to 10), from the specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_receiveMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_receiveMessageCmd).Standalone()

		sqs_receiveMessageCmd.Flags().String("attribute-names", "", "This parameter has been discontinued but will be supported for backward compatibility.")
		sqs_receiveMessageCmd.Flags().String("max-number-of-messages", "", "The maximum number of messages to return.")
		sqs_receiveMessageCmd.Flags().String("message-attribute-names", "", "The name of the message attribute, where *N* is the index.")
		sqs_receiveMessageCmd.Flags().String("message-system-attribute-names", "", "A list of attributes that need to be returned along with each message.")
		sqs_receiveMessageCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue from which messages are received.")
		sqs_receiveMessageCmd.Flags().String("receive-request-attempt-id", "", "This parameter applies only to FIFO (first-in-first-out) queues.")
		sqs_receiveMessageCmd.Flags().String("visibility-timeout", "", "The duration (in seconds) that the received messages are hidden from subsequent retrieve requests after being retrieved by a `ReceiveMessage` request.")
		sqs_receiveMessageCmd.Flags().String("wait-time-seconds", "", "The duration (in seconds) for which the call waits for a message to arrive in the queue before returning.")
		sqs_receiveMessageCmd.MarkFlagRequired("queue-url")
	})
	sqsCmd.AddCommand(sqs_receiveMessageCmd)
}
