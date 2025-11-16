package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_rejectTransitGatewayMulticastDomainAssociationsCmd = &cobra.Command{
	Use:   "reject-transit-gateway-multicast-domain-associations",
	Short: "Rejects a request to associate cross-account subnets with a transit gateway multicast domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_rejectTransitGatewayMulticastDomainAssociationsCmd).Standalone()

	ec2_rejectTransitGatewayMulticastDomainAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_rejectTransitGatewayMulticastDomainAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_rejectTransitGatewayMulticastDomainAssociationsCmd.Flags().String("subnet-ids", "", "The IDs of the subnets to associate with the transit gateway multicast domain.")
	ec2_rejectTransitGatewayMulticastDomainAssociationsCmd.Flags().String("transit-gateway-attachment-id", "", "The ID of the transit gateway attachment.")
	ec2_rejectTransitGatewayMulticastDomainAssociationsCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
	ec2_rejectTransitGatewayMulticastDomainAssociationsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_rejectTransitGatewayMulticastDomainAssociationsCmd)
}
