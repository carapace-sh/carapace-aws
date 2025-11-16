package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deregisterTargetsCmd = &cobra.Command{
	Use:   "deregister-targets",
	Short: "Deregisters the specified targets from the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deregisterTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deregisterTargetsCmd).Standalone()

		vpcLattice_deregisterTargetsCmd.Flags().String("target-group-identifier", "", "The ID or ARN of the target group.")
		vpcLattice_deregisterTargetsCmd.Flags().String("targets", "", "The targets to deregister.")
		vpcLattice_deregisterTargetsCmd.MarkFlagRequired("target-group-identifier")
		vpcLattice_deregisterTargetsCmd.MarkFlagRequired("targets")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deregisterTargetsCmd)
}
