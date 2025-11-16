package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateRuleCmd = &cobra.Command{
	Use:   "update-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateRuleCmd).Standalone()

	waf_updateRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_updateRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the `Rule` that you want to update.")
	waf_updateRuleCmd.Flags().String("updates", "", "An array of `RuleUpdate` objects that you want to insert into or delete from a [Rule]().")
	waf_updateRuleCmd.MarkFlagRequired("change-token")
	waf_updateRuleCmd.MarkFlagRequired("rule-id")
	waf_updateRuleCmd.MarkFlagRequired("updates")
	wafCmd.AddCommand(waf_updateRuleCmd)
}
