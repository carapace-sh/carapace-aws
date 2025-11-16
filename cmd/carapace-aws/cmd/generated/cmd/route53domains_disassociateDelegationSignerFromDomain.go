package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_disassociateDelegationSignerFromDomainCmd = &cobra.Command{
	Use:   "disassociate-delegation-signer-from-domain",
	Short: "Deletes a delegation signer (DS) record in the registry zone for this domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_disassociateDelegationSignerFromDomainCmd).Standalone()

	route53domains_disassociateDelegationSignerFromDomainCmd.Flags().String("domain-name", "", "Name of the domain.")
	route53domains_disassociateDelegationSignerFromDomainCmd.Flags().String("id", "", "An internal identification number assigned to each DS record after it’s created.")
	route53domains_disassociateDelegationSignerFromDomainCmd.MarkFlagRequired("domain-name")
	route53domains_disassociateDelegationSignerFromDomainCmd.MarkFlagRequired("id")
	route53domainsCmd.AddCommand(route53domains_disassociateDelegationSignerFromDomainCmd)
}
