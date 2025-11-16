package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateTaskExecutionCmd = &cobra.Command{
	Use:   "update-task-execution",
	Short: "Updates the configuration of a running DataSync task execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateTaskExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_updateTaskExecutionCmd).Standalone()

		datasync_updateTaskExecutionCmd.Flags().String("options", "", "")
		datasync_updateTaskExecutionCmd.Flags().String("task-execution-arn", "", "Specifies the Amazon Resource Name (ARN) of the task execution that you're updating.")
		datasync_updateTaskExecutionCmd.MarkFlagRequired("options")
		datasync_updateTaskExecutionCmd.MarkFlagRequired("task-execution-arn")
	})
	datasyncCmd.AddCommand(datasync_updateTaskExecutionCmd)
}
