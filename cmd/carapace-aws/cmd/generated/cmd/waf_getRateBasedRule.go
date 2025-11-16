package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getRateBasedRuleCmd = &cobra.Command{
	Use:   "get-rate-based-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getRateBasedRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_getRateBasedRuleCmd).Standalone()

		waf_getRateBasedRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the [RateBasedRule]() that you want to get.")
		waf_getRateBasedRuleCmd.MarkFlagRequired("rule-id")
	})
	wafCmd.AddCommand(waf_getRateBasedRuleCmd)
}
