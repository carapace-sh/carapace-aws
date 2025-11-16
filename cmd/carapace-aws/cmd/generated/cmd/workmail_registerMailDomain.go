package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_registerMailDomainCmd = &cobra.Command{
	Use:   "register-mail-domain",
	Short: "Registers a new domain in WorkMail and SES, and configures it for use by WorkMail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_registerMailDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_registerMailDomainCmd).Standalone()

		workmail_registerMailDomainCmd.Flags().String("client-token", "", "Idempotency token used when retrying requests.")
		workmail_registerMailDomainCmd.Flags().String("domain-name", "", "The name of the mail domain to create in WorkMail and SES.")
		workmail_registerMailDomainCmd.Flags().String("organization-id", "", "The WorkMail organization under which you're creating the domain.")
		workmail_registerMailDomainCmd.MarkFlagRequired("domain-name")
		workmail_registerMailDomainCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_registerMailDomainCmd)
}
