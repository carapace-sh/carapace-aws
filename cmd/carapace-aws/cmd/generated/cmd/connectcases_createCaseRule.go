package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_createCaseRuleCmd = &cobra.Command{
	Use:   "create-case-rule",
	Short: "Creates a new case rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_createCaseRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_createCaseRuleCmd).Standalone()

		connectcases_createCaseRuleCmd.Flags().String("description", "", "The description of a case rule.")
		connectcases_createCaseRuleCmd.Flags().String("domain-id", "", "Unique identifier of a Cases domain.")
		connectcases_createCaseRuleCmd.Flags().String("name", "", "Name of the case rule.")
		connectcases_createCaseRuleCmd.Flags().String("rule", "", "Represents what rule type should take place, under what conditions.")
		connectcases_createCaseRuleCmd.MarkFlagRequired("domain-id")
		connectcases_createCaseRuleCmd.MarkFlagRequired("name")
		connectcases_createCaseRuleCmd.MarkFlagRequired("rule")
	})
	connectcasesCmd.AddCommand(connectcases_createCaseRuleCmd)
}
