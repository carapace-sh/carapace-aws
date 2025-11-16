package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createNatGatewayCmd = &cobra.Command{
	Use:   "create-nat-gateway",
	Short: "Creates a NAT gateway in the specified subnet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createNatGatewayCmd).Standalone()

	ec2_createNatGatewayCmd.Flags().String("allocation-id", "", "\\[Public NAT gateways only] The allocation ID of an Elastic IP address to associate with the NAT gateway.")
	ec2_createNatGatewayCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createNatGatewayCmd.Flags().String("connectivity-type", "", "Indicates whether the NAT gateway supports public or private connectivity.")
	ec2_createNatGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createNatGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createNatGatewayCmd.Flags().String("private-ip-address", "", "The private IPv4 address to assign to the NAT gateway.")
	ec2_createNatGatewayCmd.Flags().String("secondary-allocation-ids", "", "Secondary EIP allocation IDs.")
	ec2_createNatGatewayCmd.Flags().String("secondary-private-ip-address-count", "", "\\[Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway.")
	ec2_createNatGatewayCmd.Flags().String("secondary-private-ip-addresses", "", "Secondary private IPv4 addresses.")
	ec2_createNatGatewayCmd.Flags().String("subnet-id", "", "The ID of the subnet in which to create the NAT gateway.")
	ec2_createNatGatewayCmd.Flags().String("tag-specifications", "", "The tags to assign to the NAT gateway.")
	ec2_createNatGatewayCmd.Flag("no-dry-run").Hidden = true
	ec2_createNatGatewayCmd.MarkFlagRequired("subnet-id")
	ec2Cmd.AddCommand(ec2_createNatGatewayCmd)
}
