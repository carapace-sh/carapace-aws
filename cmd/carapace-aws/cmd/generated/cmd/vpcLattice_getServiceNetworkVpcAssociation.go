package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getServiceNetworkVpcAssociationCmd = &cobra.Command{
	Use:   "get-service-network-vpc-association",
	Short: "Retrieves information about the specified association between a service network and a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getServiceNetworkVpcAssociationCmd).Standalone()

	vpcLattice_getServiceNetworkVpcAssociationCmd.Flags().String("service-network-vpc-association-identifier", "", "The ID or ARN of the association.")
	vpcLattice_getServiceNetworkVpcAssociationCmd.MarkFlagRequired("service-network-vpc-association-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getServiceNetworkVpcAssociationCmd)
}
