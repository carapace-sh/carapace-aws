package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listTargetsCmd = &cobra.Command{
	Use:   "list-targets",
	Short: "Lists the targets for the target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listTargetsCmd).Standalone()

	vpcLattice_listTargetsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	vpcLattice_listTargetsCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	vpcLattice_listTargetsCmd.Flags().String("target-group-identifier", "", "The ID or ARN of the target group.")
	vpcLattice_listTargetsCmd.Flags().String("targets", "", "The targets.")
	vpcLattice_listTargetsCmd.MarkFlagRequired("target-group-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_listTargetsCmd)
}
