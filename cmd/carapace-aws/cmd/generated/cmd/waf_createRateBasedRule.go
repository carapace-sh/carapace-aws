package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createRateBasedRuleCmd = &cobra.Command{
	Use:   "create-rate-based-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createRateBasedRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createRateBasedRuleCmd).Standalone()

		waf_createRateBasedRuleCmd.Flags().String("change-token", "", "The `ChangeToken` that you used to submit the `CreateRateBasedRule` request.")
		waf_createRateBasedRuleCmd.Flags().String("metric-name", "", "A friendly name or description for the metrics for this `RateBasedRule`.")
		waf_createRateBasedRuleCmd.Flags().String("name", "", "A friendly name or description of the [RateBasedRule]().")
		waf_createRateBasedRuleCmd.Flags().String("rate-key", "", "The field that AWS WAF uses to determine if requests are likely arriving from a single source and thus subject to rate monitoring.")
		waf_createRateBasedRuleCmd.Flags().String("rate-limit", "", "The maximum number of requests, which have an identical value in the field that is specified by `RateKey`, allowed in a five-minute period.")
		waf_createRateBasedRuleCmd.Flags().String("tags", "", "")
		waf_createRateBasedRuleCmd.MarkFlagRequired("change-token")
		waf_createRateBasedRuleCmd.MarkFlagRequired("metric-name")
		waf_createRateBasedRuleCmd.MarkFlagRequired("name")
		waf_createRateBasedRuleCmd.MarkFlagRequired("rate-key")
		waf_createRateBasedRuleCmd.MarkFlagRequired("rate-limit")
	})
	wafCmd.AddCommand(waf_createRateBasedRuleCmd)
}
