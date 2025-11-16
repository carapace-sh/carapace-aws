package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateFirewallRuleCmd = &cobra.Command{
	Use:   "update-firewall-rule",
	Short: "Updates the specified firewall rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateFirewallRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_updateFirewallRuleCmd).Standalone()

		route53resolver_updateFirewallRuleCmd.Flags().String("action", "", "The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list, or a threat in a DNS Firewall Advanced rule:")
		route53resolver_updateFirewallRuleCmd.Flags().String("block-override-dns-type", "", "The DNS record's type.")
		route53resolver_updateFirewallRuleCmd.Flags().String("block-override-domain", "", "The custom DNS record to send back in response to the query.")
		route53resolver_updateFirewallRuleCmd.Flags().String("block-override-ttl", "", "The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.")
		route53resolver_updateFirewallRuleCmd.Flags().String("block-response", "", "The way that you want DNS Firewall to block the request.")
		route53resolver_updateFirewallRuleCmd.Flags().String("confidence-threshold", "", "The confidence threshold for DNS Firewall Advanced.")
		route53resolver_updateFirewallRuleCmd.Flags().String("dns-threat-protection", "", "The type of the DNS Firewall Advanced rule.")
		route53resolver_updateFirewallRuleCmd.Flags().String("firewall-domain-list-id", "", "The ID of the domain list to use in the rule.")
		route53resolver_updateFirewallRuleCmd.Flags().String("firewall-domain-redirection-action", "", "How you want the the rule to evaluate DNS redirection in the DNS redirection chain, such as CNAME or DNAME.")
		route53resolver_updateFirewallRuleCmd.Flags().String("firewall-rule-group-id", "", "The unique identifier of the firewall rule group for the rule.")
		route53resolver_updateFirewallRuleCmd.Flags().String("firewall-threat-protection-id", "", "The DNS Firewall Advanced rule ID.")
		route53resolver_updateFirewallRuleCmd.Flags().String("name", "", "The name of the rule.")
		route53resolver_updateFirewallRuleCmd.Flags().String("priority", "", "The setting that determines the processing order of the rule in the rule group.")
		route53resolver_updateFirewallRuleCmd.Flags().String("qtype", "", "The DNS query type you want the rule to evaluate.")
		route53resolver_updateFirewallRuleCmd.MarkFlagRequired("firewall-rule-group-id")
	})
	route53resolverCmd.AddCommand(route53resolver_updateFirewallRuleCmd)
}
