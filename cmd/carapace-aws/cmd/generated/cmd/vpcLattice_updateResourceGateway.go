package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateResourceGatewayCmd = &cobra.Command{
	Use:   "update-resource-gateway",
	Short: "Updates the specified resource gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateResourceGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_updateResourceGatewayCmd).Standalone()

		vpcLattice_updateResourceGatewayCmd.Flags().String("resource-gateway-identifier", "", "The ID or ARN of the resource gateway.")
		vpcLattice_updateResourceGatewayCmd.Flags().String("security-group-ids", "", "The IDs of the security groups associated with the resource gateway.")
		vpcLattice_updateResourceGatewayCmd.MarkFlagRequired("resource-gateway-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_updateResourceGatewayCmd)
}
