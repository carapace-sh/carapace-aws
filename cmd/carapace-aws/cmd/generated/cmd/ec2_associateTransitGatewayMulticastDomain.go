package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateTransitGatewayMulticastDomainCmd = &cobra.Command{
	Use:   "associate-transit-gateway-multicast-domain",
	Short: "Associates the specified subnets and transit gateway attachments with the specified transit gateway multicast domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateTransitGatewayMulticastDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_associateTransitGatewayMulticastDomainCmd).Standalone()

		ec2_associateTransitGatewayMulticastDomainCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateTransitGatewayMulticastDomainCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateTransitGatewayMulticastDomainCmd.Flags().String("subnet-ids", "", "The IDs of the subnets to associate with the transit gateway multicast domain.")
		ec2_associateTransitGatewayMulticastDomainCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the transit gateway attachment to associate with the transit gateway multicast domain.")
		ec2_associateTransitGatewayMulticastDomainCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
		ec2_associateTransitGatewayMulticastDomainCmd.Flag("no-dry-run").Hidden = true
		ec2_associateTransitGatewayMulticastDomainCmd.MarkFlagRequired("subnet-ids")
		ec2_associateTransitGatewayMulticastDomainCmd.MarkFlagRequired("transit-gateway-attachment-id")
		ec2_associateTransitGatewayMulticastDomainCmd.MarkFlagRequired("transit-gateway-multicast-domain-id")
	})
	ec2Cmd.AddCommand(ec2_associateTransitGatewayMulticastDomainCmd)
}
