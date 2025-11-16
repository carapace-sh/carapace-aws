package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteChannelNamespaceCmd = &cobra.Command{
	Use:   "delete-channel-namespace",
	Short: "Deletes a `ChannelNamespace`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteChannelNamespaceCmd).Standalone()

	appsync_deleteChannelNamespaceCmd.Flags().String("api-id", "", "The ID of the `Api` associated with the `ChannelNamespace`.")
	appsync_deleteChannelNamespaceCmd.Flags().String("name", "", "The name of the `ChannelNamespace`.")
	appsync_deleteChannelNamespaceCmd.MarkFlagRequired("api-id")
	appsync_deleteChannelNamespaceCmd.MarkFlagRequired("name")
	appsyncCmd.AddCommand(appsync_deleteChannelNamespaceCmd)
}
