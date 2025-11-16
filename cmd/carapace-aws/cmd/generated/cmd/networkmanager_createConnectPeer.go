package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createConnectPeerCmd = &cobra.Command{
	Use:   "create-connect-peer",
	Short: "Creates a core network Connect peer for a specified core network connect attachment between a core network and an appliance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createConnectPeerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createConnectPeerCmd).Standalone()

		networkmanager_createConnectPeerCmd.Flags().String("bgp-options", "", "The Connect peer BGP options.")
		networkmanager_createConnectPeerCmd.Flags().String("client-token", "", "The client token associated with the request.")
		networkmanager_createConnectPeerCmd.Flags().String("connect-attachment-id", "", "The ID of the connection attachment.")
		networkmanager_createConnectPeerCmd.Flags().String("core-network-address", "", "A Connect peer core network address.")
		networkmanager_createConnectPeerCmd.Flags().String("inside-cidr-blocks", "", "The inside IP addresses used for BGP peering.")
		networkmanager_createConnectPeerCmd.Flags().String("peer-address", "", "The Connect peer address.")
		networkmanager_createConnectPeerCmd.Flags().String("subnet-arn", "", "The subnet ARN for the Connect peer.")
		networkmanager_createConnectPeerCmd.Flags().String("tags", "", "The tags associated with the peer request.")
		networkmanager_createConnectPeerCmd.MarkFlagRequired("connect-attachment-id")
		networkmanager_createConnectPeerCmd.MarkFlagRequired("peer-address")
	})
	networkmanagerCmd.AddCommand(networkmanager_createConnectPeerCmd)
}
