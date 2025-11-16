package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getRuleCmd = &cobra.Command{
	Use:   "get-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_getRuleCmd).Standalone()

		waf_getRuleCmd.Flags().String("rule-id", "", "The `RuleId` of the [Rule]() that you want to get.")
		waf_getRuleCmd.MarkFlagRequired("rule-id")
	})
	wafCmd.AddCommand(waf_getRuleCmd)
}
