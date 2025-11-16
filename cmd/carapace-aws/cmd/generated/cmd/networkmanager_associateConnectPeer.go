package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_associateConnectPeerCmd = &cobra.Command{
	Use:   "associate-connect-peer",
	Short: "Associates a core network Connect peer with a device and optionally, with a link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_associateConnectPeerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_associateConnectPeerCmd).Standalone()

		networkmanager_associateConnectPeerCmd.Flags().String("connect-peer-id", "", "The ID of the Connect peer.")
		networkmanager_associateConnectPeerCmd.Flags().String("device-id", "", "The ID of the device.")
		networkmanager_associateConnectPeerCmd.Flags().String("global-network-id", "", "The ID of your global network.")
		networkmanager_associateConnectPeerCmd.Flags().String("link-id", "", "The ID of the link.")
		networkmanager_associateConnectPeerCmd.MarkFlagRequired("connect-peer-id")
		networkmanager_associateConnectPeerCmd.MarkFlagRequired("device-id")
		networkmanager_associateConnectPeerCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_associateConnectPeerCmd)
}
