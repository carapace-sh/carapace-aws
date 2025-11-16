package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getFirewallDomainListCmd = &cobra.Command{
	Use:   "get-firewall-domain-list",
	Short: "Retrieves the specified firewall domain list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getFirewallDomainListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_getFirewallDomainListCmd).Standalone()

		route53resolver_getFirewallDomainListCmd.Flags().String("firewall-domain-list-id", "", "The ID of the domain list.")
		route53resolver_getFirewallDomainListCmd.MarkFlagRequired("firewall-domain-list-id")
	})
	route53resolverCmd.AddCommand(route53resolver_getFirewallDomainListCmd)
}
