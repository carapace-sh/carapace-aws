package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_disassociateFirewallRuleGroupCmd = &cobra.Command{
	Use:   "disassociate-firewall-rule-group",
	Short: "Disassociates a [FirewallRuleGroup]() from a VPC, to remove DNS filtering from the VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_disassociateFirewallRuleGroupCmd).Standalone()

	route53resolver_disassociateFirewallRuleGroupCmd.Flags().String("firewall-rule-group-association-id", "", "The identifier of the [FirewallRuleGroupAssociation]().")
	route53resolver_disassociateFirewallRuleGroupCmd.MarkFlagRequired("firewall-rule-group-association-id")
	route53resolverCmd.AddCommand(route53resolver_disassociateFirewallRuleGroupCmd)
}
