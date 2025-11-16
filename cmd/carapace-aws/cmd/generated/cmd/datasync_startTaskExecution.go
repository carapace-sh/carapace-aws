package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_startTaskExecutionCmd = &cobra.Command{
	Use:   "start-task-execution",
	Short: "Starts an DataSync transfer task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_startTaskExecutionCmd).Standalone()

	datasync_startTaskExecutionCmd.Flags().String("excludes", "", "Specifies a list of filter rules that determines which files to exclude from a task.")
	datasync_startTaskExecutionCmd.Flags().String("includes", "", "Specifies a list of filter rules that determines which files to include when running a task.")
	datasync_startTaskExecutionCmd.Flags().String("manifest-config", "", "Configures a manifest, which is a list of files or objects that you want DataSync to transfer.")
	datasync_startTaskExecutionCmd.Flags().String("override-options", "", "")
	datasync_startTaskExecutionCmd.Flags().String("tags", "", "Specifies the tags that you want to apply to the Amazon Resource Name (ARN) representing the task execution.")
	datasync_startTaskExecutionCmd.Flags().String("task-arn", "", "Specifies the Amazon Resource Name (ARN) of the task that you want to start.")
	datasync_startTaskExecutionCmd.Flags().String("task-report-config", "", "Specifies how you want to configure a task report, which provides detailed information about your DataSync transfer.")
	datasync_startTaskExecutionCmd.MarkFlagRequired("task-arn")
	datasyncCmd.AddCommand(datasync_startTaskExecutionCmd)
}
