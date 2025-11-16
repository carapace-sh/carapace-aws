package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateSubnetCidrBlockCmd = &cobra.Command{
	Use:   "associate-subnet-cidr-block",
	Short: "Associates a CIDR block with your subnet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateSubnetCidrBlockCmd).Standalone()

	ec2_associateSubnetCidrBlockCmd.Flags().String("ipv6-cidr-block", "", "The IPv6 CIDR block for your subnet.")
	ec2_associateSubnetCidrBlockCmd.Flags().String("ipv6-ipam-pool-id", "", "An IPv6 IPAM pool ID.")
	ec2_associateSubnetCidrBlockCmd.Flags().String("ipv6-netmask-length", "", "An IPv6 netmask length.")
	ec2_associateSubnetCidrBlockCmd.Flags().String("subnet-id", "", "The ID of your subnet.")
	ec2_associateSubnetCidrBlockCmd.MarkFlagRequired("subnet-id")
	ec2Cmd.AddCommand(ec2_associateSubnetCidrBlockCmd)
}
