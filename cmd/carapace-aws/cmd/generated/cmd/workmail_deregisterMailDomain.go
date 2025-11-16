package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deregisterMailDomainCmd = &cobra.Command{
	Use:   "deregister-mail-domain",
	Short: "Removes a domain from WorkMail, stops email routing to WorkMail, and removes the authorization allowing WorkMail use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deregisterMailDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_deregisterMailDomainCmd).Standalone()

		workmail_deregisterMailDomainCmd.Flags().String("domain-name", "", "The domain to deregister in WorkMail and SES.")
		workmail_deregisterMailDomainCmd.Flags().String("organization-id", "", "The WorkMail organization for which the domain will be deregistered.")
		workmail_deregisterMailDomainCmd.MarkFlagRequired("domain-name")
		workmail_deregisterMailDomainCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_deregisterMailDomainCmd)
}
