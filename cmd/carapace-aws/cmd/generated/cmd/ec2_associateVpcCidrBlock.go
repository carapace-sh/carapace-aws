package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateVpcCidrBlockCmd = &cobra.Command{
	Use:   "associate-vpc-cidr-block",
	Short: "Associates a CIDR block with your VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateVpcCidrBlockCmd).Standalone()

	ec2_associateVpcCidrBlockCmd.Flags().Bool("amazon-provided-ipv6-cidr-block", false, "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC.")
	ec2_associateVpcCidrBlockCmd.Flags().String("cidr-block", "", "An IPv4 CIDR block to associate with the VPC.")
	ec2_associateVpcCidrBlockCmd.Flags().String("ipv4-ipam-pool-id", "", "Associate a CIDR allocated from an IPv4 IPAM pool to a VPC.")
	ec2_associateVpcCidrBlockCmd.Flags().String("ipv4-netmask-length", "", "The netmask length of the IPv4 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool.")
	ec2_associateVpcCidrBlockCmd.Flags().String("ipv6-cidr-block", "", "An IPv6 CIDR block from the IPv6 address pool.")
	ec2_associateVpcCidrBlockCmd.Flags().String("ipv6-cidr-block-network-border-group", "", "The name of the location from which we advertise the IPV6 CIDR block.")
	ec2_associateVpcCidrBlockCmd.Flags().String("ipv6-ipam-pool-id", "", "Associates a CIDR allocated from an IPv6 IPAM pool to a VPC.")
	ec2_associateVpcCidrBlockCmd.Flags().String("ipv6-netmask-length", "", "The netmask length of the IPv6 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool.")
	ec2_associateVpcCidrBlockCmd.Flags().String("ipv6-pool", "", "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.")
	ec2_associateVpcCidrBlockCmd.Flags().Bool("no-amazon-provided-ipv6-cidr-block", false, "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC.")
	ec2_associateVpcCidrBlockCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_associateVpcCidrBlockCmd.Flag("no-amazon-provided-ipv6-cidr-block").Hidden = true
	ec2_associateVpcCidrBlockCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_associateVpcCidrBlockCmd)
}
