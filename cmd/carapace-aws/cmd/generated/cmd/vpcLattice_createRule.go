package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createRuleCmd = &cobra.Command{
	Use:   "create-rule",
	Short: "Creates a listener rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_createRuleCmd).Standalone()

		vpcLattice_createRuleCmd.Flags().String("action", "", "The action for the default rule.")
		vpcLattice_createRuleCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		vpcLattice_createRuleCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
		vpcLattice_createRuleCmd.Flags().String("match", "", "The rule match.")
		vpcLattice_createRuleCmd.Flags().String("name", "", "The name of the rule.")
		vpcLattice_createRuleCmd.Flags().String("priority", "", "The priority assigned to the rule.")
		vpcLattice_createRuleCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_createRuleCmd.Flags().String("tags", "", "The tags for the rule.")
		vpcLattice_createRuleCmd.MarkFlagRequired("action")
		vpcLattice_createRuleCmd.MarkFlagRequired("listener-identifier")
		vpcLattice_createRuleCmd.MarkFlagRequired("match")
		vpcLattice_createRuleCmd.MarkFlagRequired("name")
		vpcLattice_createRuleCmd.MarkFlagRequired("priority")
		vpcLattice_createRuleCmd.MarkFlagRequired("service-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_createRuleCmd)
}
