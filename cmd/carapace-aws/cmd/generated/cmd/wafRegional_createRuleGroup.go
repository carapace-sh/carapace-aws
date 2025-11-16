package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createRuleGroupCmd = &cobra.Command{
	Use:   "create-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createRuleGroupCmd).Standalone()

		wafRegional_createRuleGroupCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_createRuleGroupCmd.Flags().String("metric-name", "", "A friendly name or description for the metrics for this `RuleGroup`.")
		wafRegional_createRuleGroupCmd.Flags().String("name", "", "A friendly name or description of the [RuleGroup]().")
		wafRegional_createRuleGroupCmd.Flags().String("tags", "", "")
		wafRegional_createRuleGroupCmd.MarkFlagRequired("change-token")
		wafRegional_createRuleGroupCmd.MarkFlagRequired("metric-name")
		wafRegional_createRuleGroupCmd.MarkFlagRequired("name")
	})
	wafRegionalCmd.AddCommand(wafRegional_createRuleGroupCmd)
}
