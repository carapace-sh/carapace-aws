package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listActivatedRulesInRuleGroupCmd = &cobra.Command{
	Use:   "list-activated-rules-in-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listActivatedRulesInRuleGroupCmd).Standalone()

	wafRegional_listActivatedRulesInRuleGroupCmd.Flags().String("limit", "", "Specifies the number of `ActivatedRules` that you want AWS WAF to return for this request.")
	wafRegional_listActivatedRulesInRuleGroupCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `ActivatedRules` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `ActivatedRules`.")
	wafRegional_listActivatedRulesInRuleGroupCmd.Flags().String("rule-group-id", "", "The `RuleGroupId` of the [RuleGroup]() for which you want to get a list of [ActivatedRule]() objects.")
	wafRegionalCmd.AddCommand(wafRegional_listActivatedRulesInRuleGroupCmd)
}
