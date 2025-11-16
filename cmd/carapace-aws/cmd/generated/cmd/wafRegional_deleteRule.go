package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_deleteRuleCmd).Standalone()

		wafRegional_deleteRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_deleteRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the [Rule]() that you want to delete.")
		wafRegional_deleteRuleCmd.MarkFlagRequired("change-token")
		wafRegional_deleteRuleCmd.MarkFlagRequired("rule-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_deleteRuleCmd)
}
