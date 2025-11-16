package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_renewDomainCmd = &cobra.Command{
	Use:   "renew-domain",
	Short: "This operation renews a domain for the specified number of years.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_renewDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_renewDomainCmd).Standalone()

		route53domains_renewDomainCmd.Flags().String("current-expiry-year", "", "The year when the registration for the domain is set to expire.")
		route53domains_renewDomainCmd.Flags().String("domain-name", "", "The name of the domain that you want to renew.")
		route53domains_renewDomainCmd.Flags().String("duration-in-years", "", "The number of years that you want to renew the domain for.")
		route53domains_renewDomainCmd.MarkFlagRequired("current-expiry-year")
		route53domains_renewDomainCmd.MarkFlagRequired("domain-name")
	})
	route53domainsCmd.AddCommand(route53domains_renewDomainCmd)
}
