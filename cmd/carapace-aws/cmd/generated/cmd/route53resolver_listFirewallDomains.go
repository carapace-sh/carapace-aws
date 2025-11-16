package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listFirewallDomainsCmd = &cobra.Command{
	Use:   "list-firewall-domains",
	Short: "Retrieves the domains that you have defined for the specified firewall domain list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listFirewallDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listFirewallDomainsCmd).Standalone()

		route53resolver_listFirewallDomainsCmd.Flags().String("firewall-domain-list-id", "", "The ID of the domain list whose domains you want to retrieve.")
		route53resolver_listFirewallDomainsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Resolver to return for this request.")
		route53resolver_listFirewallDomainsCmd.Flags().String("next-token", "", "For the first call to this list request, omit this value.")
		route53resolver_listFirewallDomainsCmd.MarkFlagRequired("firewall-domain-list-id")
	})
	route53resolverCmd.AddCommand(route53resolver_listFirewallDomainsCmd)
}
