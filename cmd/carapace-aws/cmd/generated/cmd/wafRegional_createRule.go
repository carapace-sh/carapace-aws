package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createRuleCmd = &cobra.Command{
	Use:   "create-rule",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createRuleCmd).Standalone()

		wafRegional_createRuleCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_createRuleCmd.Flags().String("metric-name", "", "A friendly name or description for the metrics for this `Rule`.")
		wafRegional_createRuleCmd.Flags().String("name", "", "A friendly name or description of the [Rule]().")
		wafRegional_createRuleCmd.Flags().String("tags", "", "")
		wafRegional_createRuleCmd.MarkFlagRequired("change-token")
		wafRegional_createRuleCmd.MarkFlagRequired("metric-name")
		wafRegional_createRuleCmd.MarkFlagRequired("name")
	})
	wafRegionalCmd.AddCommand(wafRegional_createRuleCmd)
}
