package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_deleteMessagingStreamingConfigurationsCmd = &cobra.Command{
	Use:   "delete-messaging-streaming-configurations",
	Short: "Deletes the streaming configurations for an `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_deleteMessagingStreamingConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_deleteMessagingStreamingConfigurationsCmd).Standalone()

		chimeSdkMessaging_deleteMessagingStreamingConfigurationsCmd.Flags().String("app-instance-arn", "", "The ARN of the streaming configurations being deleted.")
		chimeSdkMessaging_deleteMessagingStreamingConfigurationsCmd.MarkFlagRequired("app-instance-arn")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_deleteMessagingStreamingConfigurationsCmd)
}
