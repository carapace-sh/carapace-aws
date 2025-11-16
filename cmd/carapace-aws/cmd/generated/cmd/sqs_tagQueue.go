package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_tagQueueCmd = &cobra.Command{
	Use:   "tag-queue",
	Short: "Add cost allocation tags to the specified Amazon SQS queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_tagQueueCmd).Standalone()

	sqs_tagQueueCmd.Flags().String("queue-url", "", "The URL of the queue.")
	sqs_tagQueueCmd.Flags().String("tags", "", "The list of tags to be added to the specified queue.")
	sqs_tagQueueCmd.MarkFlagRequired("queue-url")
	sqs_tagQueueCmd.MarkFlagRequired("tags")
	sqsCmd.AddCommand(sqs_tagQueueCmd)
}
