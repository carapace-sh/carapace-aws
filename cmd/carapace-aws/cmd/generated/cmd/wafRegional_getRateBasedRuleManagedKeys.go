package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getRateBasedRuleManagedKeysCmd = &cobra.Command{
	Use:   "get-rate-based-rule-managed-keys",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getRateBasedRuleManagedKeysCmd).Standalone()

	wafRegional_getRateBasedRuleManagedKeysCmd.Flags().String("next-marker", "", "A null value and not currently used.")
	wafRegional_getRateBasedRuleManagedKeysCmd.Flags().String("rule-id", "", "The `RuleId` of the [RateBasedRule]() for which you want to get a list of `ManagedKeys`.")
	wafRegional_getRateBasedRuleManagedKeysCmd.MarkFlagRequired("rule-id")
	wafRegionalCmd.AddCommand(wafRegional_getRateBasedRuleManagedKeysCmd)
}
