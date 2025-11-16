package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listActivatedRulesInRuleGroupCmd = &cobra.Command{
	Use:   "list-activated-rules-in-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listActivatedRulesInRuleGroupCmd).Standalone()

	waf_listActivatedRulesInRuleGroupCmd.Flags().String("limit", "", "Specifies the number of `ActivatedRules` that you want AWS WAF to return for this request.")
	waf_listActivatedRulesInRuleGroupCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `ActivatedRules` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `ActivatedRules`.")
	waf_listActivatedRulesInRuleGroupCmd.Flags().String("rule-group-id", "", "The `RuleGroupId` of the [RuleGroup]() for which you want to get a list of [ActivatedRule]() objects.")
	wafCmd.AddCommand(waf_listActivatedRulesInRuleGroupCmd)
}
