package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateTargetGroupCmd = &cobra.Command{
	Use:   "update-target-group",
	Short: "Updates the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateTargetGroupCmd).Standalone()

	vpcLattice_updateTargetGroupCmd.Flags().String("health-check", "", "The health check configuration.")
	vpcLattice_updateTargetGroupCmd.Flags().String("target-group-identifier", "", "The ID or ARN of the target group.")
	vpcLattice_updateTargetGroupCmd.MarkFlagRequired("health-check")
	vpcLattice_updateTargetGroupCmd.MarkFlagRequired("target-group-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_updateTargetGroupCmd)
}
