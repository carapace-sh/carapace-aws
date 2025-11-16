package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_checkDomainTransferabilityCmd = &cobra.Command{
	Use:   "check-domain-transferability",
	Short: "Checks whether a domain name can be transferred to Amazon Route 53.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_checkDomainTransferabilityCmd).Standalone()

	route53domains_checkDomainTransferabilityCmd.Flags().String("auth-code", "", "If the registrar for the top-level domain (TLD) requires an authorization code to transfer the domain, the code that you got from the current registrar for the domain.")
	route53domains_checkDomainTransferabilityCmd.Flags().String("domain-name", "", "The name of the domain that you want to transfer to Route 53. The top-level domain (TLD), such as .com, must be a TLD that Route 53 supports.")
	route53domains_checkDomainTransferabilityCmd.MarkFlagRequired("domain-name")
	route53domainsCmd.AddCommand(route53domains_checkDomainTransferabilityCmd)
}
