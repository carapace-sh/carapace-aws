package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateServiceNetworkVpcAssociationCmd = &cobra.Command{
	Use:   "update-service-network-vpc-association",
	Short: "Updates the service network and VPC association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateServiceNetworkVpcAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_updateServiceNetworkVpcAssociationCmd).Standalone()

		vpcLattice_updateServiceNetworkVpcAssociationCmd.Flags().String("security-group-ids", "", "The IDs of the security groups.")
		vpcLattice_updateServiceNetworkVpcAssociationCmd.Flags().String("service-network-vpc-association-identifier", "", "The ID or ARN of the association.")
		vpcLattice_updateServiceNetworkVpcAssociationCmd.MarkFlagRequired("security-group-ids")
		vpcLattice_updateServiceNetworkVpcAssociationCmd.MarkFlagRequired("service-network-vpc-association-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_updateServiceNetworkVpcAssociationCmd)
}
