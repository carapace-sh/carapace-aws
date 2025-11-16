package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getConnectPeerAssociationsCmd = &cobra.Command{
	Use:   "get-connect-peer-associations",
	Short: "Returns information about a core network Connect peer associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getConnectPeerAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getConnectPeerAssociationsCmd).Standalone()

		networkmanager_getConnectPeerAssociationsCmd.Flags().String("connect-peer-ids", "", "The IDs of the Connect peers.")
		networkmanager_getConnectPeerAssociationsCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_getConnectPeerAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getConnectPeerAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getConnectPeerAssociationsCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getConnectPeerAssociationsCmd)
}
