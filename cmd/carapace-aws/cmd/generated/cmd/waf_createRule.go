package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createRuleCmd = &cobra.Command{
	Use:   "create-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createRuleCmd).Standalone()

		waf_createRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createRuleCmd.Flags().String("metric-name", "", "A friendly name or description for the metrics for this `Rule`.")
		waf_createRuleCmd.Flags().String("name", "", "A friendly name or description of the [Rule]().")
		waf_createRuleCmd.Flags().String("tags", "", "")
		waf_createRuleCmd.MarkFlagRequired("change-token")
		waf_createRuleCmd.MarkFlagRequired("metric-name")
		waf_createRuleCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createRuleCmd)
}
