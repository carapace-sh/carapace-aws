package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_getMessagingSessionEndpointCmd = &cobra.Command{
	Use:   "get-messaging-session-endpoint",
	Short: "The details of the endpoint for the messaging session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_getMessagingSessionEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_getMessagingSessionEndpointCmd).Standalone()

		chimeSdkMessaging_getMessagingSessionEndpointCmd.Flags().String("network-type", "", "The type of network for the messaging session endpoint.")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_getMessagingSessionEndpointCmd)
}
