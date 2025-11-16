package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_createRuleGroupCmd = &cobra.Command{
	Use:   "create-rule-group",
	Short: "Creates the specified stateless or stateful rule group, which includes the rules for network traffic inspection, a capacity setting, and tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_createRuleGroupCmd).Standalone()

	networkFirewall_createRuleGroupCmd.Flags().Bool("analyze-rule-group", false, "Indicates whether you want Network Firewall to analyze the stateless rules in the rule group for rule behavior such as asymmetric routing.")
	networkFirewall_createRuleGroupCmd.Flags().String("capacity", "", "The maximum operating resources that this rule group can use.")
	networkFirewall_createRuleGroupCmd.Flags().String("description", "", "A description of the rule group.")
	networkFirewall_createRuleGroupCmd.Flags().Bool("dry-run", false, "Indicates whether you want Network Firewall to just check the validity of the request, rather than run the request.")
	networkFirewall_createRuleGroupCmd.Flags().String("encryption-configuration", "", "A complex type that contains settings for encryption of your rule group resources.")
	networkFirewall_createRuleGroupCmd.Flags().Bool("no-analyze-rule-group", false, "Indicates whether you want Network Firewall to analyze the stateless rules in the rule group for rule behavior such as asymmetric routing.")
	networkFirewall_createRuleGroupCmd.Flags().Bool("no-dry-run", false, "Indicates whether you want Network Firewall to just check the validity of the request, rather than run the request.")
	networkFirewall_createRuleGroupCmd.Flags().String("rule-group", "", "An object that defines the rule group rules.")
	networkFirewall_createRuleGroupCmd.Flags().String("rule-group-name", "", "The descriptive name of the rule group.")
	networkFirewall_createRuleGroupCmd.Flags().String("rules", "", "A string containing stateful rule group rules specifications in Suricata flat format, with one rule per line.")
	networkFirewall_createRuleGroupCmd.Flags().String("source-metadata", "", "A complex type that contains metadata about the rule group that your own rule group is copied from.")
	networkFirewall_createRuleGroupCmd.Flags().String("summary-configuration", "", "An object that contains a `RuleOptions` array of strings.")
	networkFirewall_createRuleGroupCmd.Flags().String("tags", "", "The key:value pairs to associate with the resource.")
	networkFirewall_createRuleGroupCmd.Flags().String("type", "", "Indicates whether the rule group is stateless or stateful.")
	networkFirewall_createRuleGroupCmd.MarkFlagRequired("capacity")
	networkFirewall_createRuleGroupCmd.Flag("no-analyze-rule-group").Hidden = true
	networkFirewall_createRuleGroupCmd.Flag("no-dry-run").Hidden = true
	networkFirewall_createRuleGroupCmd.MarkFlagRequired("rule-group-name")
	networkFirewall_createRuleGroupCmd.MarkFlagRequired("type")
	networkFirewallCmd.AddCommand(networkFirewall_createRuleGroupCmd)
}
