package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateListenerCmd = &cobra.Command{
	Use:   "update-listener",
	Short: "Updates the specified listener for the specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateListenerCmd).Standalone()

	vpcLattice_updateListenerCmd.Flags().String("default-action", "", "The action for the default rule.")
	vpcLattice_updateListenerCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
	vpcLattice_updateListenerCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
	vpcLattice_updateListenerCmd.MarkFlagRequired("default-action")
	vpcLattice_updateListenerCmd.MarkFlagRequired("listener-identifier")
	vpcLattice_updateListenerCmd.MarkFlagRequired("service-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_updateListenerCmd)
}
