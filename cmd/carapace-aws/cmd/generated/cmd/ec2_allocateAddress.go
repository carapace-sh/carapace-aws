package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_allocateAddressCmd = &cobra.Command{
	Use:   "allocate-address",
	Short: "Allocates an Elastic IP address to your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_allocateAddressCmd).Standalone()

	ec2_allocateAddressCmd.Flags().String("address", "", "The Elastic IP address to recover or an IPv4 address from an address pool.")
	ec2_allocateAddressCmd.Flags().String("customer-owned-ipv4-pool", "", "The ID of a customer-owned address pool.")
	ec2_allocateAddressCmd.Flags().String("domain", "", "The network (`vpc`).")
	ec2_allocateAddressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_allocateAddressCmd.Flags().String("ipam-pool-id", "", "The ID of an IPAM pool which has an Amazon-provided or BYOIP public IPv4 CIDR provisioned to it.")
	ec2_allocateAddressCmd.Flags().String("network-border-group", "", "A unique set of Availability Zones, Local Zones, or Wavelength Zones from which Amazon Web Services advertises IP addresses.")
	ec2_allocateAddressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_allocateAddressCmd.Flags().String("public-ipv4-pool", "", "The ID of an address pool that you own.")
	ec2_allocateAddressCmd.Flags().String("tag-specifications", "", "The tags to assign to the Elastic IP address.")
	ec2_allocateAddressCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_allocateAddressCmd)
}
