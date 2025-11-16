package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getRateBasedRuleCmd = &cobra.Command{
	Use:   "get-rate-based-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getRateBasedRuleCmd).Standalone()

	wafRegional_getRateBasedRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the [RateBasedRule]() that you want to get.")
	wafRegional_getRateBasedRuleCmd.MarkFlagRequired("rule-id")
	wafRegionalCmd.AddCommand(wafRegional_getRateBasedRuleCmd)
}
