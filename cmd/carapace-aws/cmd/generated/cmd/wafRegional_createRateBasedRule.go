package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createRateBasedRuleCmd = &cobra.Command{
	Use:   "create-rate-based-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createRateBasedRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createRateBasedRuleCmd).Standalone()

		wafRegional_createRateBasedRuleCmd.Flags().String("change-token", "", "The `ChangeToken` that you used to submit the `CreateRateBasedRule` request.")
		wafRegional_createRateBasedRuleCmd.Flags().String("metric-name", "", "A friendly name or description for the metrics for this `RateBasedRule`.")
		wafRegional_createRateBasedRuleCmd.Flags().String("name", "", "A friendly name or description of the [RateBasedRule]().")
		wafRegional_createRateBasedRuleCmd.Flags().String("rate-key", "", "The field that AWS WAF uses to determine if requests are likely arriving from a single source and thus subject to rate monitoring.")
		wafRegional_createRateBasedRuleCmd.Flags().String("rate-limit", "", "The maximum number of requests, which have an identical value in the field that is specified by `RateKey`, allowed in a five-minute period.")
		wafRegional_createRateBasedRuleCmd.Flags().String("tags", "", "")
		wafRegional_createRateBasedRuleCmd.MarkFlagRequired("change-token")
		wafRegional_createRateBasedRuleCmd.MarkFlagRequired("metric-name")
		wafRegional_createRateBasedRuleCmd.MarkFlagRequired("name")
		wafRegional_createRateBasedRuleCmd.MarkFlagRequired("rate-key")
		wafRegional_createRateBasedRuleCmd.MarkFlagRequired("rate-limit")
	})
	wafRegionalCmd.AddCommand(wafRegional_createRateBasedRuleCmd)
}
