package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_updatePipelineNotificationsCmd = &cobra.Command{
	Use:   "update-pipeline-notifications",
	Short: "With the UpdatePipelineNotifications operation, you can update Amazon Simple Notification Service (Amazon SNS) notifications for a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_updatePipelineNotificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_updatePipelineNotificationsCmd).Standalone()

		elastictranscoder_updatePipelineNotificationsCmd.Flags().String("id", "", "The identifier of the pipeline for which you want to change notification settings.")
		elastictranscoder_updatePipelineNotificationsCmd.Flags().String("notifications", "", "The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status.")
		elastictranscoder_updatePipelineNotificationsCmd.MarkFlagRequired("id")
		elastictranscoder_updatePipelineNotificationsCmd.MarkFlagRequired("notifications")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_updatePipelineNotificationsCmd)
}
