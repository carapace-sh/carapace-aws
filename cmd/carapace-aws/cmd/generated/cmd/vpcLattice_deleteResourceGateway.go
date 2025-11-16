package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteResourceGatewayCmd = &cobra.Command{
	Use:   "delete-resource-gateway",
	Short: "Deletes the specified resource gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteResourceGatewayCmd).Standalone()

	vpcLattice_deleteResourceGatewayCmd.Flags().String("resource-gateway-identifier", "", "The ID or ARN of the resource gateway.")
	vpcLattice_deleteResourceGatewayCmd.MarkFlagRequired("resource-gateway-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_deleteResourceGatewayCmd)
}
