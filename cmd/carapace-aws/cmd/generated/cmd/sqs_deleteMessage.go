package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_deleteMessageCmd = &cobra.Command{
	Use:   "delete-message",
	Short: "Deletes the specified message from the specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_deleteMessageCmd).Standalone()

	sqs_deleteMessageCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue from which messages are deleted.")
	sqs_deleteMessageCmd.Flags().String("receipt-handle", "", "The receipt handle associated with the message to delete.")
	sqs_deleteMessageCmd.MarkFlagRequired("queue-url")
	sqs_deleteMessageCmd.MarkFlagRequired("receipt-handle")
	sqsCmd.AddCommand(sqs_deleteMessageCmd)
}
