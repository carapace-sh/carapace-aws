package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getFirewallRuleGroupCmd = &cobra.Command{
	Use:   "get-firewall-rule-group",
	Short: "Retrieves the specified firewall rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getFirewallRuleGroupCmd).Standalone()

	route53resolver_getFirewallRuleGroupCmd.Flags().String("firewall-rule-group-id", "", "The unique identifier of the firewall rule group.")
	route53resolver_getFirewallRuleGroupCmd.MarkFlagRequired("firewall-rule-group-id")
	route53resolverCmd.AddCommand(route53resolver_getFirewallRuleGroupCmd)
}
