package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_setIdentityMailFromDomainCmd = &cobra.Command{
	Use:   "set-identity-mail-from-domain",
	Short: "Enables or disables the custom MAIL FROM domain setup for a verified identity (an email address or a domain).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_setIdentityMailFromDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_setIdentityMailFromDomainCmd).Standalone()

		ses_setIdentityMailFromDomainCmd.Flags().String("behavior-on-mxfailure", "", "The action for Amazon SES to take if it cannot successfully read the required MX record when you send an email.")
		ses_setIdentityMailFromDomainCmd.Flags().String("identity", "", "The verified identity.")
		ses_setIdentityMailFromDomainCmd.Flags().String("mail-from-domain", "", "The custom MAIL FROM domain for the verified identity to use.")
		ses_setIdentityMailFromDomainCmd.MarkFlagRequired("identity")
	})
	sesCmd.AddCommand(ses_setIdentityMailFromDomainCmd)
}
