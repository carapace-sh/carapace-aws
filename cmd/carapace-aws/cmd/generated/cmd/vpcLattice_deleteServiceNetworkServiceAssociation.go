package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteServiceNetworkServiceAssociationCmd = &cobra.Command{
	Use:   "delete-service-network-service-association",
	Short: "Deletes the association between a service and a service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteServiceNetworkServiceAssociationCmd).Standalone()

	vpcLattice_deleteServiceNetworkServiceAssociationCmd.Flags().String("service-network-service-association-identifier", "", "The ID or ARN of the association.")
	vpcLattice_deleteServiceNetworkServiceAssociationCmd.MarkFlagRequired("service-network-service-association-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_deleteServiceNetworkServiceAssociationCmd)
}
