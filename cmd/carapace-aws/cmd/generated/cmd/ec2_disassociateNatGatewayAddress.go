package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateNatGatewayAddressCmd = &cobra.Command{
	Use:   "disassociate-nat-gateway-address",
	Short: "Disassociates secondary Elastic IP addresses (EIPs) from a public NAT gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateNatGatewayAddressCmd).Standalone()

	ec2_disassociateNatGatewayAddressCmd.Flags().String("association-ids", "", "The association IDs of EIPs that have been associated with the NAT gateway.")
	ec2_disassociateNatGatewayAddressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateNatGatewayAddressCmd.Flags().String("max-drain-duration-seconds", "", "The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress.")
	ec2_disassociateNatGatewayAddressCmd.Flags().String("nat-gateway-id", "", "The ID of the NAT gateway.")
	ec2_disassociateNatGatewayAddressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disassociateNatGatewayAddressCmd.MarkFlagRequired("association-ids")
	ec2_disassociateNatGatewayAddressCmd.MarkFlagRequired("nat-gateway-id")
	ec2_disassociateNatGatewayAddressCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disassociateNatGatewayAddressCmd)
}
