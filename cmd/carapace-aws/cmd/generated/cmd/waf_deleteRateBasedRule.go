package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteRateBasedRuleCmd = &cobra.Command{
	Use:   "delete-rate-based-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteRateBasedRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteRateBasedRuleCmd).Standalone()

		waf_deleteRateBasedRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteRateBasedRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the [RateBasedRule]() that you want to delete.")
		waf_deleteRateBasedRuleCmd.MarkFlagRequired("change-token")
		waf_deleteRateBasedRuleCmd.MarkFlagRequired("rule-id")
	})
	wafCmd.AddCommand(waf_deleteRateBasedRuleCmd)
}
