package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_deleteFirewallDomainListCmd = &cobra.Command{
	Use:   "delete-firewall-domain-list",
	Short: "Deletes the specified domain list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_deleteFirewallDomainListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_deleteFirewallDomainListCmd).Standalone()

		route53resolver_deleteFirewallDomainListCmd.Flags().String("firewall-domain-list-id", "", "The ID of the domain list that you want to delete.")
		route53resolver_deleteFirewallDomainListCmd.MarkFlagRequired("firewall-domain-list-id")
	})
	route53resolverCmd.AddCommand(route53resolver_deleteFirewallDomainListCmd)
}
