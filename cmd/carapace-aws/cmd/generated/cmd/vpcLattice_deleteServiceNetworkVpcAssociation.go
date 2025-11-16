package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteServiceNetworkVpcAssociationCmd = &cobra.Command{
	Use:   "delete-service-network-vpc-association",
	Short: "Disassociates the VPC from the service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteServiceNetworkVpcAssociationCmd).Standalone()

	vpcLattice_deleteServiceNetworkVpcAssociationCmd.Flags().String("service-network-vpc-association-identifier", "", "The ID or ARN of the association.")
	vpcLattice_deleteServiceNetworkVpcAssociationCmd.MarkFlagRequired("service-network-vpc-association-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_deleteServiceNetworkVpcAssociationCmd)
}
