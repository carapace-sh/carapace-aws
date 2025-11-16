package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listRuleGroupsCmd = &cobra.Command{
	Use:   "list-rule-groups",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listRuleGroupsCmd).Standalone()

	waf_listRuleGroupsCmd.Flags().String("limit", "", "Specifies the number of `RuleGroups` that you want AWS WAF to return for this request.")
	waf_listRuleGroupsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `RuleGroups` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `RuleGroups`.")
	wafCmd.AddCommand(waf_listRuleGroupsCmd)
}
