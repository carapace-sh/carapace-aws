package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getChannelNamespaceCmd = &cobra.Command{
	Use:   "get-channel-namespace",
	Short: "Retrieves the channel namespace for a specified `Api`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getChannelNamespaceCmd).Standalone()

	appsync_getChannelNamespaceCmd.Flags().String("api-id", "", "The `Api` ID.")
	appsync_getChannelNamespaceCmd.Flags().String("name", "", "The name of the `ChannelNamespace`.")
	appsync_getChannelNamespaceCmd.MarkFlagRequired("api-id")
	appsync_getChannelNamespaceCmd.MarkFlagRequired("name")
	appsyncCmd.AddCommand(appsync_getChannelNamespaceCmd)
}
