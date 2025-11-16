package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createRuleGroupCmd = &cobra.Command{
	Use:   "create-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createRuleGroupCmd).Standalone()

		waf_createRuleGroupCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createRuleGroupCmd.Flags().String("metric-name", "", "A friendly name or description for the metrics for this `RuleGroup`.")
		waf_createRuleGroupCmd.Flags().String("name", "", "A friendly name or description of the [RuleGroup]().")
		waf_createRuleGroupCmd.Flags().String("tags", "", "")
		waf_createRuleGroupCmd.MarkFlagRequired("change-token")
		waf_createRuleGroupCmd.MarkFlagRequired("metric-name")
		waf_createRuleGroupCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createRuleGroupCmd)
}
