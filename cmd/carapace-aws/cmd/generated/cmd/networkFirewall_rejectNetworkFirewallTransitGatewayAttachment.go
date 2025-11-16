package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_rejectNetworkFirewallTransitGatewayAttachmentCmd = &cobra.Command{
	Use:   "reject-network-firewall-transit-gateway-attachment",
	Short: "Rejects a transit gateway attachment request for Network Firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_rejectNetworkFirewallTransitGatewayAttachmentCmd).Standalone()

	networkFirewall_rejectNetworkFirewallTransitGatewayAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "Required.")
	networkFirewall_rejectNetworkFirewallTransitGatewayAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	networkFirewallCmd.AddCommand(networkFirewall_rejectNetworkFirewallTransitGatewayAttachmentCmd)
}
