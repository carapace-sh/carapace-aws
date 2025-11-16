package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getMailDomainCmd = &cobra.Command{
	Use:   "get-mail-domain",
	Short: "Gets details for a mail domain, including domain records required to configure your domain with recommended security.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getMailDomainCmd).Standalone()

	workmail_getMailDomainCmd.Flags().String("domain-name", "", "The domain from which you want to retrieve details.")
	workmail_getMailDomainCmd.Flags().String("organization-id", "", "The WorkMail organization for which the domain is retrieved.")
	workmail_getMailDomainCmd.MarkFlagRequired("domain-name")
	workmail_getMailDomainCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_getMailDomainCmd)
}
