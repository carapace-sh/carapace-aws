package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listChannelNamespacesCmd = &cobra.Command{
	Use:   "list-channel-namespaces",
	Short: "Lists the channel namespaces for a specified `Api`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listChannelNamespacesCmd).Standalone()

	appsync_listChannelNamespacesCmd.Flags().String("api-id", "", "The `Api` ID.")
	appsync_listChannelNamespacesCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
	appsync_listChannelNamespacesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
	appsync_listChannelNamespacesCmd.MarkFlagRequired("api-id")
	appsyncCmd.AddCommand(appsync_listChannelNamespacesCmd)
}
