package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteServiceNetworkCmd = &cobra.Command{
	Use:   "delete-service-network",
	Short: "Deletes a service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteServiceNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deleteServiceNetworkCmd).Standalone()

		vpcLattice_deleteServiceNetworkCmd.Flags().String("service-network-identifier", "", "The ID or ARN of the service network.")
		vpcLattice_deleteServiceNetworkCmd.MarkFlagRequired("service-network-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deleteServiceNetworkCmd)
}
