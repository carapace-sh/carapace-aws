package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getRuleCmd = &cobra.Command{
	Use:   "get-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getRuleCmd).Standalone()

	wafRegional_getRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the [Rule]() that you want to get.")
	wafRegional_getRuleCmd.MarkFlagRequired("rule-id")
	wafRegionalCmd.AddCommand(wafRegional_getRuleCmd)
}
