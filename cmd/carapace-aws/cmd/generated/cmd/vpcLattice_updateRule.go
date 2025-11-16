package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateRuleCmd = &cobra.Command{
	Use:   "update-rule",
	Short: "Updates a specified rule for the listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateRuleCmd).Standalone()

	vpcLattice_updateRuleCmd.Flags().String("action", "", "Information about the action for the specified listener rule.")
	vpcLattice_updateRuleCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
	vpcLattice_updateRuleCmd.Flags().String("match", "", "The rule match.")
	vpcLattice_updateRuleCmd.Flags().String("priority", "", "The rule priority.")
	vpcLattice_updateRuleCmd.Flags().String("rule-identifier", "", "The ID or ARN of the rule.")
	vpcLattice_updateRuleCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
	vpcLattice_updateRuleCmd.MarkFlagRequired("listener-identifier")
	vpcLattice_updateRuleCmd.MarkFlagRequired("rule-identifier")
	vpcLattice_updateRuleCmd.MarkFlagRequired("service-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_updateRuleCmd)
}
