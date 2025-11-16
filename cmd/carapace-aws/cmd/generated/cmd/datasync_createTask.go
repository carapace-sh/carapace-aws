package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createTaskCmd = &cobra.Command{
	Use:   "create-task",
	Short: "Configures a *task*, which defines where and how DataSync transfers your data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_createTaskCmd).Standalone()

		datasync_createTaskCmd.Flags().String("cloud-watch-log-group-arn", "", "Specifies the Amazon Resource Name (ARN) of an Amazon CloudWatch log group for monitoring your task.")
		datasync_createTaskCmd.Flags().String("destination-location-arn", "", "Specifies the ARN of your transfer's destination location.")
		datasync_createTaskCmd.Flags().String("excludes", "", "Specifies exclude filters that define the files, objects, and folders in your source location that you don't want DataSync to transfer.")
		datasync_createTaskCmd.Flags().String("includes", "", "Specifies include filters that define the files, objects, and folders in your source location that you want DataSync to transfer.")
		datasync_createTaskCmd.Flags().String("manifest-config", "", "Configures a manifest, which is a list of files or objects that you want DataSync to transfer.")
		datasync_createTaskCmd.Flags().String("name", "", "Specifies the name of your task.")
		datasync_createTaskCmd.Flags().String("options", "", "Specifies your task's settings, such as preserving file metadata, verifying data integrity, among other options.")
		datasync_createTaskCmd.Flags().String("schedule", "", "Specifies a schedule for when you want your task to run.")
		datasync_createTaskCmd.Flags().String("source-location-arn", "", "Specifies the ARN of your transfer's source location.")
		datasync_createTaskCmd.Flags().String("tags", "", "Specifies the tags that you want to apply to your task.")
		datasync_createTaskCmd.Flags().String("task-mode", "", "Specifies one of the following task modes for your data transfer:")
		datasync_createTaskCmd.Flags().String("task-report-config", "", "Specifies how you want to configure a task report, which provides detailed information about your DataSync transfer.")
		datasync_createTaskCmd.MarkFlagRequired("destination-location-arn")
		datasync_createTaskCmd.MarkFlagRequired("source-location-arn")
	})
	datasyncCmd.AddCommand(datasync_createTaskCmd)
}
