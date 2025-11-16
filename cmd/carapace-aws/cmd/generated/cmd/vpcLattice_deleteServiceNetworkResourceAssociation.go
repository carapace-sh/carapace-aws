package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteServiceNetworkResourceAssociationCmd = &cobra.Command{
	Use:   "delete-service-network-resource-association",
	Short: "Deletes the association between a service network and a resource configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteServiceNetworkResourceAssociationCmd).Standalone()

	vpcLattice_deleteServiceNetworkResourceAssociationCmd.Flags().String("service-network-resource-association-identifier", "", "The ID of the association.")
	vpcLattice_deleteServiceNetworkResourceAssociationCmd.MarkFlagRequired("service-network-resource-association-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_deleteServiceNetworkResourceAssociationCmd)
}
