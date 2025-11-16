package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_checkDomainAvailabilityCmd = &cobra.Command{
	Use:   "check-domain-availability",
	Short: "This operation checks the availability of one domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_checkDomainAvailabilityCmd).Standalone()

	route53domains_checkDomainAvailabilityCmd.Flags().String("domain-name", "", "The name of the domain that you want to get availability for.")
	route53domains_checkDomainAvailabilityCmd.Flags().String("idn-lang-code", "", "Reserved for future use.")
	route53domains_checkDomainAvailabilityCmd.MarkFlagRequired("domain-name")
	route53domainsCmd.AddCommand(route53domains_checkDomainAvailabilityCmd)
}
