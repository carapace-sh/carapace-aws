package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateRuleGroupCmd = &cobra.Command{
	Use:   "update-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_updateRuleGroupCmd).Standalone()

		waf_updateRuleGroupCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_updateRuleGroupCmd.Flags().String("rule-group-id", "", "The `RuleGroupId` of the [RuleGroup]() that you want to update.")
		waf_updateRuleGroupCmd.Flags().String("updates", "", "An array of `RuleGroupUpdate` objects that you want to insert into or delete from a [RuleGroup]().")
		waf_updateRuleGroupCmd.MarkFlagRequired("change-token")
		waf_updateRuleGroupCmd.MarkFlagRequired("rule-group-id")
		waf_updateRuleGroupCmd.MarkFlagRequired("updates")
	})
	wafCmd.AddCommand(waf_updateRuleGroupCmd)
}
