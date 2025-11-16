package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpcCmd = &cobra.Command{
	Use:   "create-vpc",
	Short: "Creates a VPC with the specified CIDR blocks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpcCmd).Standalone()

	ec2_createVpcCmd.Flags().Bool("amazon-provided-ipv6-cidr-block", false, "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC.")
	ec2_createVpcCmd.Flags().String("cidr-block", "", "The IPv4 network range for the VPC, in CIDR notation.")
	ec2_createVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVpcCmd.Flags().String("instance-tenancy", "", "The tenancy options for instances launched into the VPC.")
	ec2_createVpcCmd.Flags().String("ipv4-ipam-pool-id", "", "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR.")
	ec2_createVpcCmd.Flags().String("ipv4-netmask-length", "", "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool.")
	ec2_createVpcCmd.Flags().String("ipv6-cidr-block", "", "The IPv6 CIDR block from the IPv6 address pool.")
	ec2_createVpcCmd.Flags().String("ipv6-cidr-block-network-border-group", "", "The name of the location from which we advertise the IPV6 CIDR block.")
	ec2_createVpcCmd.Flags().String("ipv6-ipam-pool-id", "", "The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6 CIDR.")
	ec2_createVpcCmd.Flags().String("ipv6-netmask-length", "", "The netmask length of the IPv6 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool.")
	ec2_createVpcCmd.Flags().String("ipv6-pool", "", "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.")
	ec2_createVpcCmd.Flags().Bool("no-amazon-provided-ipv6-cidr-block", false, "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC.")
	ec2_createVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVpcCmd.Flags().String("tag-specifications", "", "The tags to assign to the VPC.")
	ec2_createVpcCmd.Flag("no-amazon-provided-ipv6-cidr-block").Hidden = true
	ec2_createVpcCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createVpcCmd)
}
