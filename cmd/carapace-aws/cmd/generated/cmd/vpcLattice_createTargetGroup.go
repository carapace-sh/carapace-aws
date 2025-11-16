package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createTargetGroupCmd = &cobra.Command{
	Use:   "create-target-group",
	Short: "Creates a target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createTargetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_createTargetGroupCmd).Standalone()

		vpcLattice_createTargetGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		vpcLattice_createTargetGroupCmd.Flags().String("config", "", "The target group configuration.")
		vpcLattice_createTargetGroupCmd.Flags().String("name", "", "The name of the target group.")
		vpcLattice_createTargetGroupCmd.Flags().String("tags", "", "The tags for the target group.")
		vpcLattice_createTargetGroupCmd.Flags().String("type", "", "The type of target group.")
		vpcLattice_createTargetGroupCmd.MarkFlagRequired("name")
		vpcLattice_createTargetGroupCmd.MarkFlagRequired("type")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_createTargetGroupCmd)
}
