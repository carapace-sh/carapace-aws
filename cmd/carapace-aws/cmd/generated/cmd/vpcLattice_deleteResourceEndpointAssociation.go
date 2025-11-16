package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteResourceEndpointAssociationCmd = &cobra.Command{
	Use:   "delete-resource-endpoint-association",
	Short: "Disassociates the resource configuration from the resource VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteResourceEndpointAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deleteResourceEndpointAssociationCmd).Standalone()

		vpcLattice_deleteResourceEndpointAssociationCmd.Flags().String("resource-endpoint-association-identifier", "", "The ID or ARN of the association.")
		vpcLattice_deleteResourceEndpointAssociationCmd.MarkFlagRequired("resource-endpoint-association-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deleteResourceEndpointAssociationCmd)
}
