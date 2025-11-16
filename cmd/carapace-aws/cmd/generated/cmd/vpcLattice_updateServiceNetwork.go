package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateServiceNetworkCmd = &cobra.Command{
	Use:   "update-service-network",
	Short: "Updates the specified service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateServiceNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_updateServiceNetworkCmd).Standalone()

		vpcLattice_updateServiceNetworkCmd.Flags().String("auth-type", "", "The type of IAM policy.")
		vpcLattice_updateServiceNetworkCmd.Flags().String("service-network-identifier", "", "The ID or ARN of the service network.")
		vpcLattice_updateServiceNetworkCmd.MarkFlagRequired("auth-type")
		vpcLattice_updateServiceNetworkCmd.MarkFlagRequired("service-network-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_updateServiceNetworkCmd)
}
