package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_transferDomainCmd = &cobra.Command{
	Use:   "transfer-domain",
	Short: "Transfers a domain from another registrar to Amazon Route 53.\n\nFor more information about transferring domains, see the following topics:\n\n- For transfer requirements, a detailed procedure, and information about viewing the status of a domain that you're transferring to Route 53, see [Transferring Registration for a Domain to Amazon Route 53](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html) in the *Amazon Route 53 Developer Guide*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_transferDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_transferDomainCmd).Standalone()

		route53domains_transferDomainCmd.Flags().String("admin-contact", "", "Provides detailed contact information.")
		route53domains_transferDomainCmd.Flags().String("auth-code", "", "The authorization code for the domain.")
		route53domains_transferDomainCmd.Flags().Bool("auto-renew", false, "Indicates whether the domain will be automatically renewed (true) or not (false).")
		route53domains_transferDomainCmd.Flags().String("billing-contact", "", "Provides detailed contact information.")
		route53domains_transferDomainCmd.Flags().String("domain-name", "", "The name of the domain that you want to transfer to Route 53. The top-level domain (TLD), such as .com, must be a TLD that Route 53 supports.")
		route53domains_transferDomainCmd.Flags().String("duration-in-years", "", "The number of years that you want to register the domain for.")
		route53domains_transferDomainCmd.Flags().String("idn-lang-code", "", "Reserved for future use.")
		route53domains_transferDomainCmd.Flags().String("nameservers", "", "Contains details for the host and glue IP addresses.")
		route53domains_transferDomainCmd.Flags().Bool("no-auto-renew", false, "Indicates whether the domain will be automatically renewed (true) or not (false).")
		route53domains_transferDomainCmd.Flags().Bool("no-privacy-protect-admin-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_transferDomainCmd.Flags().Bool("no-privacy-protect-billing-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_transferDomainCmd.Flags().Bool("no-privacy-protect-registrant-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_transferDomainCmd.Flags().Bool("no-privacy-protect-tech-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_transferDomainCmd.Flags().Bool("privacy-protect-admin-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_transferDomainCmd.Flags().Bool("privacy-protect-billing-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_transferDomainCmd.Flags().Bool("privacy-protect-registrant-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_transferDomainCmd.Flags().Bool("privacy-protect-tech-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_transferDomainCmd.Flags().String("registrant-contact", "", "Provides detailed contact information.")
		route53domains_transferDomainCmd.Flags().String("tech-contact", "", "Provides detailed contact information.")
		route53domains_transferDomainCmd.MarkFlagRequired("admin-contact")
		route53domains_transferDomainCmd.MarkFlagRequired("domain-name")
		route53domains_transferDomainCmd.MarkFlagRequired("duration-in-years")
		route53domains_transferDomainCmd.Flag("no-auto-renew").Hidden = true
		route53domains_transferDomainCmd.Flag("no-privacy-protect-admin-contact").Hidden = true
		route53domains_transferDomainCmd.Flag("no-privacy-protect-billing-contact").Hidden = true
		route53domains_transferDomainCmd.Flag("no-privacy-protect-registrant-contact").Hidden = true
		route53domains_transferDomainCmd.Flag("no-privacy-protect-tech-contact").Hidden = true
		route53domains_transferDomainCmd.MarkFlagRequired("registrant-contact")
		route53domains_transferDomainCmd.MarkFlagRequired("tech-contact")
	})
	route53domainsCmd.AddCommand(route53domains_transferDomainCmd)
}
