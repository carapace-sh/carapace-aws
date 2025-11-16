package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateTransitGatewayMulticastDomainCmd = &cobra.Command{
	Use:   "disassociate-transit-gateway-multicast-domain",
	Short: "Disassociates the specified subnets from the transit gateway multicast domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateTransitGatewayMulticastDomainCmd).Standalone()

	ec2_disassociateTransitGatewayMulticastDomainCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateTransitGatewayMulticastDomainCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateTransitGatewayMulticastDomainCmd.Flags().String("subnet-ids", "", "The IDs of the subnets;")
	ec2_disassociateTransitGatewayMulticastDomainCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the attachment.")
	ec2_disassociateTransitGatewayMulticastDomainCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
	ec2_disassociateTransitGatewayMulticastDomainCmd.Flag("no-dry-run").Hidden = true
	ec2_disassociateTransitGatewayMulticastDomainCmd.MarkFlagRequired("subnet-ids")
	ec2_disassociateTransitGatewayMulticastDomainCmd.MarkFlagRequired("transit-gateway-attachment-id")
	ec2_disassociateTransitGatewayMulticastDomainCmd.MarkFlagRequired("transit-gateway-multicast-domain-id")
	ec2Cmd.AddCommand(ec2_disassociateTransitGatewayMulticastDomainCmd)
}
