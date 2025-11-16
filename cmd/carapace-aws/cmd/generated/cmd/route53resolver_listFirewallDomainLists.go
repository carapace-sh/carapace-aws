package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listFirewallDomainListsCmd = &cobra.Command{
	Use:   "list-firewall-domain-lists",
	Short: "Retrieves the firewall domain lists that you have defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listFirewallDomainListsCmd).Standalone()

	route53resolver_listFirewallDomainListsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Resolver to return for this request.")
	route53resolver_listFirewallDomainListsCmd.Flags().String("next-token", "", "For the first call to this list request, omit this value.")
	route53resolverCmd.AddCommand(route53resolver_listFirewallDomainListsCmd)
}
