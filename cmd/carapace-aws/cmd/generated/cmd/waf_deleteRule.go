package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteRuleCmd).Standalone()

		waf_deleteRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the [Rule]() that you want to delete.")
		waf_deleteRuleCmd.MarkFlagRequired("change-token")
		waf_deleteRuleCmd.MarkFlagRequired("rule-id")
	})
	wafCmd.AddCommand(waf_deleteRuleCmd)
}
