package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteRateBasedRuleCmd = &cobra.Command{
	Use:   "delete-rate-based-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteRateBasedRuleCmd).Standalone()

	wafRegional_deleteRateBasedRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_deleteRateBasedRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the [RateBasedRule]() that you want to delete.")
	wafRegional_deleteRateBasedRuleCmd.MarkFlagRequired("change-token")
	wafRegional_deleteRateBasedRuleCmd.MarkFlagRequired("rule-id")
	wafRegionalCmd.AddCommand(wafRegional_deleteRateBasedRuleCmd)
}
