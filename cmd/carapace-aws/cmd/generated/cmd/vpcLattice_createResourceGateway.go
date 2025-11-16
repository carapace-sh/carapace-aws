package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createResourceGatewayCmd = &cobra.Command{
	Use:   "create-resource-gateway",
	Short: "A resource gateway is a point of ingress into the VPC where a resource resides.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createResourceGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_createResourceGatewayCmd).Standalone()

		vpcLattice_createResourceGatewayCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		vpcLattice_createResourceGatewayCmd.Flags().String("ip-address-type", "", "A resource gateway can have IPv4, IPv6 or dualstack addresses.")
		vpcLattice_createResourceGatewayCmd.Flags().String("ipv4-addresses-per-eni", "", "The number of IPv4 addresses in each ENI for the resource gateway.")
		vpcLattice_createResourceGatewayCmd.Flags().String("name", "", "The name of the resource gateway.")
		vpcLattice_createResourceGatewayCmd.Flags().String("security-group-ids", "", "The IDs of the security groups to apply to the resource gateway.")
		vpcLattice_createResourceGatewayCmd.Flags().String("subnet-ids", "", "The IDs of the VPC subnets in which to create the resource gateway.")
		vpcLattice_createResourceGatewayCmd.Flags().String("tags", "", "The tags for the resource gateway.")
		vpcLattice_createResourceGatewayCmd.Flags().String("vpc-identifier", "", "The ID of the VPC for the resource gateway.")
		vpcLattice_createResourceGatewayCmd.MarkFlagRequired("name")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_createResourceGatewayCmd)
}
