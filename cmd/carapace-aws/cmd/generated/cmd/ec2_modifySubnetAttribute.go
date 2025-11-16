package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifySubnetAttributeCmd = &cobra.Command{
	Use:   "modify-subnet-attribute",
	Short: "Modifies a subnet attribute.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifySubnetAttributeCmd).Standalone()

	ec2_modifySubnetAttributeCmd.Flags().String("assign-ipv6-address-on-creation", "", "Specify `true` to indicate that network interfaces created in the specified subnet should be assigned an IPv6 address.")
	ec2_modifySubnetAttributeCmd.Flags().String("customer-owned-ipv4-pool", "", "The customer-owned IPv4 address pool associated with the subnet.")
	ec2_modifySubnetAttributeCmd.Flags().String("disable-lni-at-device-index", "", "Specify `true` to indicate that local network interfaces at the current position should be disabled.")
	ec2_modifySubnetAttributeCmd.Flags().String("enable-dns64", "", "Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.")
	ec2_modifySubnetAttributeCmd.Flags().String("enable-lni-at-device-index", "", "Indicates the device position for local network interfaces in this subnet.")
	ec2_modifySubnetAttributeCmd.Flags().String("enable-resource-name-dns-aaaarecord-on-launch", "", "Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.")
	ec2_modifySubnetAttributeCmd.Flags().String("enable-resource-name-dns-arecord-on-launch", "", "Indicates whether to respond to DNS queries for instance hostnames with DNS A records.")
	ec2_modifySubnetAttributeCmd.Flags().String("map-customer-owned-ip-on-launch", "", "Specify `true` to indicate that network interfaces attached to instances created in the specified subnet should be assigned a customer-owned IPv4 address.")
	ec2_modifySubnetAttributeCmd.Flags().String("map-public-ip-on-launch", "", "Specify `true` to indicate that network interfaces attached to instances created in the specified subnet should be assigned a public IPv4 address.")
	ec2_modifySubnetAttributeCmd.Flags().String("private-dns-hostname-type-on-launch", "", "The type of hostname to assign to instances in the subnet at launch.")
	ec2_modifySubnetAttributeCmd.Flags().String("subnet-id", "", "The ID of the subnet.")
	ec2_modifySubnetAttributeCmd.MarkFlagRequired("subnet-id")
	ec2Cmd.AddCommand(ec2_modifySubnetAttributeCmd)
}
