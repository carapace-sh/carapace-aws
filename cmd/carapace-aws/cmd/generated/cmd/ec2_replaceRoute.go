package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_replaceRouteCmd = &cobra.Command{
	Use:   "replace-route",
	Short: "Replaces an existing route within a route table in a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_replaceRouteCmd).Standalone()

	ec2_replaceRouteCmd.Flags().String("carrier-gateway-id", "", "\\[IPv4 traffic only] The ID of a carrier gateway.")
	ec2_replaceRouteCmd.Flags().String("core-network-arn", "", "The Amazon Resource Name (ARN) of the core network.")
	ec2_replaceRouteCmd.Flags().String("destination-cidr-block", "", "The IPv4 CIDR address block used for the destination match.")
	ec2_replaceRouteCmd.Flags().String("destination-ipv6-cidr-block", "", "The IPv6 CIDR address block used for the destination match.")
	ec2_replaceRouteCmd.Flags().String("destination-prefix-list-id", "", "The ID of the prefix list for the route.")
	ec2_replaceRouteCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_replaceRouteCmd.Flags().String("egress-only-internet-gateway-id", "", "\\[IPv6 traffic only] The ID of an egress-only internet gateway.")
	ec2_replaceRouteCmd.Flags().String("gateway-id", "", "The ID of an internet gateway or virtual private gateway.")
	ec2_replaceRouteCmd.Flags().String("instance-id", "", "The ID of a NAT instance in your VPC.")
	ec2_replaceRouteCmd.Flags().String("local-gateway-id", "", "The ID of the local gateway.")
	ec2_replaceRouteCmd.Flags().Bool("local-target", false, "Specifies whether to reset the local route to its default target (`local`).")
	ec2_replaceRouteCmd.Flags().String("nat-gateway-id", "", "\\[IPv4 traffic only] The ID of a NAT gateway.")
	ec2_replaceRouteCmd.Flags().String("network-interface-id", "", "The ID of a network interface.")
	ec2_replaceRouteCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_replaceRouteCmd.Flags().Bool("no-local-target", false, "Specifies whether to reset the local route to its default target (`local`).")
	ec2_replaceRouteCmd.Flags().String("odb-network-arn", "", "The Amazon Resource Name (ARN) of the ODB network.")
	ec2_replaceRouteCmd.Flags().String("route-table-id", "", "The ID of the route table.")
	ec2_replaceRouteCmd.Flags().String("transit-gateway-id", "", "The ID of a transit gateway.")
	ec2_replaceRouteCmd.Flags().String("vpc-endpoint-id", "", "The ID of a VPC endpoint.")
	ec2_replaceRouteCmd.Flags().String("vpc-peering-connection-id", "", "The ID of a VPC peering connection.")
	ec2_replaceRouteCmd.Flag("no-dry-run").Hidden = true
	ec2_replaceRouteCmd.Flag("no-local-target").Hidden = true
	ec2_replaceRouteCmd.MarkFlagRequired("route-table-id")
	ec2Cmd.AddCommand(ec2_replaceRouteCmd)
}
