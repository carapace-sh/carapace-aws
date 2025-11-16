package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createRouteCmd = &cobra.Command{
	Use:   "create-route",
	Short: "Creates a route in a route table within a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createRouteCmd).Standalone()

		ec2_createRouteCmd.Flags().String("carrier-gateway-id", "", "The ID of the carrier gateway.")
		ec2_createRouteCmd.Flags().String("core-network-arn", "", "The Amazon Resource Name (ARN) of the core network.")
		ec2_createRouteCmd.Flags().String("destination-cidr-block", "", "The IPv4 CIDR address block used for the destination match.")
		ec2_createRouteCmd.Flags().String("destination-ipv6-cidr-block", "", "The IPv6 CIDR block used for the destination match.")
		ec2_createRouteCmd.Flags().String("destination-prefix-list-id", "", "The ID of a prefix list used for the destination match.")
		ec2_createRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createRouteCmd.Flags().String("egress-only-internet-gateway-id", "", "\\[IPv6 traffic only] The ID of an egress-only internet gateway.")
		ec2_createRouteCmd.Flags().String("gateway-id", "", "The ID of an internet gateway or virtual private gateway attached to your VPC.")
		ec2_createRouteCmd.Flags().String("instance-id", "", "The ID of a NAT instance in your VPC.")
		ec2_createRouteCmd.Flags().String("local-gateway-id", "", "The ID of the local gateway.")
		ec2_createRouteCmd.Flags().String("nat-gateway-id", "", "\\[IPv4 traffic only] The ID of a NAT gateway.")
		ec2_createRouteCmd.Flags().String("network-interface-id", "", "The ID of a network interface.")
		ec2_createRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createRouteCmd.Flags().String("odb-network-arn", "", "The Amazon Resource Name (ARN) of the ODB network.")
		ec2_createRouteCmd.Flags().String("route-table-id", "", "The ID of the route table for the route.")
		ec2_createRouteCmd.Flags().String("transit-gateway-id", "", "The ID of a transit gateway.")
		ec2_createRouteCmd.Flags().String("vpc-endpoint-id", "", "The ID of a VPC endpoint.")
		ec2_createRouteCmd.Flags().String("vpc-peering-connection-id", "", "The ID of a VPC peering connection.")
		ec2_createRouteCmd.Flag("no-dry-run").Hidden = true
		ec2_createRouteCmd.MarkFlagRequired("route-table-id")
	})
	ec2Cmd.AddCommand(ec2_createRouteCmd)
}
