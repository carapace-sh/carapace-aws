package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getServiceCmd = &cobra.Command{
	Use:   "get-service",
	Short: "Retrieves information about the specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_getServiceCmd).Standalone()

		vpcLattice_getServiceCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_getServiceCmd.MarkFlagRequired("service-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_getServiceCmd)
}
