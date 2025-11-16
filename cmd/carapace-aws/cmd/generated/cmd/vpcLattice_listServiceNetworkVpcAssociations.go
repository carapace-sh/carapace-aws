package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listServiceNetworkVpcAssociationsCmd = &cobra.Command{
	Use:   "list-service-network-vpc-associations",
	Short: "Lists the associations between a service network and a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listServiceNetworkVpcAssociationsCmd).Standalone()

	vpcLattice_listServiceNetworkVpcAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	vpcLattice_listServiceNetworkVpcAssociationsCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	vpcLattice_listServiceNetworkVpcAssociationsCmd.Flags().String("service-network-identifier", "", "The ID or ARN of the service network.")
	vpcLattice_listServiceNetworkVpcAssociationsCmd.Flags().String("vpc-identifier", "", "The ID or ARN of the VPC.")
	vpcLatticeCmd.AddCommand(vpcLattice_listServiceNetworkVpcAssociationsCmd)
}
