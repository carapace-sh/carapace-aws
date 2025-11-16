package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_startMessageMoveTaskCmd = &cobra.Command{
	Use:   "start-message-move-task",
	Short: "Starts an asynchronous task to move messages from a specified source queue to a specified destination queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_startMessageMoveTaskCmd).Standalone()

	sqs_startMessageMoveTaskCmd.Flags().String("destination-arn", "", "The ARN of the queue that receives the moved messages.")
	sqs_startMessageMoveTaskCmd.Flags().String("max-number-of-messages-per-second", "", "The number of messages to be moved per second (the message movement rate).")
	sqs_startMessageMoveTaskCmd.Flags().String("source-arn", "", "The ARN of the queue that contains the messages to be moved to another queue.")
	sqs_startMessageMoveTaskCmd.MarkFlagRequired("source-arn")
	sqsCmd.AddCommand(sqs_startMessageMoveTaskCmd)
}
