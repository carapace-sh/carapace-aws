package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_acceptNetworkFirewallTransitGatewayAttachmentCmd = &cobra.Command{
	Use:   "accept-network-firewall-transit-gateway-attachment",
	Short: "Accepts a transit gateway attachment request for Network Firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_acceptNetworkFirewallTransitGatewayAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_acceptNetworkFirewallTransitGatewayAttachmentCmd).Standalone()

		networkFirewall_acceptNetworkFirewallTransitGatewayAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "Required.")
		networkFirewall_acceptNetworkFirewallTransitGatewayAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	})
	networkFirewallCmd.AddCommand(networkFirewall_acceptNetworkFirewallTransitGatewayAttachmentCmd)
}
