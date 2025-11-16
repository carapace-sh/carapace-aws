package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateChannelNamespaceCmd = &cobra.Command{
	Use:   "update-channel-namespace",
	Short: "Updates a `ChannelNamespace` associated with an `Api`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateChannelNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_updateChannelNamespaceCmd).Standalone()

		appsync_updateChannelNamespaceCmd.Flags().String("api-id", "", "The `Api` ID.")
		appsync_updateChannelNamespaceCmd.Flags().String("code-handlers", "", "The event handler functions that run custom business logic to process published events and subscribe requests.")
		appsync_updateChannelNamespaceCmd.Flags().String("handler-configs", "", "The configuration for the `OnPublish` and `OnSubscribe` handlers.")
		appsync_updateChannelNamespaceCmd.Flags().String("name", "", "The name of the `ChannelNamespace`.")
		appsync_updateChannelNamespaceCmd.Flags().String("publish-auth-modes", "", "The authorization mode to use for publishing messages on the channel namespace.")
		appsync_updateChannelNamespaceCmd.Flags().String("subscribe-auth-modes", "", "The authorization mode to use for subscribing to messages on the channel namespace.")
		appsync_updateChannelNamespaceCmd.MarkFlagRequired("api-id")
		appsync_updateChannelNamespaceCmd.MarkFlagRequired("name")
	})
	appsyncCmd.AddCommand(appsync_updateChannelNamespaceCmd)
}
