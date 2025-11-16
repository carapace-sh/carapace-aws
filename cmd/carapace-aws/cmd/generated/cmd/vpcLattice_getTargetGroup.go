package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getTargetGroupCmd = &cobra.Command{
	Use:   "get-target-group",
	Short: "Retrieves information about the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getTargetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_getTargetGroupCmd).Standalone()

		vpcLattice_getTargetGroupCmd.Flags().String("target-group-identifier", "", "The ID or ARN of the target group.")
		vpcLattice_getTargetGroupCmd.MarkFlagRequired("target-group-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_getTargetGroupCmd)
}
