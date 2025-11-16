package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateRuleGroupCmd = &cobra.Command{
	Use:   "update-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateRuleGroupCmd).Standalone()

	wafRegional_updateRuleGroupCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_updateRuleGroupCmd.Flags().String("rule-group-id", "", "The `RuleGroupId` of the [RuleGroup]() that you want to update.")
	wafRegional_updateRuleGroupCmd.Flags().String("updates", "", "An array of `RuleGroupUpdate` objects that you want to insert into or delete from a [RuleGroup]().")
	wafRegional_updateRuleGroupCmd.MarkFlagRequired("change-token")
	wafRegional_updateRuleGroupCmd.MarkFlagRequired("rule-group-id")
	wafRegional_updateRuleGroupCmd.MarkFlagRequired("updates")
	wafRegionalCmd.AddCommand(wafRegional_updateRuleGroupCmd)
}
