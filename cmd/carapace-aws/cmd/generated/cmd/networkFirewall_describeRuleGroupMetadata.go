package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_describeRuleGroupMetadataCmd = &cobra.Command{
	Use:   "describe-rule-group-metadata",
	Short: "High-level information about a rule group, returned by operations like create and describe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_describeRuleGroupMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_describeRuleGroupMetadataCmd).Standalone()

		networkFirewall_describeRuleGroupMetadataCmd.Flags().String("rule-group-arn", "", "The descriptive name of the rule group.")
		networkFirewall_describeRuleGroupMetadataCmd.Flags().String("rule-group-name", "", "The descriptive name of the rule group.")
		networkFirewall_describeRuleGroupMetadataCmd.Flags().String("type", "", "Indicates whether the rule group is stateless or stateful.")
	})
	networkFirewallCmd.AddCommand(networkFirewall_describeRuleGroupMetadataCmd)
}
