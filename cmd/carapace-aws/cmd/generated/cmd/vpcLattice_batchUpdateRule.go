package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_batchUpdateRuleCmd = &cobra.Command{
	Use:   "batch-update-rule",
	Short: "Updates the listener rules in a batch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_batchUpdateRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_batchUpdateRuleCmd).Standalone()

		vpcLattice_batchUpdateRuleCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
		vpcLattice_batchUpdateRuleCmd.Flags().String("rules", "", "The rules for the specified listener.")
		vpcLattice_batchUpdateRuleCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_batchUpdateRuleCmd.MarkFlagRequired("listener-identifier")
		vpcLattice_batchUpdateRuleCmd.MarkFlagRequired("rules")
		vpcLattice_batchUpdateRuleCmd.MarkFlagRequired("service-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_batchUpdateRuleCmd)
}
