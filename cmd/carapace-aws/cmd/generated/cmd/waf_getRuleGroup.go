package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getRuleGroupCmd = &cobra.Command{
	Use:   "get-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_getRuleGroupCmd).Standalone()

		waf_getRuleGroupCmd.Flags().String("rule-group-id", "", "The `RuleGroupId` of the [RuleGroup]() that you want to get.")
		waf_getRuleGroupCmd.MarkFlagRequired("rule-group-id")
	})
	wafCmd.AddCommand(waf_getRuleGroupCmd)
}
