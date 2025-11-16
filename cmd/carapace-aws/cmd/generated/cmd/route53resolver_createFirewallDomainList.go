package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_createFirewallDomainListCmd = &cobra.Command{
	Use:   "create-firewall-domain-list",
	Short: "Creates an empty firewall domain list for use in DNS Firewall rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_createFirewallDomainListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_createFirewallDomainListCmd).Standalone()

		route53resolver_createFirewallDomainListCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows you to retry failed requests without the risk of running the operation twice.")
		route53resolver_createFirewallDomainListCmd.Flags().String("name", "", "A name that lets you identify the domain list to manage and use it.")
		route53resolver_createFirewallDomainListCmd.Flags().String("tags", "", "A list of the tag keys and values that you want to associate with the domain list.")
		route53resolver_createFirewallDomainListCmd.MarkFlagRequired("creator-request-id")
		route53resolver_createFirewallDomainListCmd.MarkFlagRequired("name")
	})
	route53resolverCmd.AddCommand(route53resolver_createFirewallDomainListCmd)
}
