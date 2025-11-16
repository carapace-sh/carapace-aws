package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeRuleGroupSummaryCmd = &cobra.Command{
	Use:   "describe-rule-group-summary",
	Short: "Returns detailed information for a stateful rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeRuleGroupSummaryCmd).Standalone()

	networkFirewall_describeRuleGroupSummaryCmd.Flags().String("rule-group-arn", "", "Required.")
	networkFirewall_describeRuleGroupSummaryCmd.Flags().String("rule-group-name", "", "The descriptive name of the rule group.")
	networkFirewall_describeRuleGroupSummaryCmd.Flags().String("type", "", "The type of rule group you want a summary for.")
	networkFirewallCmd.AddCommand(networkFirewall_describeRuleGroupSummaryCmd)
}
