package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deregisterTransitGatewayMulticastGroupMembersCmd = &cobra.Command{
	Use:   "deregister-transit-gateway-multicast-group-members",
	Short: "Deregisters the specified members (network interfaces) from the transit gateway multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deregisterTransitGatewayMulticastGroupMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deregisterTransitGatewayMulticastGroupMembersCmd).Standalone()

		ec2_deregisterTransitGatewayMulticastGroupMembersCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deregisterTransitGatewayMulticastGroupMembersCmd.Flags().String("group-ip-address", "", "The IP address assigned to the transit gateway multicast group.")
		ec2_deregisterTransitGatewayMulticastGroupMembersCmd.Flags().String("network-interface-ids", "", "The IDs of the group members' network interfaces.")
		ec2_deregisterTransitGatewayMulticastGroupMembersCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deregisterTransitGatewayMulticastGroupMembersCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
		ec2_deregisterTransitGatewayMulticastGroupMembersCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deregisterTransitGatewayMulticastGroupMembersCmd)
}
