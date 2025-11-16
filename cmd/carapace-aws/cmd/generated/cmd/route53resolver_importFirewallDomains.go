package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_importFirewallDomainsCmd = &cobra.Command{
	Use:   "import-firewall-domains",
	Short: "Imports domain names from a file into a domain list, for use in a DNS firewall rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_importFirewallDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_importFirewallDomainsCmd).Standalone()

		route53resolver_importFirewallDomainsCmd.Flags().String("domain-file-url", "", "The fully qualified URL or URI of the file stored in Amazon Simple Storage Service (Amazon S3) that contains the list of domains to import.")
		route53resolver_importFirewallDomainsCmd.Flags().String("firewall-domain-list-id", "", "The ID of the domain list that you want to modify with the import operation.")
		route53resolver_importFirewallDomainsCmd.Flags().String("operation", "", "What you want DNS Firewall to do with the domains that are listed in the file.")
		route53resolver_importFirewallDomainsCmd.MarkFlagRequired("domain-file-url")
		route53resolver_importFirewallDomainsCmd.MarkFlagRequired("firewall-domain-list-id")
		route53resolver_importFirewallDomainsCmd.MarkFlagRequired("operation")
	})
	route53resolverCmd.AddCommand(route53resolver_importFirewallDomainsCmd)
}
