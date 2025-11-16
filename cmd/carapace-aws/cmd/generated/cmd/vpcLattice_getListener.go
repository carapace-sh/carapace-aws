package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getListenerCmd = &cobra.Command{
	Use:   "get-listener",
	Short: "Retrieves information about the specified listener for the specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getListenerCmd).Standalone()

	vpcLattice_getListenerCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
	vpcLattice_getListenerCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
	vpcLattice_getListenerCmd.MarkFlagRequired("listener-identifier")
	vpcLattice_getListenerCmd.MarkFlagRequired("service-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getListenerCmd)
}
