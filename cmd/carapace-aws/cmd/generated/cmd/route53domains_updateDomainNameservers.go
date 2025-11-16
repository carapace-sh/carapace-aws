package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_updateDomainNameserversCmd = &cobra.Command{
	Use:   "update-domain-nameservers",
	Short: "This operation replaces the current set of name servers for the domain with the specified set of name servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_updateDomainNameserversCmd).Standalone()

	route53domains_updateDomainNameserversCmd.Flags().String("domain-name", "", "The name of the domain that you want to change name servers for.")
	route53domains_updateDomainNameserversCmd.Flags().String("fiauth-key", "", "The authorization key for .fi domains")
	route53domains_updateDomainNameserversCmd.Flags().String("nameservers", "", "A list of new name servers for the domain.")
	route53domains_updateDomainNameserversCmd.MarkFlagRequired("domain-name")
	route53domains_updateDomainNameserversCmd.MarkFlagRequired("nameservers")
	route53domainsCmd.AddCommand(route53domains_updateDomainNameserversCmd)
}
