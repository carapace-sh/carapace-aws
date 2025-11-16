package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "This operation returns all the domain names registered with Amazon Route 53 for the current Amazon Web Services account if no filtering conditions are used.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_listDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_listDomainsCmd).Standalone()

		route53domains_listDomainsCmd.Flags().String("filter-conditions", "", "A complex type that contains information about the filters applied during the `ListDomains` request.")
		route53domains_listDomainsCmd.Flags().String("marker", "", "For an initial request for a list of domains, omit this element.")
		route53domains_listDomainsCmd.Flags().String("max-items", "", "Number of domains to be returned.")
		route53domains_listDomainsCmd.Flags().String("sort-condition", "", "A complex type that contains information about the requested ordering of domains in the returned list.")
	})
	route53domainsCmd.AddCommand(route53domains_listDomainsCmd)
}
