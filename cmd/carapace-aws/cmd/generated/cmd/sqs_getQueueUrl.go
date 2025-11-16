package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_getQueueUrlCmd = &cobra.Command{
	Use:   "get-queue-url",
	Short: "The `GetQueueUrl` API returns the URL of an existing Amazon SQS queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_getQueueUrlCmd).Standalone()

	sqs_getQueueUrlCmd.Flags().String("queue-name", "", "(Required) The name of the queue for which you want to fetch the URL.")
	sqs_getQueueUrlCmd.Flags().String("queue-owner-awsaccount-id", "", "(Optional) The Amazon Web Services account ID of the account that created the queue.")
	sqs_getQueueUrlCmd.MarkFlagRequired("queue-name")
	sqsCmd.AddCommand(sqs_getQueueUrlCmd)
}
