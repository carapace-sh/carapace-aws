package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteServiceCmd = &cobra.Command{
	Use:   "delete-service",
	Short: "Deletes a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deleteServiceCmd).Standalone()

		vpcLattice_deleteServiceCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_deleteServiceCmd.MarkFlagRequired("service-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deleteServiceCmd)
}
