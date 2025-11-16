package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_associateFirewallRuleGroupCmd = &cobra.Command{
	Use:   "associate-firewall-rule-group",
	Short: "Associates a [FirewallRuleGroup]() with a VPC, to provide DNS filtering for the VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_associateFirewallRuleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_associateFirewallRuleGroupCmd).Standalone()

		route53resolver_associateFirewallRuleGroupCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed requests to be retried without the risk of running the operation twice.")
		route53resolver_associateFirewallRuleGroupCmd.Flags().String("firewall-rule-group-id", "", "The unique identifier of the firewall rule group.")
		route53resolver_associateFirewallRuleGroupCmd.Flags().String("mutation-protection", "", "If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.")
		route53resolver_associateFirewallRuleGroupCmd.Flags().String("name", "", "A name that lets you identify the association, to manage and use it.")
		route53resolver_associateFirewallRuleGroupCmd.Flags().String("priority", "", "The setting that determines the processing order of the rule group among the rule groups that you associate with the specified VPC.")
		route53resolver_associateFirewallRuleGroupCmd.Flags().String("tags", "", "A list of the tag keys and values that you want to associate with the rule group association.")
		route53resolver_associateFirewallRuleGroupCmd.Flags().String("vpc-id", "", "The unique identifier of the VPC that you want to associate with the rule group.")
		route53resolver_associateFirewallRuleGroupCmd.MarkFlagRequired("creator-request-id")
		route53resolver_associateFirewallRuleGroupCmd.MarkFlagRequired("firewall-rule-group-id")
		route53resolver_associateFirewallRuleGroupCmd.MarkFlagRequired("name")
		route53resolver_associateFirewallRuleGroupCmd.MarkFlagRequired("priority")
		route53resolver_associateFirewallRuleGroupCmd.MarkFlagRequired("vpc-id")
	})
	route53resolverCmd.AddCommand(route53resolver_associateFirewallRuleGroupCmd)
}
