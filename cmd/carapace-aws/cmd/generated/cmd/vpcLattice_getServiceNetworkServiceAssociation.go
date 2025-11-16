package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getServiceNetworkServiceAssociationCmd = &cobra.Command{
	Use:   "get-service-network-service-association",
	Short: "Retrieves information about the specified association between a service network and a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getServiceNetworkServiceAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_getServiceNetworkServiceAssociationCmd).Standalone()

		vpcLattice_getServiceNetworkServiceAssociationCmd.Flags().String("service-network-service-association-identifier", "", "The ID or ARN of the association.")
		vpcLattice_getServiceNetworkServiceAssociationCmd.MarkFlagRequired("service-network-service-association-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_getServiceNetworkServiceAssociationCmd)
}
