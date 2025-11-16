package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTransitGatewayMulticastDomainCmd = &cobra.Command{
	Use:   "delete-transit-gateway-multicast-domain",
	Short: "Deletes the specified transit gateway multicast domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTransitGatewayMulticastDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteTransitGatewayMulticastDomainCmd).Standalone()

		ec2_deleteTransitGatewayMulticastDomainCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayMulticastDomainCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteTransitGatewayMulticastDomainCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
		ec2_deleteTransitGatewayMulticastDomainCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteTransitGatewayMulticastDomainCmd.MarkFlagRequired("transit-gateway-multicast-domain-id")
	})
	ec2Cmd.AddCommand(ec2_deleteTransitGatewayMulticastDomainCmd)
}
