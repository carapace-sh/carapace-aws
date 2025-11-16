package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_unassignPrivateNatGatewayAddressCmd = &cobra.Command{
	Use:   "unassign-private-nat-gateway-address",
	Short: "Unassigns secondary private IPv4 addresses from a private NAT gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_unassignPrivateNatGatewayAddressCmd).Standalone()

	ec2_unassignPrivateNatGatewayAddressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_unassignPrivateNatGatewayAddressCmd.Flags().String("max-drain-duration-seconds", "", "The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress.")
	ec2_unassignPrivateNatGatewayAddressCmd.Flags().String("nat-gateway-id", "", "The ID of the NAT gateway.")
	ec2_unassignPrivateNatGatewayAddressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_unassignPrivateNatGatewayAddressCmd.Flags().String("private-ip-addresses", "", "The private IPv4 addresses you want to unassign.")
	ec2_unassignPrivateNatGatewayAddressCmd.MarkFlagRequired("nat-gateway-id")
	ec2_unassignPrivateNatGatewayAddressCmd.Flag("no-dry-run").Hidden = true
	ec2_unassignPrivateNatGatewayAddressCmd.MarkFlagRequired("private-ip-addresses")
	ec2Cmd.AddCommand(ec2_unassignPrivateNatGatewayAddressCmd)
}
