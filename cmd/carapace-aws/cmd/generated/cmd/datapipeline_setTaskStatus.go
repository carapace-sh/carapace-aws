package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_setTaskStatusCmd = &cobra.Command{
	Use:   "set-task-status",
	Short: "Task runners call `SetTaskStatus` to notify AWS Data Pipeline that a task is completed and provide information about the final status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_setTaskStatusCmd).Standalone()

	datapipeline_setTaskStatusCmd.Flags().String("error-id", "", "If an error occurred during the task, this value specifies the error code.")
	datapipeline_setTaskStatusCmd.Flags().String("error-message", "", "If an error occurred during the task, this value specifies a text description of the error.")
	datapipeline_setTaskStatusCmd.Flags().String("error-stack-trace", "", "If an error occurred during the task, this value specifies the stack trace associated with the error.")
	datapipeline_setTaskStatusCmd.Flags().String("task-id", "", "The ID of the task assigned to the task runner.")
	datapipeline_setTaskStatusCmd.Flags().String("task-status", "", "If `FINISHED`, the task successfully completed.")
	datapipeline_setTaskStatusCmd.MarkFlagRequired("task-id")
	datapipeline_setTaskStatusCmd.MarkFlagRequired("task-status")
	datapipelineCmd.AddCommand(datapipeline_setTaskStatusCmd)
}
