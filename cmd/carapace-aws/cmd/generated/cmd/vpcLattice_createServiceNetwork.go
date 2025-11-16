package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createServiceNetworkCmd = &cobra.Command{
	Use:   "create-service-network",
	Short: "Creates a service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createServiceNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_createServiceNetworkCmd).Standalone()

		vpcLattice_createServiceNetworkCmd.Flags().String("auth-type", "", "The type of IAM policy.")
		vpcLattice_createServiceNetworkCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		vpcLattice_createServiceNetworkCmd.Flags().String("name", "", "The name of the service network.")
		vpcLattice_createServiceNetworkCmd.Flags().String("sharing-config", "", "Specify if the service network should be enabled for sharing.")
		vpcLattice_createServiceNetworkCmd.Flags().String("tags", "", "The tags for the service network.")
		vpcLattice_createServiceNetworkCmd.MarkFlagRequired("name")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_createServiceNetworkCmd)
}
