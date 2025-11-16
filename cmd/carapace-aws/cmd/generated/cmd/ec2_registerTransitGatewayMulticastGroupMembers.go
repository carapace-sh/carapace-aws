package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_registerTransitGatewayMulticastGroupMembersCmd = &cobra.Command{
	Use:   "register-transit-gateway-multicast-group-members",
	Short: "Registers members (network interfaces) with the transit gateway multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_registerTransitGatewayMulticastGroupMembersCmd).Standalone()

	ec2_registerTransitGatewayMulticastGroupMembersCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_registerTransitGatewayMulticastGroupMembersCmd.Flags().String("group-ip-address", "", "The IP address assigned to the transit gateway multicast group.")
	ec2_registerTransitGatewayMulticastGroupMembersCmd.Flags().String("network-interface-ids", "", "The group members' network interface IDs to register with the transit gateway multicast group.")
	ec2_registerTransitGatewayMulticastGroupMembersCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_registerTransitGatewayMulticastGroupMembersCmd.Flags().String("transit-gateway-multicast-domain-id", "", "The ID of the transit gateway multicast domain.")
	ec2_registerTransitGatewayMulticastGroupMembersCmd.MarkFlagRequired("network-interface-ids")
	ec2_registerTransitGatewayMulticastGroupMembersCmd.Flag("no-dry-run").Hidden = true
	ec2_registerTransitGatewayMulticastGroupMembersCmd.MarkFlagRequired("transit-gateway-multicast-domain-id")
	ec2Cmd.AddCommand(ec2_registerTransitGatewayMulticastGroupMembersCmd)
}
