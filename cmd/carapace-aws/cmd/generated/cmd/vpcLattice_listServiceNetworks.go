package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listServiceNetworksCmd = &cobra.Command{
	Use:   "list-service-networks",
	Short: "Lists the service networks owned by or shared with this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listServiceNetworksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_listServiceNetworksCmd).Standalone()

		vpcLattice_listServiceNetworksCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		vpcLattice_listServiceNetworksCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_listServiceNetworksCmd)
}
