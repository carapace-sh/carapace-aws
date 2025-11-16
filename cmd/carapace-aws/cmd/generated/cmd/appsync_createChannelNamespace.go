package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createChannelNamespaceCmd = &cobra.Command{
	Use:   "create-channel-namespace",
	Short: "Creates a `ChannelNamespace` for an `Api`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createChannelNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_createChannelNamespaceCmd).Standalone()

		appsync_createChannelNamespaceCmd.Flags().String("api-id", "", "The `Api` ID.")
		appsync_createChannelNamespaceCmd.Flags().String("code-handlers", "", "The event handler functions that run custom business logic to process published events and subscribe requests.")
		appsync_createChannelNamespaceCmd.Flags().String("handler-configs", "", "The configuration for the `OnPublish` and `OnSubscribe` handlers.")
		appsync_createChannelNamespaceCmd.Flags().String("name", "", "The name of the `ChannelNamespace`.")
		appsync_createChannelNamespaceCmd.Flags().String("publish-auth-modes", "", "The authorization mode to use for publishing messages on the channel namespace.")
		appsync_createChannelNamespaceCmd.Flags().String("subscribe-auth-modes", "", "The authorization mode to use for subscribing to messages on the channel namespace.")
		appsync_createChannelNamespaceCmd.Flags().String("tags", "", "")
		appsync_createChannelNamespaceCmd.MarkFlagRequired("api-id")
		appsync_createChannelNamespaceCmd.MarkFlagRequired("name")
	})
	appsyncCmd.AddCommand(appsync_createChannelNamespaceCmd)
}
