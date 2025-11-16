package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_listQueueTagsCmd = &cobra.Command{
	Use:   "list-queue-tags",
	Short: "List all cost allocation tags added to the specified Amazon SQS queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_listQueueTagsCmd).Standalone()

	sqs_listQueueTagsCmd.Flags().String("queue-url", "", "The URL of the queue.")
	sqs_listQueueTagsCmd.MarkFlagRequired("queue-url")
	sqsCmd.AddCommand(sqs_listQueueTagsCmd)
}
