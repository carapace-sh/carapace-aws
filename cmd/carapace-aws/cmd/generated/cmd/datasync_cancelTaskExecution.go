package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_cancelTaskExecutionCmd = &cobra.Command{
	Use:   "cancel-task-execution",
	Short: "Stops an DataSync task execution that's in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_cancelTaskExecutionCmd).Standalone()

	datasync_cancelTaskExecutionCmd.Flags().String("task-execution-arn", "", "The Amazon Resource Name (ARN) of the task execution to stop.")
	datasync_cancelTaskExecutionCmd.MarkFlagRequired("task-execution-arn")
	datasyncCmd.AddCommand(datasync_cancelTaskExecutionCmd)
}
