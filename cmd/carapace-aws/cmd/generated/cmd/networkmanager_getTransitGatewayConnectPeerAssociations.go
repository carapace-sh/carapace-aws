package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getTransitGatewayConnectPeerAssociationsCmd = &cobra.Command{
	Use:   "get-transit-gateway-connect-peer-associations",
	Short: "Gets information about one or more of your transit gateway Connect peer associations in a global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getTransitGatewayConnectPeerAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getTransitGatewayConnectPeerAssociationsCmd).Standalone()

		networkmanager_getTransitGatewayConnectPeerAssociationsCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_getTransitGatewayConnectPeerAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getTransitGatewayConnectPeerAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getTransitGatewayConnectPeerAssociationsCmd.Flags().String("transit-gateway-connect-peer-arns", "", "One or more transit gateway Connect peer Amazon Resource Names (ARNs).")
		networkmanager_getTransitGatewayConnectPeerAssociationsCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getTransitGatewayConnectPeerAssociationsCmd)
}
