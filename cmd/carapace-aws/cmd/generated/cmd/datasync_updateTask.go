package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateTaskCmd = &cobra.Command{
	Use:   "update-task",
	Short: "Updates the configuration of a *task*, which defines where and how DataSync transfers your data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateTaskCmd).Standalone()

	datasync_updateTaskCmd.Flags().String("cloud-watch-log-group-arn", "", "Specifies the Amazon Resource Name (ARN) of an Amazon CloudWatch log group for monitoring your task.")
	datasync_updateTaskCmd.Flags().String("excludes", "", "Specifies exclude filters that define the files, objects, and folders in your source location that you don't want DataSync to transfer.")
	datasync_updateTaskCmd.Flags().String("includes", "", "Specifies include filters define the files, objects, and folders in your source location that you want DataSync to transfer.")
	datasync_updateTaskCmd.Flags().String("manifest-config", "", "Configures a manifest, which is a list of files or objects that you want DataSync to transfer.")
	datasync_updateTaskCmd.Flags().String("name", "", "Specifies the name of your task.")
	datasync_updateTaskCmd.Flags().String("options", "", "")
	datasync_updateTaskCmd.Flags().String("schedule", "", "Specifies a schedule for when you want your task to run.")
	datasync_updateTaskCmd.Flags().String("task-arn", "", "Specifies the ARN of the task that you want to update.")
	datasync_updateTaskCmd.Flags().String("task-report-config", "", "Specifies how you want to configure a task report, which provides detailed information about your DataSync transfer.")
	datasync_updateTaskCmd.MarkFlagRequired("task-arn")
	datasyncCmd.AddCommand(datasync_updateTaskCmd)
}
