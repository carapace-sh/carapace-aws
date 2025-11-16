package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateNatGatewayAddressCmd = &cobra.Command{
	Use:   "associate-nat-gateway-address",
	Short: "Associates Elastic IP addresses (EIPs) and private IPv4 addresses with a public NAT gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateNatGatewayAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_associateNatGatewayAddressCmd).Standalone()

		ec2_associateNatGatewayAddressCmd.Flags().String("allocation-ids", "", "The allocation IDs of EIPs that you want to associate with your NAT gateway.")
		ec2_associateNatGatewayAddressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateNatGatewayAddressCmd.Flags().String("nat-gateway-id", "", "The ID of the NAT gateway.")
		ec2_associateNatGatewayAddressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateNatGatewayAddressCmd.Flags().String("private-ip-addresses", "", "The private IPv4 addresses that you want to assign to the NAT gateway.")
		ec2_associateNatGatewayAddressCmd.MarkFlagRequired("allocation-ids")
		ec2_associateNatGatewayAddressCmd.MarkFlagRequired("nat-gateway-id")
		ec2_associateNatGatewayAddressCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_associateNatGatewayAddressCmd)
}
