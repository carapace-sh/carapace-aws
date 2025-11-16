package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateAddressCmd = &cobra.Command{
	Use:   "associate-address",
	Short: "Associates an Elastic IP address, or carrier IP address (for instances that are in subnets in Wavelength Zones) with an instance or a network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateAddressCmd).Standalone()

	ec2_associateAddressCmd.Flags().String("allocation-id", "", "The allocation ID.")
	ec2_associateAddressCmd.Flags().Bool("allow-reassociation", false, "Reassociation is automatic, but you can specify false to ensure the operation fails if the Elastic IP address is already associated with another resource.")
	ec2_associateAddressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateAddressCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_associateAddressCmd.Flags().String("network-interface-id", "", "The ID of the network interface.")
	ec2_associateAddressCmd.Flags().Bool("no-allow-reassociation", false, "Reassociation is automatic, but you can specify false to ensure the operation fails if the Elastic IP address is already associated with another resource.")
	ec2_associateAddressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateAddressCmd.Flags().String("private-ip-address", "", "The primary or secondary private IP address to associate with the Elastic IP address.")
	ec2_associateAddressCmd.Flags().String("public-ip", "", "Deprecated.")
	ec2_associateAddressCmd.Flag("no-allow-reassociation").Hidden = true
	ec2_associateAddressCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_associateAddressCmd)
}
