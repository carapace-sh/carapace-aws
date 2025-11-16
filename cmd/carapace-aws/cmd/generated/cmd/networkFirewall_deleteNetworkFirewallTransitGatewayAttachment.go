package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_deleteNetworkFirewallTransitGatewayAttachmentCmd = &cobra.Command{
	Use:   "delete-network-firewall-transit-gateway-attachment",
	Short: "Deletes a transit gateway attachment from a Network Firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_deleteNetworkFirewallTransitGatewayAttachmentCmd).Standalone()

	networkFirewall_deleteNetworkFirewallTransitGatewayAttachmentCmd.Flags().String("transit-gateway-attachment-id", "", "Required.")
	networkFirewall_deleteNetworkFirewallTransitGatewayAttachmentCmd.MarkFlagRequired("transit-gateway-attachment-id")
	networkFirewallCmd.AddCommand(networkFirewall_deleteNetworkFirewallTransitGatewayAttachmentCmd)
}
