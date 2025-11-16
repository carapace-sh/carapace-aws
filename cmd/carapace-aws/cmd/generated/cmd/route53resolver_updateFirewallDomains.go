package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateFirewallDomainsCmd = &cobra.Command{
	Use:   "update-firewall-domains",
	Short: "Updates the firewall domain list from an array of domain specifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateFirewallDomainsCmd).Standalone()

	route53resolver_updateFirewallDomainsCmd.Flags().String("domains", "", "A list of domains to use in the update operation.")
	route53resolver_updateFirewallDomainsCmd.Flags().String("firewall-domain-list-id", "", "The ID of the domain list whose domains you want to update.")
	route53resolver_updateFirewallDomainsCmd.Flags().String("operation", "", "What you want DNS Firewall to do with the domains that you are providing:")
	route53resolver_updateFirewallDomainsCmd.MarkFlagRequired("domains")
	route53resolver_updateFirewallDomainsCmd.MarkFlagRequired("firewall-domain-list-id")
	route53resolver_updateFirewallDomainsCmd.MarkFlagRequired("operation")
	route53resolverCmd.AddCommand(route53resolver_updateFirewallDomainsCmd)
}
