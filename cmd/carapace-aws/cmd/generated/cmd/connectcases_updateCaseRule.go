package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_updateCaseRuleCmd = &cobra.Command{
	Use:   "update-case-rule",
	Short: "Updates a case rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_updateCaseRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_updateCaseRuleCmd).Standalone()

		connectcases_updateCaseRuleCmd.Flags().String("case-rule-id", "", "Unique identifier of a case rule.")
		connectcases_updateCaseRuleCmd.Flags().String("description", "", "Description of a case rule.")
		connectcases_updateCaseRuleCmd.Flags().String("domain-id", "", "Unique identifier of a Cases domain.")
		connectcases_updateCaseRuleCmd.Flags().String("name", "", "Name of the case rule.")
		connectcases_updateCaseRuleCmd.Flags().String("rule", "", "Represents what rule type should take place, under what conditions.")
		connectcases_updateCaseRuleCmd.MarkFlagRequired("case-rule-id")
		connectcases_updateCaseRuleCmd.MarkFlagRequired("domain-id")
	})
	connectcasesCmd.AddCommand(connectcases_updateCaseRuleCmd)
}
