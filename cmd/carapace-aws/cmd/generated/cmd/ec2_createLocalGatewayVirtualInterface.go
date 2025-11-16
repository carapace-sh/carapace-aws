package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createLocalGatewayVirtualInterfaceCmd = &cobra.Command{
	Use:   "create-local-gateway-virtual-interface",
	Short: "Create a virtual interface for a local gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createLocalGatewayVirtualInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createLocalGatewayVirtualInterfaceCmd).Standalone()

		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().String("local-address", "", "The IP address assigned to the local gateway virtual interface on the Outpost side.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().String("local-gateway-virtual-interface-group-id", "", "The ID of the local gateway virtual interface group.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().String("outpost-lag-id", "", "References the Link Aggregation Group (LAG) that connects the Outpost to on-premises network devices.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().String("peer-address", "", "The peer IP address for the local gateway virtual interface.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().String("peer-bgp-asn", "", "The Autonomous System Number (ASN) of the Border Gateway Protocol (BGP) peer.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().String("peer-bgp-asn-extended", "", "The extended 32-bit ASN of the BGP peer for use with larger ASN values.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().String("tag-specifications", "", "The tags to apply to a resource when the local gateway virtual interface is being created.")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flags().String("vlan", "", "The virtual local area network (VLAN) used for the local gateway virtual interface.")
		ec2_createLocalGatewayVirtualInterfaceCmd.MarkFlagRequired("local-address")
		ec2_createLocalGatewayVirtualInterfaceCmd.MarkFlagRequired("local-gateway-virtual-interface-group-id")
		ec2_createLocalGatewayVirtualInterfaceCmd.Flag("no-dry-run").Hidden = true
		ec2_createLocalGatewayVirtualInterfaceCmd.MarkFlagRequired("outpost-lag-id")
		ec2_createLocalGatewayVirtualInterfaceCmd.MarkFlagRequired("peer-address")
		ec2_createLocalGatewayVirtualInterfaceCmd.MarkFlagRequired("vlan")
	})
	ec2Cmd.AddCommand(ec2_createLocalGatewayVirtualInterfaceCmd)
}
