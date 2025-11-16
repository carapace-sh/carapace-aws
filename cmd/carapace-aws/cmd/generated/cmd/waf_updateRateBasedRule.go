package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateRateBasedRuleCmd = &cobra.Command{
	Use:   "update-rate-based-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateRateBasedRuleCmd).Standalone()

	waf_updateRateBasedRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_updateRateBasedRuleCmd.Flags().String("rate-limit", "", "The maximum number of requests, which have an identical value in the field specified by the `RateKey`, allowed in a five-minute period.")
	waf_updateRateBasedRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the `RateBasedRule` that you want to update.")
	waf_updateRateBasedRuleCmd.Flags().String("updates", "", "An array of `RuleUpdate` objects that you want to insert into or delete from a [RateBasedRule]().")
	waf_updateRateBasedRuleCmd.MarkFlagRequired("change-token")
	waf_updateRateBasedRuleCmd.MarkFlagRequired("rate-limit")
	waf_updateRateBasedRuleCmd.MarkFlagRequired("rule-id")
	waf_updateRateBasedRuleCmd.MarkFlagRequired("updates")
	wafCmd.AddCommand(waf_updateRateBasedRuleCmd)
}
