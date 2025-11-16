package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listServiceNetworkServiceAssociationsCmd = &cobra.Command{
	Use:   "list-service-network-service-associations",
	Short: "Lists the associations between a service network and a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listServiceNetworkServiceAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_listServiceNetworkServiceAssociationsCmd).Standalone()

		vpcLattice_listServiceNetworkServiceAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		vpcLattice_listServiceNetworkServiceAssociationsCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
		vpcLattice_listServiceNetworkServiceAssociationsCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_listServiceNetworkServiceAssociationsCmd.Flags().String("service-network-identifier", "", "The ID or ARN of the service network.")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_listServiceNetworkServiceAssociationsCmd)
}
