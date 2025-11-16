package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_deleteFirewallRuleCmd = &cobra.Command{
	Use:   "delete-firewall-rule",
	Short: "Deletes the specified firewall rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_deleteFirewallRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_deleteFirewallRuleCmd).Standalone()

		route53resolver_deleteFirewallRuleCmd.Flags().String("firewall-domain-list-id", "", "The ID of the domain list that's used in the rule.")
		route53resolver_deleteFirewallRuleCmd.Flags().String("firewall-rule-group-id", "", "The unique identifier of the firewall rule group that you want to delete the rule from.")
		route53resolver_deleteFirewallRuleCmd.Flags().String("firewall-threat-protection-id", "", "The ID that is created for a DNS Firewall Advanced rule.")
		route53resolver_deleteFirewallRuleCmd.Flags().String("qtype", "", "The DNS query type that the rule you are deleting evaluates.")
		route53resolver_deleteFirewallRuleCmd.MarkFlagRequired("firewall-rule-group-id")
	})
	route53resolverCmd.AddCommand(route53resolver_deleteFirewallRuleCmd)
}
