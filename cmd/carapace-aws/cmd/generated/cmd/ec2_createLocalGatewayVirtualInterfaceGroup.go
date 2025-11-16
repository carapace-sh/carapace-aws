package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createLocalGatewayVirtualInterfaceGroupCmd = &cobra.Command{
	Use:   "create-local-gateway-virtual-interface-group",
	Short: "Create a local gateway virtual interface group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createLocalGatewayVirtualInterfaceGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createLocalGatewayVirtualInterfaceGroupCmd).Standalone()

		ec2_createLocalGatewayVirtualInterfaceGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createLocalGatewayVirtualInterfaceGroupCmd.Flags().String("local-bgp-asn", "", "The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP).")
		ec2_createLocalGatewayVirtualInterfaceGroupCmd.Flags().String("local-bgp-asn-extended", "", "The extended 32-bit ASN for the local BGP configuration.")
		ec2_createLocalGatewayVirtualInterfaceGroupCmd.Flags().String("local-gateway-id", "", "The ID of the local gateway.")
		ec2_createLocalGatewayVirtualInterfaceGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createLocalGatewayVirtualInterfaceGroupCmd.Flags().String("tag-specifications", "", "The tags to apply to the local gateway virtual interface group when the resource is being created.")
		ec2_createLocalGatewayVirtualInterfaceGroupCmd.MarkFlagRequired("local-gateway-id")
		ec2_createLocalGatewayVirtualInterfaceGroupCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createLocalGatewayVirtualInterfaceGroupCmd)
}
