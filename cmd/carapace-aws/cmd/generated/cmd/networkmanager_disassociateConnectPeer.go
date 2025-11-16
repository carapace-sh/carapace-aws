package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_disassociateConnectPeerCmd = &cobra.Command{
	Use:   "disassociate-connect-peer",
	Short: "Disassociates a core network Connect peer from a device and a link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_disassociateConnectPeerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_disassociateConnectPeerCmd).Standalone()

		networkmanager_disassociateConnectPeerCmd.Flags().String("connect-peer-id", "", "The ID of the Connect peer to disassociate from a device.")
		networkmanager_disassociateConnectPeerCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_disassociateConnectPeerCmd.MarkFlagRequired("connect-peer-id")
		networkmanager_disassociateConnectPeerCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_disassociateConnectPeerCmd)
}
