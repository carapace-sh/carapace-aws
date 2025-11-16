package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getServiceNetworkCmd = &cobra.Command{
	Use:   "get-service-network",
	Short: "Retrieves information about the specified service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getServiceNetworkCmd).Standalone()

	vpcLattice_getServiceNetworkCmd.Flags().String("service-network-identifier", "", "The ID or ARN of the service network.")
	vpcLattice_getServiceNetworkCmd.MarkFlagRequired("service-network-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getServiceNetworkCmd)
}
