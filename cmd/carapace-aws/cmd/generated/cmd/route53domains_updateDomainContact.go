package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_updateDomainContactCmd = &cobra.Command{
	Use:   "update-domain-contact",
	Short: "This operation updates the contact information for a particular domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_updateDomainContactCmd).Standalone()

	route53domains_updateDomainContactCmd.Flags().String("admin-contact", "", "Provides detailed contact information.")
	route53domains_updateDomainContactCmd.Flags().String("billing-contact", "", "Provides detailed contact information.")
	route53domains_updateDomainContactCmd.Flags().String("consent", "", "Customer's consent for the owner change request.")
	route53domains_updateDomainContactCmd.Flags().String("domain-name", "", "The name of the domain that you want to update contact information for.")
	route53domains_updateDomainContactCmd.Flags().String("registrant-contact", "", "Provides detailed contact information.")
	route53domains_updateDomainContactCmd.Flags().String("tech-contact", "", "Provides detailed contact information.")
	route53domains_updateDomainContactCmd.MarkFlagRequired("domain-name")
	route53domainsCmd.AddCommand(route53domains_updateDomainContactCmd)
}
