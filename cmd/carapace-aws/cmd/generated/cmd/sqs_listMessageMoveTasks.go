package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_listMessageMoveTasksCmd = &cobra.Command{
	Use:   "list-message-move-tasks",
	Short: "Gets the most recent message movement tasks (up to 10) under a specific source queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_listMessageMoveTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_listMessageMoveTasksCmd).Standalone()

		sqs_listMessageMoveTasksCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		sqs_listMessageMoveTasksCmd.Flags().String("source-arn", "", "The ARN of the queue whose message movement tasks are to be listed.")
		sqs_listMessageMoveTasksCmd.MarkFlagRequired("source-arn")
	})
	sqsCmd.AddCommand(sqs_listMessageMoveTasksCmd)
}
