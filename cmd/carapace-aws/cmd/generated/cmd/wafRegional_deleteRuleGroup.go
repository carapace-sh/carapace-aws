package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteRuleGroupCmd = &cobra.Command{
	Use:   "delete-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_deleteRuleGroupCmd).Standalone()

		wafRegional_deleteRuleGroupCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_deleteRuleGroupCmd.Flags().String("rule-group-id", "", "The `RuleGroupId` of the [RuleGroup]() that you want to delete.")
		wafRegional_deleteRuleGroupCmd.MarkFlagRequired("change-token")
		wafRegional_deleteRuleGroupCmd.MarkFlagRequired("rule-group-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_deleteRuleGroupCmd)
}
