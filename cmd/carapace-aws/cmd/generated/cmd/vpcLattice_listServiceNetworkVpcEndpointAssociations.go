package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listServiceNetworkVpcEndpointAssociationsCmd = &cobra.Command{
	Use:   "list-service-network-vpc-endpoint-associations",
	Short: "Lists the associations between a service network and a VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listServiceNetworkVpcEndpointAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_listServiceNetworkVpcEndpointAssociationsCmd).Standalone()

		vpcLattice_listServiceNetworkVpcEndpointAssociationsCmd.Flags().String("max-results", "", "The maximum page size.")
		vpcLattice_listServiceNetworkVpcEndpointAssociationsCmd.Flags().String("next-token", "", "If there are additional results, a pagination token for the next page of results.")
		vpcLattice_listServiceNetworkVpcEndpointAssociationsCmd.Flags().String("service-network-identifier", "", "The ID of the service network associated with the VPC endpoint.")
		vpcLattice_listServiceNetworkVpcEndpointAssociationsCmd.MarkFlagRequired("service-network-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_listServiceNetworkVpcEndpointAssociationsCmd)
}
