package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd = &cobra.Command{
	Use:   "list-channel-memberships-for-app-instance-user",
	Short: "Lists all channels that an `AppInstanceUser` or `AppInstanceBot` is a part of.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd).Standalone()

		chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd.Flags().String("app-instance-user-arn", "", "The ARN of the user or bot.")
		chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd.Flags().String("max-results", "", "The maximum number of users that you want returned.")
		chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd.Flags().String("next-token", "", "The token returned from previous API requests until the number of channel memberships is reached.")
		chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd.MarkFlagRequired("chime-bearer")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelMembershipsForAppInstanceUserCmd)
}
