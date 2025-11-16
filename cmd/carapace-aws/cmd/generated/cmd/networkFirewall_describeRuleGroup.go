package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeRuleGroupCmd = &cobra.Command{
	Use:   "describe-rule-group",
	Short: "Returns the data objects for the specified rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeRuleGroupCmd).Standalone()

	networkFirewall_describeRuleGroupCmd.Flags().Bool("analyze-rule-group", false, "Indicates whether you want Network Firewall to analyze the stateless rules in the rule group for rule behavior such as asymmetric routing.")
	networkFirewall_describeRuleGroupCmd.Flags().Bool("no-analyze-rule-group", false, "Indicates whether you want Network Firewall to analyze the stateless rules in the rule group for rule behavior such as asymmetric routing.")
	networkFirewall_describeRuleGroupCmd.Flags().String("rule-group-arn", "", "The Amazon Resource Name (ARN) of the rule group.")
	networkFirewall_describeRuleGroupCmd.Flags().String("rule-group-name", "", "The descriptive name of the rule group.")
	networkFirewall_describeRuleGroupCmd.Flags().String("type", "", "Indicates whether the rule group is stateless or stateful.")
	networkFirewall_describeRuleGroupCmd.Flag("no-analyze-rule-group").Hidden = true
	networkFirewallCmd.AddCommand(networkFirewall_describeRuleGroupCmd)
}
