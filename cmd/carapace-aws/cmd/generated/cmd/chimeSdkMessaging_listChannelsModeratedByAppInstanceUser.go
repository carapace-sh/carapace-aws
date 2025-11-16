package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelsModeratedByAppInstanceUserCmd = &cobra.Command{
	Use:   "list-channels-moderated-by-app-instance-user",
	Short: "A list of the channels moderated by an `AppInstanceUser`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelsModeratedByAppInstanceUserCmd).Standalone()

	chimeSdkMessaging_listChannelsModeratedByAppInstanceUserCmd.Flags().String("app-instance-user-arn", "", "The ARN of the user or bot in the moderated channel.")
	chimeSdkMessaging_listChannelsModeratedByAppInstanceUserCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_listChannelsModeratedByAppInstanceUserCmd.Flags().String("max-results", "", "The maximum number of channels in the request.")
	chimeSdkMessaging_listChannelsModeratedByAppInstanceUserCmd.Flags().String("next-token", "", "The token returned from previous API requests until the number of channels moderated by the user is reached.")
	chimeSdkMessaging_listChannelsModeratedByAppInstanceUserCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelsModeratedByAppInstanceUserCmd)
}
