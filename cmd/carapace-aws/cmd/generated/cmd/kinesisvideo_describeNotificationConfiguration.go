package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_describeNotificationConfigurationCmd = &cobra.Command{
	Use:   "describe-notification-configuration",
	Short: "Gets the `NotificationConfiguration` for a given Kinesis video stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_describeNotificationConfigurationCmd).Standalone()

	kinesisvideo_describeNotificationConfigurationCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the Kinesis video stream from where you want to retrieve the notification configuration.")
	kinesisvideo_describeNotificationConfigurationCmd.Flags().String("stream-name", "", "The name of the stream from which to retrieve the notification configuration.")
	kinesisvideoCmd.AddCommand(kinesisvideo_describeNotificationConfigurationCmd)
}
