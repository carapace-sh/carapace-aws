package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateFirewallRuleGroupAssociationCmd = &cobra.Command{
	Use:   "update-firewall-rule-group-association",
	Short: "Changes the association of a [FirewallRuleGroup]() with a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateFirewallRuleGroupAssociationCmd).Standalone()

	route53resolver_updateFirewallRuleGroupAssociationCmd.Flags().String("firewall-rule-group-association-id", "", "The identifier of the [FirewallRuleGroupAssociation]().")
	route53resolver_updateFirewallRuleGroupAssociationCmd.Flags().String("mutation-protection", "", "If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.")
	route53resolver_updateFirewallRuleGroupAssociationCmd.Flags().String("name", "", "The name of the rule group association.")
	route53resolver_updateFirewallRuleGroupAssociationCmd.Flags().String("priority", "", "The setting that determines the processing order of the rule group among the rule groups that you associate with the specified VPC.")
	route53resolver_updateFirewallRuleGroupAssociationCmd.MarkFlagRequired("firewall-rule-group-association-id")
	route53resolverCmd.AddCommand(route53resolver_updateFirewallRuleGroupAssociationCmd)
}
