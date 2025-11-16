package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getConnectPeerCmd = &cobra.Command{
	Use:   "get-connect-peer",
	Short: "Returns information about a core network Connect peer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getConnectPeerCmd).Standalone()

	networkmanager_getConnectPeerCmd.Flags().String("connect-peer-id", "", "The ID of the Connect peer.")
	networkmanager_getConnectPeerCmd.MarkFlagRequired("connect-peer-id")
	networkmanagerCmd.AddCommand(networkmanager_getConnectPeerCmd)
}
