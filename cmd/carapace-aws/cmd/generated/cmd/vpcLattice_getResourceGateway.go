package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getResourceGatewayCmd = &cobra.Command{
	Use:   "get-resource-gateway",
	Short: "Retrieves information about the specified resource gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getResourceGatewayCmd).Standalone()

	vpcLattice_getResourceGatewayCmd.Flags().String("resource-gateway-identifier", "", "The ID of the resource gateway.")
	vpcLattice_getResourceGatewayCmd.MarkFlagRequired("resource-gateway-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getResourceGatewayCmd)
}
