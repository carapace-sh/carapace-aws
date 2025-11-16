package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateRuleGroupCmd = &cobra.Command{
	Use:   "update-rule-group",
	Short: "Updates the rule settings for the specified rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkFirewall_updateRuleGroupCmd).Standalone()

		networkFirewall_updateRuleGroupCmd.Flags().Bool("analyze-rule-group", false, "Indicates whether you want Network Firewall to analyze the stateless rules in the rule group for rule behavior such as asymmetric routing.")
		networkFirewall_updateRuleGroupCmd.Flags().String("description", "", "A description of the rule group.")
		networkFirewall_updateRuleGroupCmd.Flags().Bool("dry-run", false, "Indicates whether you want Network Firewall to just check the validity of the request, rather than run the request.")
		networkFirewall_updateRuleGroupCmd.Flags().String("encryption-configuration", "", "A complex type that contains settings for encryption of your rule group resources.")
		networkFirewall_updateRuleGroupCmd.Flags().Bool("no-analyze-rule-group", false, "Indicates whether you want Network Firewall to analyze the stateless rules in the rule group for rule behavior such as asymmetric routing.")
		networkFirewall_updateRuleGroupCmd.Flags().Bool("no-dry-run", false, "Indicates whether you want Network Firewall to just check the validity of the request, rather than run the request.")
		networkFirewall_updateRuleGroupCmd.Flags().String("rule-group", "", "An object that defines the rule group rules.")
		networkFirewall_updateRuleGroupCmd.Flags().String("rule-group-arn", "", "The Amazon Resource Name (ARN) of the rule group.")
		networkFirewall_updateRuleGroupCmd.Flags().String("rule-group-name", "", "The descriptive name of the rule group.")
		networkFirewall_updateRuleGroupCmd.Flags().String("rules", "", "A string containing stateful rule group rules specifications in Suricata flat format, with one rule per line.")
		networkFirewall_updateRuleGroupCmd.Flags().String("source-metadata", "", "A complex type that contains metadata about the rule group that your own rule group is copied from.")
		networkFirewall_updateRuleGroupCmd.Flags().String("summary-configuration", "", "Updates the selected summary configuration for a rule group.")
		networkFirewall_updateRuleGroupCmd.Flags().String("type", "", "Indicates whether the rule group is stateless or stateful.")
		networkFirewall_updateRuleGroupCmd.Flags().String("update-token", "", "A token used for optimistic locking.")
		networkFirewall_updateRuleGroupCmd.Flag("no-analyze-rule-group").Hidden = true
		networkFirewall_updateRuleGroupCmd.Flag("no-dry-run").Hidden = true
		networkFirewall_updateRuleGroupCmd.MarkFlagRequired("update-token")
	})
	networkFirewallCmd.AddCommand(networkFirewall_updateRuleGroupCmd)
}
