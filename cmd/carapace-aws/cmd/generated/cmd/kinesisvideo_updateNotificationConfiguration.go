package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_updateNotificationConfigurationCmd = &cobra.Command{
	Use:   "update-notification-configuration",
	Short: "Updates the notification information for a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_updateNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_updateNotificationConfigurationCmd).Standalone()

		kinesisvideo_updateNotificationConfigurationCmd.Flags().String("notification-configuration", "", "The structure containing the information required for notifications.")
		kinesisvideo_updateNotificationConfigurationCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the Kinesis video stream from where you want to update the notification configuration.")
		kinesisvideo_updateNotificationConfigurationCmd.Flags().String("stream-name", "", "The name of the stream from which to update the notification configuration.")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_updateNotificationConfigurationCmd)
}
