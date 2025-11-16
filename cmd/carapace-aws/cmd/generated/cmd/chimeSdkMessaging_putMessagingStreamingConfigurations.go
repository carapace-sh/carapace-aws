package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_putMessagingStreamingConfigurationsCmd = &cobra.Command{
	Use:   "put-messaging-streaming-configurations",
	Short: "Sets the data streaming configuration for an `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_putMessagingStreamingConfigurationsCmd).Standalone()

	chimeSdkMessaging_putMessagingStreamingConfigurationsCmd.Flags().String("app-instance-arn", "", "The ARN of the streaming configuration.")
	chimeSdkMessaging_putMessagingStreamingConfigurationsCmd.Flags().String("streaming-configurations", "", "The streaming configurations.")
	chimeSdkMessaging_putMessagingStreamingConfigurationsCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkMessaging_putMessagingStreamingConfigurationsCmd.MarkFlagRequired("streaming-configurations")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_putMessagingStreamingConfigurationsCmd)
}
