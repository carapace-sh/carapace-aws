package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_changeMessageVisibilityBatchCmd = &cobra.Command{
	Use:   "change-message-visibility-batch",
	Short: "Changes the visibility timeout of multiple messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_changeMessageVisibilityBatchCmd).Standalone()

	sqs_changeMessageVisibilityBatchCmd.Flags().String("entries", "", "Lists the receipt handles of the messages for which the visibility timeout must be changed.")
	sqs_changeMessageVisibilityBatchCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue whose messages' visibility is changed.")
	sqs_changeMessageVisibilityBatchCmd.MarkFlagRequired("entries")
	sqs_changeMessageVisibilityBatchCmd.MarkFlagRequired("queue-url")
	sqsCmd.AddCommand(sqs_changeMessageVisibilityBatchCmd)
}
