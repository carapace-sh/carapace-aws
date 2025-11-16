package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getRuleGroupCmd = &cobra.Command{
	Use:   "get-rule-group",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getRuleGroupCmd).Standalone()

	wafRegional_getRuleGroupCmd.Flags().String("rule-group-id", "", "The `RuleGroupId` of the [RuleGroup]() that you want to get.")
	wafRegional_getRuleGroupCmd.MarkFlagRequired("rule-group-id")
	wafRegionalCmd.AddCommand(wafRegional_getRuleGroupCmd)
}
