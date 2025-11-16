package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getTransitGatewayMulticastDomainAssociationsCmd = &cobra.Command{
	Use:   "get-transit-gateway-multicast-domain-associations",
	Short: "Gets information about the associations for the transit gateway multicast domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getTransitGatewayMulticastDomainAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getTransitGatewayMulticastDomainAssociationsCmd).Standalone()

		ec2_getTransitGatewayMulticastDomainAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayMulticastDomainAssociationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_getTransitGatewayMulticastDomainAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_getTransitGatewayMulticastDomainAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getTransitGatewayMulticastDomainAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getTransitGatewayMulticastDomainAssociationsCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
		ec2_getTransitGatewayMulticastDomainAssociationsCmd.Flag("no-dry-run").Hidden = true
		ec2_getTransitGatewayMulticastDomainAssociationsCmd.MarkFlagRequired("transit-gateway-multicast-domain-id")
	})
	ec2Cmd.AddCommand(ec2_getTransitGatewayMulticastDomainAssociationsCmd)
}
