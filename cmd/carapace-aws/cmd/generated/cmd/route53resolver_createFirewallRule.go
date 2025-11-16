package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_createFirewallRuleCmd = &cobra.Command{
	Use:   "create-firewall-rule",
	Short: "Creates a single DNS Firewall rule in the specified rule group, using the specified domain list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_createFirewallRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_createFirewallRuleCmd).Standalone()

		route53resolver_createFirewallRuleCmd.Flags().String("action", "", "The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list, or a threat in a DNS Firewall Advanced rule:")
		route53resolver_createFirewallRuleCmd.Flags().String("block-override-dns-type", "", "The DNS record's type.")
		route53resolver_createFirewallRuleCmd.Flags().String("block-override-domain", "", "The custom DNS record to send back in response to the query.")
		route53resolver_createFirewallRuleCmd.Flags().String("block-override-ttl", "", "The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.")
		route53resolver_createFirewallRuleCmd.Flags().String("block-response", "", "The way that you want DNS Firewall to block the request, used with the rule action setting `BLOCK`.")
		route53resolver_createFirewallRuleCmd.Flags().String("confidence-threshold", "", "The confidence threshold for DNS Firewall Advanced.")
		route53resolver_createFirewallRuleCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows you to retry failed requests without the risk of running the operation twice.")
		route53resolver_createFirewallRuleCmd.Flags().String("dns-threat-protection", "", "Use to create a DNS Firewall Advanced rule.")
		route53resolver_createFirewallRuleCmd.Flags().String("firewall-domain-list-id", "", "The ID of the domain list that you want to use in the rule.")
		route53resolver_createFirewallRuleCmd.Flags().String("firewall-domain-redirection-action", "", "How you want the the rule to evaluate DNS redirection in the DNS redirection chain, such as CNAME or DNAME.")
		route53resolver_createFirewallRuleCmd.Flags().String("firewall-rule-group-id", "", "The unique identifier of the firewall rule group where you want to create the rule.")
		route53resolver_createFirewallRuleCmd.Flags().String("name", "", "A name that lets you identify the rule in the rule group.")
		route53resolver_createFirewallRuleCmd.Flags().String("priority", "", "The setting that determines the processing order of the rule in the rule group.")
		route53resolver_createFirewallRuleCmd.Flags().String("qtype", "", "The DNS query type you want the rule to evaluate.")
		route53resolver_createFirewallRuleCmd.MarkFlagRequired("action")
		route53resolver_createFirewallRuleCmd.MarkFlagRequired("creator-request-id")
		route53resolver_createFirewallRuleCmd.MarkFlagRequired("firewall-rule-group-id")
		route53resolver_createFirewallRuleCmd.MarkFlagRequired("name")
		route53resolver_createFirewallRuleCmd.MarkFlagRequired("priority")
	})
	route53resolverCmd.AddCommand(route53resolver_createFirewallRuleCmd)
}
