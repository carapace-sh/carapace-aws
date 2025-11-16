package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getRateBasedRuleManagedKeysCmd = &cobra.Command{
	Use:   "get-rate-based-rule-managed-keys",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getRateBasedRuleManagedKeysCmd).Standalone()

	waf_getRateBasedRuleManagedKeysCmd.Flags().String("next-marker", "", "A null value and not currently used.")
	waf_getRateBasedRuleManagedKeysCmd.Flags().String("rule-id", "", "The `RuleId` of the [RateBasedRule]() for which you want to get a list of `ManagedKeys`.")
	waf_getRateBasedRuleManagedKeysCmd.MarkFlagRequired("rule-id")
	wafCmd.AddCommand(waf_getRateBasedRuleManagedKeysCmd)
}
