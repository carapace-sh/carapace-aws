package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_deleteRuleGroupCmd = &cobra.Command{
	Use:   "delete-rule-group",
	Short: "Deletes the specified [RuleGroup]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_deleteRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_deleteRuleGroupCmd).Standalone()

		networkFirewall_deleteRuleGroupCmd.Flags().String("rule-group-arn", "", "The Amazon Resource Name (ARN) of the rule group.")
		networkFirewall_deleteRuleGroupCmd.Flags().String("rule-group-name", "", "The descriptive name of the rule group.")
		networkFirewall_deleteRuleGroupCmd.Flags().String("type", "", "Indicates whether the rule group is stateless or stateful.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_deleteRuleGroupCmd)
}
