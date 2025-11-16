package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_reportTaskRunnerHeartbeatCmd = &cobra.Command{
	Use:   "report-task-runner-heartbeat",
	Short: "Task runners call `ReportTaskRunnerHeartbeat` every 15 minutes to indicate that they are operational.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_reportTaskRunnerHeartbeatCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_reportTaskRunnerHeartbeatCmd).Standalone()

		datapipeline_reportTaskRunnerHeartbeatCmd.Flags().String("hostname", "", "The public DNS name of the task runner.")
		datapipeline_reportTaskRunnerHeartbeatCmd.Flags().String("taskrunner-id", "", "The ID of the task runner.")
		datapipeline_reportTaskRunnerHeartbeatCmd.Flags().String("worker-group", "", "The type of task the task runner is configured to accept and process.")
		datapipeline_reportTaskRunnerHeartbeatCmd.MarkFlagRequired("taskrunner-id")
	})
	datapipelineCmd.AddCommand(datapipeline_reportTaskRunnerHeartbeatCmd)
}
