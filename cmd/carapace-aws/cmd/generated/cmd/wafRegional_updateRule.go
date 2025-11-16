package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateRuleCmd = &cobra.Command{
	Use:   "update-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateRuleCmd).Standalone()

	wafRegional_updateRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_updateRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the `Rule` that you want to update.")
	wafRegional_updateRuleCmd.Flags().String("updates", "", "An array of `RuleUpdate` objects that you want to insert into or delete from a [Rule]().")
	wafRegional_updateRuleCmd.MarkFlagRequired("change-token")
	wafRegional_updateRuleCmd.MarkFlagRequired("rule-id")
	wafRegional_updateRuleCmd.MarkFlagRequired("updates")
	wafRegionalCmd.AddCommand(wafRegional_updateRuleCmd)
}
