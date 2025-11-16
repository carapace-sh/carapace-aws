package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_changeMessageVisibilityCmd = &cobra.Command{
	Use:   "change-message-visibility",
	Short: "Changes the visibility timeout of a specified message in a queue to a new value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_changeMessageVisibilityCmd).Standalone()

	sqs_changeMessageVisibilityCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue whose message's visibility is changed.")
	sqs_changeMessageVisibilityCmd.Flags().String("receipt-handle", "", "The receipt handle associated with the message, whose visibility timeout is changed.")
	sqs_changeMessageVisibilityCmd.Flags().String("visibility-timeout", "", "The new value for the message's visibility timeout (in seconds).")
	sqs_changeMessageVisibilityCmd.MarkFlagRequired("queue-url")
	sqs_changeMessageVisibilityCmd.MarkFlagRequired("receipt-handle")
	sqs_changeMessageVisibilityCmd.MarkFlagRequired("visibility-timeout")
	sqsCmd.AddCommand(sqs_changeMessageVisibilityCmd)
}
