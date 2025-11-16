package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayMulticastDomainCmd = &cobra.Command{
	Use:   "create-transit-gateway-multicast-domain",
	Short: "Creates a multicast domain using the specified transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayMulticastDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createTransitGatewayMulticastDomainCmd).Standalone()

		ec2_createTransitGatewayMulticastDomainCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayMulticastDomainCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTransitGatewayMulticastDomainCmd.Flags().String("options", "", "The options for the transit gateway multicast domain.")
		ec2_createTransitGatewayMulticastDomainCmd.Flags().String("tag-specifications", "", "The tags for the transit gateway multicast domain.")
		ec2_createTransitGatewayMulticastDomainCmd.Flags().String("transit-gateway-id", "", "The ID of the transit gateway.")
		ec2_createTransitGatewayMulticastDomainCmd.Flag("no-dry-run").Hidden = true
		ec2_createTransitGatewayMulticastDomainCmd.MarkFlagRequired("transit-gateway-id")
	})
	ec2Cmd.AddCommand(ec2_createTransitGatewayMulticastDomainCmd)
}
