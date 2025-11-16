package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getServiceNetworkResourceAssociationCmd = &cobra.Command{
	Use:   "get-service-network-resource-association",
	Short: "Retrieves information about the specified association between a service network and a resource configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getServiceNetworkResourceAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_getServiceNetworkResourceAssociationCmd).Standalone()

		vpcLattice_getServiceNetworkResourceAssociationCmd.Flags().String("service-network-resource-association-identifier", "", "The ID of the association.")
		vpcLattice_getServiceNetworkResourceAssociationCmd.MarkFlagRequired("service-network-resource-association-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_getServiceNetworkResourceAssociationCmd)
}
