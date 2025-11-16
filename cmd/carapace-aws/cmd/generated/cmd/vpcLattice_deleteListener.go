package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteListenerCmd = &cobra.Command{
	Use:   "delete-listener",
	Short: "Deletes the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deleteListenerCmd).Standalone()

		vpcLattice_deleteListenerCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
		vpcLattice_deleteListenerCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_deleteListenerCmd.MarkFlagRequired("listener-identifier")
		vpcLattice_deleteListenerCmd.MarkFlagRequired("service-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deleteListenerCmd)
}
