package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_reportTaskProgressCmd = &cobra.Command{
	Use:   "report-task-progress",
	Short: "Task runners call `ReportTaskProgress` when assigned a task to acknowledge that it has the task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_reportTaskProgressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_reportTaskProgressCmd).Standalone()

		datapipeline_reportTaskProgressCmd.Flags().String("fields", "", "Key-value pairs that define the properties of the ReportTaskProgressInput object.")
		datapipeline_reportTaskProgressCmd.Flags().String("task-id", "", "The ID of the task assigned to the task runner.")
		datapipeline_reportTaskProgressCmd.MarkFlagRequired("task-id")
	})
	datapipelineCmd.AddCommand(datapipeline_reportTaskProgressCmd)
}
