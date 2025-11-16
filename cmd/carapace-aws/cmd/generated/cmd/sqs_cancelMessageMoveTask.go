package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_cancelMessageMoveTaskCmd = &cobra.Command{
	Use:   "cancel-message-move-task",
	Short: "Cancels a specified message movement task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_cancelMessageMoveTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_cancelMessageMoveTaskCmd).Standalone()

		sqs_cancelMessageMoveTaskCmd.Flags().String("task-handle", "", "An identifier associated with a message movement task.")
		sqs_cancelMessageMoveTaskCmd.MarkFlagRequired("task-handle")
	})
	sqsCmd.AddCommand(sqs_cancelMessageMoveTaskCmd)
}
