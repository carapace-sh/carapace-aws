package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createSubnetCmd = &cobra.Command{
	Use:   "create-subnet",
	Short: "Creates a subnet in the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createSubnetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createSubnetCmd).Standalone()

		ec2_createSubnetCmd.Flags().String("availability-zone", "", "The Availability Zone or Local Zone for the subnet.")
		ec2_createSubnetCmd.Flags().String("availability-zone-id", "", "The AZ ID or the Local Zone ID of the subnet.")
		ec2_createSubnetCmd.Flags().String("cidr-block", "", "The IPv4 network range for the subnet, in CIDR notation.")
		ec2_createSubnetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createSubnetCmd.Flags().String("ipv4-ipam-pool-id", "", "An IPv4 IPAM pool ID for the subnet.")
		ec2_createSubnetCmd.Flags().String("ipv4-netmask-length", "", "An IPv4 netmask length for the subnet.")
		ec2_createSubnetCmd.Flags().String("ipv6-cidr-block", "", "The IPv6 network range for the subnet, in CIDR notation.")
		ec2_createSubnetCmd.Flags().String("ipv6-ipam-pool-id", "", "An IPv6 IPAM pool ID for the subnet.")
		ec2_createSubnetCmd.Flags().Bool("ipv6-native", false, "Indicates whether to create an IPv6 only subnet.")
		ec2_createSubnetCmd.Flags().String("ipv6-netmask-length", "", "An IPv6 netmask length for the subnet.")
		ec2_createSubnetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createSubnetCmd.Flags().Bool("no-ipv6-native", false, "Indicates whether to create an IPv6 only subnet.")
		ec2_createSubnetCmd.Flags().String("outpost-arn", "", "The Amazon Resource Name (ARN) of the Outpost.")
		ec2_createSubnetCmd.Flags().String("tag-specifications", "", "The tags to assign to the subnet.")
		ec2_createSubnetCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_createSubnetCmd.Flag("no-dry-run").Hidden = true
		ec2_createSubnetCmd.Flag("no-ipv6-native").Hidden = true
		ec2_createSubnetCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_createSubnetCmd)
}
