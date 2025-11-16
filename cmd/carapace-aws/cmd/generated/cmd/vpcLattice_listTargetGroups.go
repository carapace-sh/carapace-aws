package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listTargetGroupsCmd = &cobra.Command{
	Use:   "list-target-groups",
	Short: "Lists your target groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listTargetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_listTargetGroupsCmd).Standalone()

		vpcLattice_listTargetGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		vpcLattice_listTargetGroupsCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
		vpcLattice_listTargetGroupsCmd.Flags().String("target-group-type", "", "The target group type.")
		vpcLattice_listTargetGroupsCmd.Flags().String("vpc-identifier", "", "The ID or ARN of the VPC.")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_listTargetGroupsCmd)
}
