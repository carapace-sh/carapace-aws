package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getConnectionsCmd = &cobra.Command{
	Use:   "get-connections",
	Short: "Gets information about one or more of your connections in a global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getConnectionsCmd).Standalone()

		networkmanager_getConnectionsCmd.Flags().String("connection-ids", "", "One or more connection IDs.")
		networkmanager_getConnectionsCmd.Flags().String("device-id", "", "The ID of the device.")
		networkmanager_getConnectionsCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_getConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getConnectionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getConnectionsCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getConnectionsCmd)
}
