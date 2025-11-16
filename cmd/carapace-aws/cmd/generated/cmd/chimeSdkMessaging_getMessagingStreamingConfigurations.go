package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_getMessagingStreamingConfigurationsCmd = &cobra.Command{
	Use:   "get-messaging-streaming-configurations",
	Short: "Retrieves the data streaming configuration for an `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_getMessagingStreamingConfigurationsCmd).Standalone()

	chimeSdkMessaging_getMessagingStreamingConfigurationsCmd.Flags().String("app-instance-arn", "", "The ARN of the streaming configurations.")
	chimeSdkMessaging_getMessagingStreamingConfigurationsCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_getMessagingStreamingConfigurationsCmd)
}
