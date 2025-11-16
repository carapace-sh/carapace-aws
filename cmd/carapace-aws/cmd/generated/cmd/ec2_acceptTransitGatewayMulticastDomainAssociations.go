package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_acceptTransitGatewayMulticastDomainAssociationsCmd = &cobra.Command{
	Use:   "accept-transit-gateway-multicast-domain-associations",
	Short: "Accepts a request to associate subnets with a transit gateway multicast domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_acceptTransitGatewayMulticastDomainAssociationsCmd).Standalone()

	ec2_acceptTransitGatewayMulticastDomainAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptTransitGatewayMulticastDomainAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_acceptTransitGatewayMulticastDomainAssociationsCmd.Flags().String("subnet-ids", "", "The IDs of the subnets to associate with the transit gateway multicast domain.")
	ec2_acceptTransitGatewayMulticastDomainAssociationsCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the transit gateway attachment.")
	ec2_acceptTransitGatewayMulticastDomainAssociationsCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
	ec2_acceptTransitGatewayMulticastDomainAssociationsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_acceptTransitGatewayMulticastDomainAssociationsCmd)
}
