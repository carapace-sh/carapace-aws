package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteTargetGroupCmd = &cobra.Command{
	Use:   "delete-target-group",
	Short: "Deletes a target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteTargetGroupCmd).Standalone()

	vpcLattice_deleteTargetGroupCmd.Flags().String("target-group-identifier", "", "The ID or ARN of the target group.")
	vpcLattice_deleteTargetGroupCmd.MarkFlagRequired("target-group-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_deleteTargetGroupCmd)
}
