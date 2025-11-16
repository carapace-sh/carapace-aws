package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_pollForTaskCmd = &cobra.Command{
	Use:   "poll-for-task",
	Short: "Task runners call `PollForTask` to receive a task to perform from AWS Data Pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_pollForTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_pollForTaskCmd).Standalone()

		datapipeline_pollForTaskCmd.Flags().String("hostname", "", "The public DNS name of the calling task runner.")
		datapipeline_pollForTaskCmd.Flags().String("instance-identity", "", "Identity information for the EC2 instance that is hosting the task runner.")
		datapipeline_pollForTaskCmd.Flags().String("worker-group", "", "The type of task the task runner is configured to accept and process.")
		datapipeline_pollForTaskCmd.MarkFlagRequired("worker-group")
	})
	datapipelineCmd.AddCommand(datapipeline_pollForTaskCmd)
}
