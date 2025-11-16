package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getFirewallRuleGroupAssociationCmd = &cobra.Command{
	Use:   "get-firewall-rule-group-association",
	Short: "Retrieves a firewall rule group association, which enables DNS filtering for a VPC with one rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getFirewallRuleGroupAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_getFirewallRuleGroupAssociationCmd).Standalone()

		route53resolver_getFirewallRuleGroupAssociationCmd.Flags().String("firewall-rule-group-association-id", "", "The identifier of the [FirewallRuleGroupAssociation]().")
		route53resolver_getFirewallRuleGroupAssociationCmd.MarkFlagRequired("firewall-rule-group-association-id")
	})
	route53resolverCmd.AddCommand(route53resolver_getFirewallRuleGroupAssociationCmd)
}
