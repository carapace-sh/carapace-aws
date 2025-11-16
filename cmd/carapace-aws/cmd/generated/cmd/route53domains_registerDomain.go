package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_registerDomainCmd = &cobra.Command{
	Use:   "register-domain",
	Short: "This operation registers a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_registerDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_registerDomainCmd).Standalone()

		route53domains_registerDomainCmd.Flags().String("admin-contact", "", "Provides detailed contact information.")
		route53domains_registerDomainCmd.Flags().Bool("auto-renew", false, "Indicates whether the domain will be automatically renewed (`true`) or not (`false`).")
		route53domains_registerDomainCmd.Flags().String("billing-contact", "", "Provides detailed contact information.")
		route53domains_registerDomainCmd.Flags().String("domain-name", "", "The domain name that you want to register.")
		route53domains_registerDomainCmd.Flags().String("duration-in-years", "", "The number of years that you want to register the domain for.")
		route53domains_registerDomainCmd.Flags().String("idn-lang-code", "", "Reserved for future use.")
		route53domains_registerDomainCmd.Flags().Bool("no-auto-renew", false, "Indicates whether the domain will be automatically renewed (`true`) or not (`false`).")
		route53domains_registerDomainCmd.Flags().Bool("no-privacy-protect-admin-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_registerDomainCmd.Flags().Bool("no-privacy-protect-billing-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_registerDomainCmd.Flags().Bool("no-privacy-protect-registrant-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_registerDomainCmd.Flags().Bool("no-privacy-protect-tech-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_registerDomainCmd.Flags().Bool("privacy-protect-admin-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_registerDomainCmd.Flags().Bool("privacy-protect-billing-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_registerDomainCmd.Flags().Bool("privacy-protect-registrant-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_registerDomainCmd.Flags().Bool("privacy-protect-tech-contact", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_registerDomainCmd.Flags().String("registrant-contact", "", "Provides detailed contact information.")
		route53domains_registerDomainCmd.Flags().String("tech-contact", "", "Provides detailed contact information.")
		route53domains_registerDomainCmd.MarkFlagRequired("admin-contact")
		route53domains_registerDomainCmd.MarkFlagRequired("domain-name")
		route53domains_registerDomainCmd.MarkFlagRequired("duration-in-years")
		route53domains_registerDomainCmd.Flag("no-auto-renew").Hidden = true
		route53domains_registerDomainCmd.Flag("no-privacy-protect-admin-contact").Hidden = true
		route53domains_registerDomainCmd.Flag("no-privacy-protect-billing-contact").Hidden = true
		route53domains_registerDomainCmd.Flag("no-privacy-protect-registrant-contact").Hidden = true
		route53domains_registerDomainCmd.Flag("no-privacy-protect-tech-contact").Hidden = true
		route53domains_registerDomainCmd.MarkFlagRequired("registrant-contact")
		route53domains_registerDomainCmd.MarkFlagRequired("tech-contact")
	})
	route53domainsCmd.AddCommand(route53domains_registerDomainCmd)
}
