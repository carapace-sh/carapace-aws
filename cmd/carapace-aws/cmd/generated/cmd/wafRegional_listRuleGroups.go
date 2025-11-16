package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listRuleGroupsCmd = &cobra.Command{
	Use:   "list-rule-groups",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listRuleGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_listRuleGroupsCmd).Standalone()

		wafRegional_listRuleGroupsCmd.Flags().String("limit", "", "Specifies the number of `RuleGroups` that you want AWS WAF to return for this request.")
		wafRegional_listRuleGroupsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `RuleGroups` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `RuleGroups`.")
	})
	wafRegionalCmd.AddCommand(wafRegional_listRuleGroupsCmd)
}
