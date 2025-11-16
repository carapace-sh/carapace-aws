package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_updateDomainContactPrivacyCmd = &cobra.Command{
	Use:   "update-domain-contact-privacy",
	Short: "This operation updates the specified domain contact's privacy setting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_updateDomainContactPrivacyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_updateDomainContactPrivacyCmd).Standalone()

		route53domains_updateDomainContactPrivacyCmd.Flags().Bool("admin-privacy", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_updateDomainContactPrivacyCmd.Flags().Bool("billing-privacy", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_updateDomainContactPrivacyCmd.Flags().String("domain-name", "", "The name of the domain that you want to update the privacy setting for.")
		route53domains_updateDomainContactPrivacyCmd.Flags().Bool("no-admin-privacy", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_updateDomainContactPrivacyCmd.Flags().Bool("no-billing-privacy", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_updateDomainContactPrivacyCmd.Flags().Bool("no-registrant-privacy", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_updateDomainContactPrivacyCmd.Flags().Bool("no-tech-privacy", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_updateDomainContactPrivacyCmd.Flags().Bool("registrant-privacy", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_updateDomainContactPrivacyCmd.Flags().Bool("tech-privacy", false, "Whether you want to conceal contact information from WHOIS queries.")
		route53domains_updateDomainContactPrivacyCmd.MarkFlagRequired("domain-name")
		route53domains_updateDomainContactPrivacyCmd.Flag("no-admin-privacy").Hidden = true
		route53domains_updateDomainContactPrivacyCmd.Flag("no-billing-privacy").Hidden = true
		route53domains_updateDomainContactPrivacyCmd.Flag("no-registrant-privacy").Hidden = true
		route53domains_updateDomainContactPrivacyCmd.Flag("no-tech-privacy").Hidden = true
	})
	route53domainsCmd.AddCommand(route53domains_updateDomainContactPrivacyCmd)
}
