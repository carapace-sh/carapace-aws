package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_associateDelegationSignerToDomainCmd = &cobra.Command{
	Use:   "associate-delegation-signer-to-domain",
	Short: "Creates a delegation signer (DS) record in the registry zone for this domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_associateDelegationSignerToDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_associateDelegationSignerToDomainCmd).Standalone()

		route53domains_associateDelegationSignerToDomainCmd.Flags().String("domain-name", "", "The name of the domain.")
		route53domains_associateDelegationSignerToDomainCmd.Flags().String("signing-attributes", "", "The information about a key, including the algorithm, public key-value, and flags.")
		route53domains_associateDelegationSignerToDomainCmd.MarkFlagRequired("domain-name")
		route53domains_associateDelegationSignerToDomainCmd.MarkFlagRequired("signing-attributes")
	})
	route53domainsCmd.AddCommand(route53domains_associateDelegationSignerToDomainCmd)
}
