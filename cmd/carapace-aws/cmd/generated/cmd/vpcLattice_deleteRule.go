package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "Deletes a listener rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deleteRuleCmd).Standalone()

		vpcLattice_deleteRuleCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
		vpcLattice_deleteRuleCmd.Flags().String("rule-identifier", "", "The ID or ARN of the rule.")
		vpcLattice_deleteRuleCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_deleteRuleCmd.MarkFlagRequired("listener-identifier")
		vpcLattice_deleteRuleCmd.MarkFlagRequired("rule-identifier")
		vpcLattice_deleteRuleCmd.MarkFlagRequired("service-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deleteRuleCmd)
}
