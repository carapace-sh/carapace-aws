package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createServiceNetworkServiceAssociationCmd = &cobra.Command{
	Use:   "create-service-network-service-association",
	Short: "Associates the specified service with the specified service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createServiceNetworkServiceAssociationCmd).Standalone()

	vpcLattice_createServiceNetworkServiceAssociationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	vpcLattice_createServiceNetworkServiceAssociationCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
	vpcLattice_createServiceNetworkServiceAssociationCmd.Flags().String("service-network-identifier", "", "The ID or ARN of the service network.")
	vpcLattice_createServiceNetworkServiceAssociationCmd.Flags().String("tags", "", "The tags for the association.")
	vpcLattice_createServiceNetworkServiceAssociationCmd.MarkFlagRequired("service-identifier")
	vpcLattice_createServiceNetworkServiceAssociationCmd.MarkFlagRequired("service-network-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_createServiceNetworkServiceAssociationCmd)
}
