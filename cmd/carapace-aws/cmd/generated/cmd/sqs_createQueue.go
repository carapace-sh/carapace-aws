package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_createQueueCmd = &cobra.Command{
	Use:   "create-queue",
	Short: "Creates a new standard or FIFO queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_createQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_createQueueCmd).Standalone()

		sqs_createQueueCmd.Flags().String("attributes", "", "A map of attributes with their corresponding values.")
		sqs_createQueueCmd.Flags().String("queue-name", "", "The name of the new queue.")
		sqs_createQueueCmd.Flags().String("tags", "", "Add cost allocation tags to the specified Amazon SQS queue.")
		sqs_createQueueCmd.MarkFlagRequired("queue-name")
	})
	sqsCmd.AddCommand(sqs_createQueueCmd)
}
