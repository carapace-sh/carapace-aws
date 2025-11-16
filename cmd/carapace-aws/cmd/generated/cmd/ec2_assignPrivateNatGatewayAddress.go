package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_assignPrivateNatGatewayAddressCmd = &cobra.Command{
	Use:   "assign-private-nat-gateway-address",
	Short: "Assigns private IPv4 addresses to a private NAT gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_assignPrivateNatGatewayAddressCmd).Standalone()

	ec2_assignPrivateNatGatewayAddressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_assignPrivateNatGatewayAddressCmd.Flags().String("nat-gateway-id", "", "The ID of the NAT gateway.")
	ec2_assignPrivateNatGatewayAddressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_assignPrivateNatGatewayAddressCmd.Flags().String("private-ip-address-count", "", "The number of private IP addresses to assign to the NAT gateway.")
	ec2_assignPrivateNatGatewayAddressCmd.Flags().String("private-ip-addresses", "", "The private IPv4 addresses you want to assign to the private NAT gateway.")
	ec2_assignPrivateNatGatewayAddressCmd.MarkFlagRequired("nat-gateway-id")
	ec2_assignPrivateNatGatewayAddressCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_assignPrivateNatGatewayAddressCmd)
}
