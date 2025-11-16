package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteRuleGroupCmd = &cobra.Command{
	Use:   "delete-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteRuleGroupCmd).Standalone()

		waf_deleteRuleGroupCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteRuleGroupCmd.Flags().String("rule-group-id", "", "The `RuleGroupId` of the [RuleGroup]() that you want to delete.")
		waf_deleteRuleGroupCmd.MarkFlagRequired("change-token")
		waf_deleteRuleGroupCmd.MarkFlagRequired("rule-group-id")
	})
	wafCmd.AddCommand(waf_deleteRuleGroupCmd)
}
